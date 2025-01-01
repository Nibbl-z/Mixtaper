package utils

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"io"
	"regexp"
)

func CheckRiq(path string) error {
	reader, err := zip.OpenReader(path)

	if err != nil {
		return err
	}

	defer reader.Close()
	
	hasRemix, hasSong := false, false
	
	for _, file := range reader.File {
		if file.Name == "remix.json" {
			hasRemix = true
		}
		
		if file.Name == "song.ogg" {
			hasSong = true
		}
	}
	
	if !hasRemix {
		return errors.New("remix.json is missing from the .riq")
	}

	if !hasSong {
		return errors.New("song.ogg is missing from the .riq")
	}

	return nil
}

func GetBPM(data map[string]interface{}) float64 {
	if data["bpm"] != nil {
		return data["bpm"].(float64)
	}
	
	if len(data["tempoChanges"].([]map[string]interface{})) >= 1 {
		return data["tempoChanges"].([]map[string]interface{})[0]["tempo"].(float64) // Erm what the go
	}
	
	return 0.0
}

func contains(array []string, search string) bool {
	for _, v := range array {
		if v == search {
			return true
		}
	}
	
	return false
}

func GetGames(data map[string]interface{}) []string {
	entities := data["entities"].([]interface{}) 
	
	var games []string
	
	for _, entity := range entities {
		datamodel := entity.(map[string]interface{})["datamodel"].(string)
		getNameExpression := regexp.MustCompile(`^(.*?)/.*`)
		
		gameName := getNameExpression.FindStringSubmatch(datamodel)
		
		if len(gameName) > 1 && !contains(games, gameName[1]) {
			games = append(games, gameName[1])
		}
	}

	return games
}

func GetRemixData(path string) (map[string]interface{}, error) {
	reader, err := zip.OpenReader(path)

	if err != nil {
		return nil, err
	}

	defer reader.Close()
	
	var remixBytes []byte
	
	for _, file := range reader.File {
		if file.Name == "remix.json" {
			fileReader, err := file.Open()
			if err != nil {
				return nil, err
			}
			defer fileReader.Close()
			
			remixBytes, err = io.ReadAll(fileReader)
			if err != nil {
				return nil, err
			}

			break
		}
	}
	
	var remixData map[string]interface{}

	err = json.Unmarshal(remixBytes, &remixData)
	if err != nil {
		return nil, err
	}

	return remixData, nil
}