package repository

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strconv"
	"strings"
)

type VoiceDTO struct {
	Voices []Voice `json:"voices"`
}

var voiceDto VoiceDTO
var voiceDto2 VoiceDTO

type Voice struct {
	FileName    string  `json:"fileName"`
	Duration    int     `json:"duration"`
	Description string  `json:"description"`
	Timestamp   int64   `json:"timestamp"`
	Signals     []int16 `json:"signals"`
}

func GetVoices() VoiceDTO {

	dir := "/var/instagram/voices"
	var voiceDTO VoiceDTO

	voices, err := readTxtFiles(dir)
	log.Println("voices count: ", len(voices))

	if err != nil {
		log.Fatalf("Error reading files: %v", err)
	}

	// Print the data for demonstration
	for _, voice := range voices {
		voiceJSON, _ := json.MarshalIndent(voice, "", "  ")
		fmt.Printf("Voice: %s\n", voiceJSON)
	}
	voiceDTO.Voices = voices
	return voiceDTO
}

// Function to read .txt files from a directory and convert their content to Voice structs
func readTxtFiles(dir string) ([]Voice, error) {
	var voices []Voice

	// Read all files in the directory
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	// Iterate through the files
	for _, file := range files {

		// Check if the file is a .txt file

		if filepath.Ext(file.Name()) == ".txt" {

			// Read the content of the file
			data, err := ioutil.ReadFile(filepath.Join(dir, file.Name()))
			if err != nil {
				log.Printf("Error reading file %s: %v", file.Name(), err)
				continue
			}

			var voice Voice
			err = json.Unmarshal(data, &voice)
			if err != nil {
				log.Fatalf("Error unmarshaling JSON: %v", err)
			}

			voices = append(voices, voice)
		}
	}

	return voices, nil
}

// Function to parse signal data from a string
func parseSignals(signalStr string) []int16 {
	var signals []int16
	signalValues := strings.Fields(signalStr) // Split by whitespace

	for _, value := range signalValues {
		if signal, err := strconv.ParseInt(value, 10, 16); err == nil {
			signals = append(signals, int16(signal))
		} else {
			log.Printf("Invalid signal value: %s, error: %v", value, err)
		}
	}

	return signals
}
