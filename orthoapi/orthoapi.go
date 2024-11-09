package orthoapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type OrthoResponse struct {
	Year                  int      `json:"year"`
	Month                 int      `json:"month"`
	Day                   int      `json:"day"`
	PaschaDistance        int      `json:"pascha_distance"`
	Titles                []string `json:"titles"`
	FastLevel             int      `json:"fast_level"`
	FastLevelDescription  string   `json:"fast_level_desc"`
	FeastLevel            int      `json:"feast_level"`
	FeastLevelDescription string   `json:"feast_level_desc"`
}

const DEFAULT_URL string = "https://orthocal.info/api/gregorian/"
const JULIAN_URL string = "https://orthocal.info/api/julian/"
const HIDDEN_DIR_NAME string = "orthocli-"

func GetFullResponse(calendarType string) (OrthoResponse, error) {

	var orthoResponse OrthoResponse

	// TODO: move all of this logic to the function meant to pull the entire year
	apiResponse, err := makeApiRequest(DEFAULT_URL)

	if err != nil {
		return orthoResponse, err
	}

	json.Unmarshal(apiResponse, &orthoResponse)

	dirName, err := createTempDir()
	if err != nil {
		fmt.Println("Failed to create dir: ", err)
	}
	fmt.Printf("Created dir: %s", dirName)

	return orthoResponse, nil

}

func makeApiRequest(url string) ([]byte, error) {
	apiResponse, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	responseData, err := io.ReadAll(apiResponse.Body)

	return responseData, err

	// json.Unmarshal(responseData, &orthoResponse)

	// return orthoResponse, nil
}

func createTempDir() (string, error) {

	dirName, err := os.MkdirTemp("", HIDDEN_DIR_NAME)
	if err != nil {
		return "", err
	}

	return dirName, err

}

func findTempDir(prefix string) (string, error) {
	// Get the temporary directory location
	tempDir := os.TempDir()

	// Read the contents of the temporary directory
	files, err := os.ReadDir(tempDir)
	if err != nil {
		return "", fmt.Errorf("error reading temp directory: %w", err)
	}

	// Search for the directory with the specified prefix
	for _, file := range files {
		if file.IsDir() && strings.HasPrefix(file.Name(), prefix) {
			return filepath.Join(tempDir, file.Name()), nil // Return found directory path
		}
	}

	return "", fmt.Errorf("no temporary directory found with prefix: %s", prefix)
}

// func createHiddenDir() error {
// 	dirName := ""
// 	if !strings.HasPrefix(HIDDEN_DIR_NAME, ".") {
// 		dirName = "." + HIDDEN_DIR_NAME
// 	}

// 	err := os.Mkdir(dirName, 0755)
// 	if err != nil {
// 		return err
// 	}

// 	return nil

// }
