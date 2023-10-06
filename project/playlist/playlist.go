package playlist

import (
	"project/decorators"
)

type PlaylistManager struct {
	playlist string
	decorators []decorators.PlaylistDecorator
}

func NewPlaylistManager(playlist string) *PlaylistManager {
	return &PlaylistManager{
		playlist: playlist,
		decorators: []decorators.PlaylistDecorator{},
	}
}

func(p *PlaylistManager) ApplyDecorator(decorator decorators.PlaylistDecorator) {
	p.decorators = append(p.decorators, decorator)

}

func (p *PlaylistManager) PlayPlaylist() string {
	result := p.playlist
	for _, decorator := range p.decorators {
		result = decorator.Decorate(result)
	}
	return result
}
