package service

import (
	"gcozy_player/config"
	"gcozy_player/internal/container"
	"gcozy_player/internal/model"
	"gcozy_player/pkg/ffmpeg"
	"gcozy_player/pkg/track"
	"io/fs"
	"log"
	"maps"
	"os"
	"path/filepath"
	"slices"
	"sync"

	"gorm.io/gorm"
)

type LibraryService interface {
	Rescan() error
}

type libraryService struct {
	conf          *config.Config
	db            *gorm.DB
	trackService  TrackService
	artistService ArtistService
}

func NewLibraryService(c container.Container) LibraryService {
	return &libraryService{
		conf:          c.GetConfig(),
		db:            c.GetConnection(),
		trackService:  NewTrackService(c),
		artistService: NewArtistService(c),
	}
}

// Rescan is the method to rescan and update library.
// Method:
// - create new artists with tracks;
// - create new tracks;
// - delete artists without tracks;
// - delete tracks that no found while rescan. 
func (l *libraryService) Rescan() error {
	artistsByNames := map[string]model.Artist{}
	artistTracks := map[uint]map[string]*model.Track{}

	if artists, err := l.artistService.GetAll(); err != nil {
		return err
	} else {
		for _, artist := range *artists {
			artistsByNames[artist.Name] = artist
		}
	}

	if tracks, err := l.trackService.GetAll(); err != nil {
		return err
	} else {
		for _, track := range *tracks {
			if _, ok := artistTracks[track.ArtistID]; !ok {
				artistTracks[track.ArtistID] = map[string]*model.Track{}
			}
			artistTracks[track.ArtistID][track.Path] = &track
		}
	}
	
	if artists, err := l.ActualizeArtists(artistsByNames); err != nil {
		return err
	} else {
		var wg sync.WaitGroup
		
		tracksForCreate := []model.Track{}
		tracksForDelete := []model.Track{}
		
		for _, artist := range artists {
			wg.Add(1)
			go func() {
				defer wg.Done()

				if _, ok := artistTracks[artist.ID]; !ok {
					artistTracks[artist.ID] = map[string]*model.Track{}
				}
				
				newTracks := l.CollectNewTracks(artist, artistTracks[artist.ID])
				tracksForCreate = append(tracksForCreate, *newTracks...)

				for _, v := range artistTracks[artist.ID] {
					tracksForDelete = append(tracksForDelete, *v)
				}
			}()
		}
		wg.Wait()
		
		// Create new tracks
		if len(tracksForCreate) > 0 {
			if err := l.trackService.BulkCreate(&tracksForCreate); err != nil {
				return err
			}
		}
		
		// Delete tracks that not found
		if len(tracksForDelete) > 0 {
			if err := l.trackService.BulkDelete(&tracksForDelete); err != nil {
				return err
			}
		}
	}
	
	return l.artistService.DeleteWithoutTracks()
}

func (l *libraryService) ActualizeArtists(artistsByNames map[string]model.Artist) (map[string]model.Artist, error) {
	entries, _ := os.ReadDir(l.conf.MusicFolder)
	newArtists := []model.Artist{}
	existedArtists := map[string]model.Artist{}
	
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		if _, ok := artistsByNames[entry.Name()]; !ok {
			log.Println("Found new artist: ", entry.Name())
			artist := model.Artist{Name: entry.Name()}
			newArtists = append(newArtists, artist)
		} else {
			existedArtists[entry.Name()] = artistsByNames[entry.Name()]
			delete(artistsByNames, entry.Name())
		}
	}
	
	if len(newArtists) > 0 {
		log.Println("Found new artist for create: ", len(newArtists))
		if err := l.artistService.BulkCreate(&newArtists); err != nil {
			return nil, err
		}
	}
	 
	for _, artist := range newArtists {
		existedArtists[artist.Name] = artist
	}
	
	artistsForDelete := slices.Collect(maps.Values(artistsByNames))
	if len(artistsForDelete) > 0 {
		log.Println("Found artist for delete: ", len(artistsForDelete))
		if err := l.artistService.BulkDelete(&artistsForDelete); err != nil {
			return nil, err
		}
	}
	
	if err := l.artistService.UpdateCovers(); err != nil {
		return nil, err
	}
	
	return existedArtists, nil
}

// CollectNewTracks returns paths to new tracks.
// Takes path to walk and map with current tracks in library. collect all new tracks. 
// While walk:
// 1) check file is track file, skip if doesn't;
// 2) if track file in library, remove it from tracks to prevent being deleted;
// 3) if track file not in library, add it to return result.
func (l *libraryService) CollectNewTracks(artist model.Artist, tracks map[string]*model.Track) *[]model.Track {
	walkPath := l.conf.MusicFolder + "/" + artist.Name
	log.Println("Scan artist directory: ", walkPath)
	
	var newTrackPaths []string
	filepath.WalkDir(
		walkPath,
		func(path string, d fs.DirEntry, err error) error {
			if err == nil && !d.IsDir() && track.IsTrackFile(d.Name()) {
				if _, ok := tracks[d.Name()]; ok {
					delete(tracks, d.Name())
				} else {
					newTrackPaths = append(newTrackPaths, path)
				}
			}
			return err
		},
	)
	log.Println("Found ", len(newTrackPaths), "tracks in artist: ", walkPath)
	
	var newTracks []model.Track

	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, trackPath := range newTrackPaths {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			info, err := ffmpeg.GetFileInfo(trackPath)
			if err != nil {
				log.Println("Error ffmpeg.GetFileInfo(trackPath): ", err.Error())
				return
			}
			newTrack := model.Track{
				Title:    info.Tags.Title,
				Album:    info.Tags.Album,
				Duration: uint16(info.Duration),
				Path:     trackPath,
				ArtistID: artist.ID,
			}

			m.Lock()
			newTracks = append(newTracks, newTrack)
			m.Unlock()
		}(trackPath)
	}
	wg.Wait()
	
	return &newTracks
}
