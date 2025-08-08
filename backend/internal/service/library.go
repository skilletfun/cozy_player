package service

import (
	"gcozy_player/config"
	"gcozy_player/internal/container"
	"gcozy_player/internal/model"
	"gcozy_player/pkg/ffmpeg"
	"gcozy_player/pkg/track"
	"io/fs"
	"log"
	"os"
	"path/filepath"
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
	artistNames := map[string]uint{}
	artistTracks := map[uint]map[string]*model.Track{}

	if artists, err := l.artistService.GetAll(); err != nil {
		return err
	} else {
		for _, artist := range *artists {
			artistNames[artist.Name] = artist.ID
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

	var wg sync.WaitGroup
	var err error
	entries, _ := os.ReadDir(l.conf.MusicFolder)
	log.Println("Scan music directory: ", l.conf.MusicFolder)
	for _, entry := range entries {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if !entry.IsDir() {
				return
			}

			// Create artist if not exist, add to maps
			if _, ok := artistNames[entry.Name()]; !ok {
				artist := &model.Artist{Name: entry.Name()}
				if err = l.artistService.Create(artist); err != nil {
					return
				}
				if err = l.artistService.UpdateCover(artist); err != nil {
					return
				}
				artistNames[artist.Name] = artist.ID
				artistTracks[artist.ID] = map[string]*model.Track{}
			}

			artistId := artistNames[entry.Name()]
			artistPath := l.conf.MusicFolder + "/" + entry.Name()
			newTrackPaths := l.CollectNewTracks(artistPath, artistTracks[artistId])

			// Delete from DB tracks because not found while WalkDir
			for _, v := range artistTracks[artistId] {
				if err = l.trackService.Delete(v); err != nil {
					return
				}
			}

			if err = l.CreateNewTracks(newTrackPaths, artistId); err != nil {
				return
			}
		}()
	}
	wg.Wait()

	return l.artistService.DeleteWithoutTracks()
}

// CollectNewTracks returns paths to new tracks.
// Takes path to walk and map with current tracks in library. collect all new tracks. 
// While walk:
// 1) check file is track file, skip if doesn't;
// 2) if track file in library, remove it from tracks to prevent being deleted;
// 3) if track file not in library, add it to return result.
func (l *libraryService) CollectNewTracks(walkPath string, tracks map[string]*model.Track) []string {
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
	return newTrackPaths
}

// CreateNewTracks is the method to create new tracks for provided artist 
func (l *libraryService) CreateNewTracks(trackPaths []string, artistId uint) error {
	var newTracks []model.Track

	m := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, trackPath := range trackPaths {
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
				ArtistID: artistId,
			}

			m.Lock()
			newTracks = append(newTracks, newTrack)
			m.Unlock()
		}(trackPath)
	}
	wg.Wait()
	return l.trackService.BulkCreate(&newTracks)
}
