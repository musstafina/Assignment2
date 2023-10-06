package genres

import "project/factory"

type RockGenreFactory struct{}

func (r *RockGenreFactory) CreateGenre() factory.MusicGenre {
	return &RockGenre{}
}

type RockGenre struct{}

func (r *RockGenre) Play() string {
	return "Playing a Rock track ..."
}