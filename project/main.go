package main

import (
	"fmt"
	"project/decorators"
	"project/factory/genres"
	"project/playlist"
)

func main() {
	rockFactory := &genres.RockGenreFactory{}	
	rockPlaylist := rockFactory.CreateGenre()

	rockManager := playlist.NewPlaylistManager("Rock Playlist")
	rockManager.ApplyDecorator(&decorators.ShuffleDecorator{})
	rockManager.ApplyDecorator(&decorators.VolumeDecorator{VolumeLevel: 50})
	rockManager.ApplyDecorator(&decorators.RepeatDecorator{RepeatTimes: 2})

	rockManager.AddSong("Song 1")
	rockManager.AddSong("Song 2")
	rockManager.AddSong("Song 3")

	fmt.Println(rockPlaylist.Play())
	fmt.Println("\nPlaying Rock Playlist:")
	fmt.Println(rockManager.PlayPlaylist())
}
