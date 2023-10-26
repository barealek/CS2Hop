package cs2hop

import (
	"encoding/json"
	"io"
	"net/http"
)

type Offsets struct {
	DwLocalPlayerController offset `json:"dwLocalPlayerController"`
	DwLocalPlayerPawn       offset `json:"dwLocalPlayerPawn"`
	DwForceJump             offset `json:"dwForceJump"`
}

type offsetsWrapper struct {
	ClientDll struct {
		Data map[string]struct {
			Value uint32 `json:"value"`
		} `json:"data"`
	} `json:"client_dll"`
}

type offset struct {
	Value float64 `json:"value"`
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

	var wrapper offsetsWrapper
	var wrapperBytes []byte

	err = json.Unmarshal(bytes, &wrapper)
	wrapperBytes, err = json.Marshal(wrapper.ClientDll.Data)
	err = json.Unmarshal(wrapperBytes, o)

	if err != nil {
		return err
	}

	return nil
}
