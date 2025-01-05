package utils

import (
	"github.com/appwrite/sdk-for-go/appwrite"
	"github.com/appwrite/sdk-for-go/models"
)

type LevelList struct {
	*models.DocumentList
	Documents []Level `json:"documents"`
}
type EditPostData struct {
	SongName string `json:"songName"`
	SongArtist string `json:"songArtist"`
	ChartName string `json:"chartName,omitempty"`
	Description string `json:"description,omitempty"`
	ID string `json:"id"`
}

type Level struct {
	SongName string `json:"songName"`
	SongArtist string `json:"songArtist"`
	ChartName string `json:"chartName,omitempty"`
	Description string `json:"description,omitempty"`
	BPM float64 `json:"bpm,omitempty"`
	GamesUsed []string `json:"gamesUsed,omitempty"`
	Uploader string `json:"uploader"`
	ID string `json:"$id"`
}

func CheckLevelExists(id string) bool {
	client := CreateClient()

	database := appwrite.NewDatabases(client)
	
	document, err := database.GetDocument(
		"mixtaper",
		"levels",
		string(id),
	)
	
	if err != nil || document == nil {
		return false
	}
	
	return true
}