package factory

type MusicGenreFactory interface{
	CreateGenre() MusicGenre
}

type MusicGenre interface {
	Play() string
}