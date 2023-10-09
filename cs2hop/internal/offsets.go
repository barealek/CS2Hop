package cs2hop

type Offsets struct {
	Timestamp string
}

func GetOffsets() (*Offsets, error) {
	var offsets Offsets

	// TODO: Download offsets from github.com/a2x/cs2-dumper

	return &offsets, nil
}
