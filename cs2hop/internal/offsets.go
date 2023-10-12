package cs2hop

import (
	"encoding/json"
	"io"
	"net/http"
)

type Offsets struct {
	ClientDll struct {
		DwLocalPlayerController float64 `json:"dwLocalPlayerController"`
		DwLocalPlayerPawn       float64 `json:"dwLocalPlayerPawn"`
		DwPlantedC4             float64 `json:"dwPlantedC4"`
		DwForceJump             float64 `json:"dwForceJump"`
	} `json:"client_dll"`
}

func (o *Offsets) FetchOffsets() error {
	// TODO: Download offsets from github.com/a2x/cs2-dumper

	resp, err := http.Get("https://raw.githubusercontent.com/a2x/cs2-dumper/main/generated/offsets.json")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &o)
	if err != nil {
		return err
	}
	return nil
}
