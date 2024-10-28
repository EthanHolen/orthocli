package orthoapi

import (
	"encoding/json"
	"io"
	"net/http"
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

func MakeRequest(url string) (OrthoResponse, error) {
	var orthoResponse OrthoResponse
	apiResponse, err := http.Get(url)

	if err != nil {
		return orthoResponse, err
	}

	responseData, err := io.ReadAll(apiResponse.Body)

	if err != nil {
		return orthoResponse, err
	}

	json.Unmarshal(responseData, &orthoResponse)

	return orthoResponse, nil
}
