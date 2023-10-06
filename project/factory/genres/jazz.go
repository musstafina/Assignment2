package genres

import "project/factory"

type JazzGenreFactory struct{}

func (j *JazzGenreFactory) ClassicalGenre() factory.MusicGenre {
	return &JazzGenre{}
}

type JazzGenre struct{}

func (j *JazzGenre) Play() string {
	return "Playong a Jazz track..."
}