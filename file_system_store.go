package main

import (
	"encoding/json"
	"io"
	"os"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, io.SeekStart)
	league, _ := NewLeague(database)
	return &FileSystemPlayerStore{
		database: &tape{database.(*os.File)},
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(playerName string) int {
	player := f.league.Find(playerName)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(playerName string) {
	player := f.league.Find(playerName)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{playerName, 1})
	}

	json.NewEncoder(f.database).Encode(f.league)
}
