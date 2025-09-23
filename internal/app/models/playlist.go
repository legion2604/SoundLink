package models

type Playlist struct {
	PlaylistName string `json:"playlist_name"`
	Category     string `json:"category"`
}
type AddMusicToPlaylist struct {
	PlaylistID string `json:"playlist_id"`
	MusicID    string `json:"music_id"`
}
type Music struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}
type PlaylistForm struct {
	Id           string `json:"id"`
	PlaylistName string `json:"playlistName"`
	Category     string `json:"category"`
}

type MusicIdPlaylistId struct {
	MusicId    string `json:"music_id"`
	PlaylistId string `json:"playlist_id"`
}
