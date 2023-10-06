package genres

import "project/factory"

type ClassicalGenreFactory struct{}

func (c *ClassicalGenreFactory) CreateGenre() factory.MusicGenre {
	return &ClassicalGenre{}
}

type ClassicalGenre struct{}

func (c *ClassicalGenre) Play() string {
	return "Playing a Classical track..."
}