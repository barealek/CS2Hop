package cs2hop

import (
	"encoding/json"
	"io"
	"net/http"
)

type Offsets struct {
	ClientDll struct {
		DwForceJump   float64 `json:"dwForceJump"`
		DwLocalPlayer float64 `json:"dwLocalPlayerController"`
	} `json:"client_dll"`
}

func GetOffsets() (*Offsets, error) {
	var offsets Offsets

	// TODO: Download offsets from github.com/a2x/cs2-dumper

	resp, err := http.Get("https://raw.githubusercontent.com/a2x/cs2-dumper/main/generated/offsets.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &offsets)
	if err != nil {
		return nil, err
	}

	return &offsets, nil
}
