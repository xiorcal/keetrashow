package cmd

import (
	"fmt"
	"strings"
)

const (
	LAST_EPISODE = "no next, show finished"
	WRONG_DATA   = "invalid data"
)

type Show struct {
	title         string
	remote_api_id string
	seasons       []Season
	last_watched  Episode
	finished      bool
}

type Season struct {
	season   int
	episodes int
}

type Episode struct {
	season  int
	episode int
}

func (e *Episode) String() string {
	return fmt.Sprintf("S%vE%v", e.season, e.episode)
}

func (s *Season) String() string {
	return fmt.Sprintf("S%v (%v episodes)", s.season, s.episodes)
}
func (s *Show) String() string {
	maxLength := len(s.title)
	seasons := ""

	for i := 0; i < len(s.seasons); i++ {
		current := s.seasons[i].String()
		if s.seasons[i].season == s.last_watched.season {
			current = current + fmt.Sprintf("(next : %v)", s.last_watched.episode)
			if maxLength < len(current) {
				maxLength = len(current)
			}
		}
		seasons = seasons + "\n" + current
	}

	return fmt.Sprintf("%v\n%v%v\n%v\n", s.title, strings.Repeat("=", maxLength), seasons, strings.Repeat("-", maxLength))
}

// Watch increment from one episode, marking show as finished if it is
func (s *Show) Watch() (Show, error) {
	next, err := s.getNextEpisode(s.last_watched)
	if err != nil {
		if err == fmt.Errorf(LAST_EPISODE) {
			s.last_watched = Episode{}
			s.finished = true
		} else {
			return Show{}, err
		}
	}
	s.last_watched = next
	return *s, nil
}

func (s *Show) getNextEpisode(e Episode) (Episode, error) {
	currentSeason := Season{}

	for i := 0; i < len(s.seasons); i++ {
		if s.seasons[i].season == e.season {
			currentSeason = s.seasons[i]
			break
		}
	}
	if currentSeason.season == 0 {
		return Episode{}, fmt.Errorf(WRONG_DATA)
	}
	if e.episode+1 <= currentSeason.episodes {
		return Episode{currentSeason.season, e.episode + 1}, nil
	} else {
		nextSeason := Season{}
		for i := 0; i < len(s.seasons); i++ {
			if s.seasons[i].season == e.season+1 {
				nextSeason = s.seasons[i]
				break
			}
		}
		if nextSeason.season == 0 {
			return Episode{}, fmt.Errorf(LAST_EPISODE)
		}
		return Episode{nextSeason.season, 1}, nil
	}
}
