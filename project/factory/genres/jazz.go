package genres

import "project/factory"

type JazzGenreFactory struct{}

func (j *JazzGenreFactory) CreateGenre() factory.MusicGenre {
	return &JazzGenre{}
}

type JazzGenre struct{}

func (j *JazzGenre) Play() string {
	return "Playing a Jazz track..."
}