package api

import (
	"encoding/json"
	"fmt"
	"github.com/mahdi-cpp/api-go-emqx/repository"
	"os"
	"os/exec"
)

func Convert3GPToWAV(inputFile string, outputFile string) error {
	// Prepare the ffmpeg command with additional audio options
	cmd := exec.Command("ffmpeg", "-i", inputFile, "-acodec", "pcm_s16le", "-ac", "1", "-ar", "16000", outputFile)

	// Run the command and capture any error
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to convert file: %w", err)
	}
	return nil
}

func SaveRecordToFile(voice repository.Voice) error {

	// Create or open the file
	file, err := os.Create("/var/instagram/voices/" + voice.FileName + ".txt")
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Marshal the struct into JSON
	jsonData, err := json.MarshalIndent(voice, "", "    ") // Indent for readability
	if err != nil {
		return fmt.Errorf("error marshaling JSON: %v", err)
	}

	// Write the JSON string to the file
	_, err = file.WriteString(string(jsonData))
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}

func SaveJSONStringToFile(jsonString, filename string) error {
	// Create or open the file
	file, err := os.Create("/var/instagram/voices/" + filename)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close() // Ensure the file is closed after writing

	// Write the JSON string to the file
	_, err = file.WriteString(jsonString)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
