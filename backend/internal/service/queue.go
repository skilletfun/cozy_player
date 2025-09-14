package service

import (
	"errors"
	"gcozy_player/internal/container"
	"gcozy_player/internal/model"
	"gcozy_player/pkg/structs"
	"gcozy_player/pkg/utils"
	"maps"
	"slices"

	"gorm.io/gorm"
)

type QueueService interface {
	GenerateNew() error
	GenerateByAll() error
	GenerateByArtist(int) error
	GenerateByTrack(int) error
	Next() (uint, error)
	Prev() (uint, error)
}

type GenerationType = int

const (
	All GenerationType = iota
	Artist
	Track
)

var QueueGenerationType = map[GenerationType]string{
	All:    "all",
	Artist: "artist",
	Track:  "track",
}

type queueService struct {
	db             *gorm.DB
	queue          structs.Stack[uint]
	history        structs.Stack[uint]
	generationType GenerationType
	currentTrack   uint
	trackService   TrackService
}

func NewQueueService(c container.Container) QueueService {
	return &queueService{
		db:           c.GetConnection(),
		trackService: NewTrackService(c),
		queue:        structs.NewStack[uint](),
		history:      structs.NewStack[uint](),
	}
}

func (q *queueService) GenerateNew() error {
	if q.generationType == All {
		return q.GenerateByAll()
	} else if q.currentTrack == 0 {
		return errors.New("cannot generate specific queue without history")
	} else if track, _ := q.trackService.GetByID(int(q.currentTrack)); q.generationType == Artist {
		return q.GenerateByArtist(int(track.ArtistID))
	} else {
		return q.GenerateByTrack(int(track.ID))
	}
}

// GenerateByAll is the method to generate new track's playing queue by all tracks.
func (q *queueService) GenerateByAll() error {
	if tracks, err := q.trackService.GetAll(); err != nil {
		return err
	} else {
		q.generationType = All
		q.Generate(tracks)
		return nil
	}
}

// GenerateByArtist is the method to generate new track's playing queue by provided artist.
func (q *queueService) GenerateByArtist(artist int) error {
	if tracks, err := q.trackService.GetAllByArtist(artist); err != nil {
		return err
	} else {
		q.generationType = Artist
		q.Generate(tracks)
		return nil
	}
}

// GenerateByTrack is the method to generate new track's playing queue with only provided track.
func (q *queueService) GenerateByTrack(artist int) error {
	if track, err := q.trackService.GetByID(artist); err != nil {
		return err
	} else {
		q.generationType = Track
		q.Generate(&[]model.Track{*track})
		return nil
	}
}

// Generate is the method to generate new track's playing queue by provided tracks.
func (q *queueService) Generate(tracks *[]model.Track) {
	q.queue.Clear()

	utils.Shuffle(tracks)
	parts := q.SortTracksByPlayCount(tracks)
	sortedPartKeys := slices.Collect(maps.Keys(parts))
	slices.Sort(sortedPartKeys)

	for _, key := range slices.Backward(sortedPartKeys) {
		for _, track := range q.ShuffleTracks(parts[key]) {
			q.queue.Push(track)
		}
	}
}

// SortTracksByPlayCount returns tracks sorted by play_count and artists.
func (q *queueService) SortTracksByPlayCount(tracks *[]model.Track) map[uint16]map[uint]structs.Stack[uint] {
	parts := map[uint16]map[uint]structs.Stack[uint]{}

	for _, track := range *tracks {
		if _, ok := parts[track.PlayCount]; !ok {
			parts[track.PlayCount] = make(map[uint]structs.Stack[uint])
		}
		if _, ok := parts[track.PlayCount][track.ArtistID]; !ok {
			parts[track.PlayCount][track.ArtistID] = structs.NewStack[uint]()
		}

		parts[track.PlayCount][track.ArtistID].Push(track.ID)
	}
	return parts
}

// ShuffleTracks returns shuffled tracks.
func (q *queueService) ShuffleTracks(tracks map[uint]structs.Stack[uint]) []uint {
	stacks := slices.Collect(maps.Values(tracks))
	slices.SortFunc(stacks, func(a structs.Stack[uint], b structs.Stack[uint]) int {
		if a.Size() > b.Size() {
			return 1
		} else if a.Size() < b.Size() {
			return -1
		}
		return 0
	})

	totalTracks := utils.SumSlice(stacks, func(s structs.Stack[uint]) int { return s.Size() })
	result := make([]uint, 0, totalTracks)

	index := len(stacks) - 1
	for len(stacks) > 0 {
		if stacks[index].IsEmpty() {
			stacks = slices.Delete(stacks, index, index+1)
			index--
			if index < 0 {
				index = len(stacks) - 1
			}
			continue
		}

		result = append(result, stacks[index].Pop())

		if index == 0 {
			index = len(stacks) - 1
			continue
		}

		// Check need to start over with the biggest stack
		hasStackOnLeft := index > 0
		hasStackOnRight := index < len(stacks)-1
		if hasStackOnRight && hasStackOnLeft {
			rightStackBiggerOrEqual := stacks[index+1].Size() >= stacks[index].Size()
			leftStackLessOrEqual := stacks[index-1].Size() <= stacks[index].Size()
			if rightStackBiggerOrEqual && leftStackLessOrEqual {
				index = len(stacks) - 1
				continue
			}
		}

		index--
	}
	return result
}

// Next returns the next track from track's playing queue.
// If playing queue is empty, new playing queue will be generated.
func (q *queueService) Next() (uint, error) {
	if q.queue.IsEmpty() {
		if err := q.GenerateNew(); err != nil {
			return 0, nil
		}
	}
	if q.queue.IsEmpty() {
		return 0, errors.New("queue empty after generation")
	}
	if q.currentTrack != 0 {
		q.trackService.IncrementPlayCount(int(q.currentTrack))
		q.history.Push(q.currentTrack)
	}

	q.currentTrack = q.queue.Pop()
	return q.currentTrack, nil
}

// Prev returns the track that played before current.
// If no played tracks in history, return error.
func (q *queueService) Prev() (uint, error) {
	if q.currentTrack == 0 {
		return 0, errors.New("empty history, no previous track")
	}
	if q.history.IsEmpty() {
		return q.currentTrack, nil
	}

	q.queue.Push(q.currentTrack)
	q.currentTrack = q.history.Pop()
	return q.currentTrack, nil
}
