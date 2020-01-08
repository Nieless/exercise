package main

type Song struct {
	name   string
	artist string
	next   *Song
}

type PlayList struct {
	name       string
	head       *Song // start
	nowPlaying *Song
}

func createPlaylist()  {
	
}