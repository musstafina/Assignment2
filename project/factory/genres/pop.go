package genres

import "project/factory"

type PopGenreFactory struct{}

func (p *PopGenreFactory) CreateGenre() factory.MusicGenre {
	return &PopGenre{}
}

type PopGenre struct{}

func (p *PopGenre) Play() string {
	return "Playing a Pop track..."
}