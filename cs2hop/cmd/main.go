package main

import (
	"fmt"

	cs2hop "github.com/barealek/cs2hop/internal"
)

var (
	offsets = cs2hop.Offsets{}
)

func main() {
	err := offsets.FetchOffsets()
	if err != nil {
		panic(err)
	}

	fmt.Println(offsets.ClientDll.DwLocalPlayer)
}
