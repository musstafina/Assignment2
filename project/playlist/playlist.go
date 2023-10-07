package playlist

import (
	"fmt"
	"project/decorators"
)

type PlaylistManager struct {
	playlist string
	decorators []decorators.PlaylistDecorator
	songs []string
}

func NewPlaylistManager(playlist string) *PlaylistManager {
	return &PlaylistManager{
		playlist: playlist,
		decorators: []decorators.PlaylistDecorator{},
		songs: []string{},
	}
}

func(p *PlaylistManager) ApplyDecorator(decorator decorators.PlaylistDecorator) {
	p.decorators = append(p.decorators, decorator)

}

func (p *PlaylistManager) AddSong(song string) {
	p.songs = append(p.songs, song)
}

func (p *PlaylistManager) PlayPlaylist() string {
	songOut := ""
	for _, song := range p.songs {
		songOut += fmt.Sprintf("Playlist song: %s\n", song)
	}

	decoratedPlaylist := ""
	for _, decorator := range p.decorators {
		decoratedPlaylist += decorator.Decorate(p.playlist) + "\n"
	}

	result := decoratedPlaylist + songOut
	return result


}
	// result := p.playlist
	// for _, decorator := range p.decorators {
	// 	result = decorator.Decorate(result)
	// }

	// for _, song := range p.songs {
	// 	result += "\n" + "Playlist song: " + song
	// }
	// return result

