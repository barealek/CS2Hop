package main

import (
	"fmt"

	cs2hop "github.com/barealek/cs2hop/internal"
	"github.com/jamesmoriarty/gomem"
)

var (
	offsets = cs2hop.Offsets{}
)

func main() {
	err := offsets.FetchOffsets()
	if err != nil {
		panic(err)
	}

	process, err := gomem.GetProcessFromName("cs2.exe")
	if err != nil {
		panic(err)
	}

	client, err := cs2hop.GetClientFrom(process, &offsets)
	if err != nil {
		panic(err)
	}

	if err = client.ForceJump(); err != nil {
		fmt.Println("ERROR:")
		fmt.Println(err)
	}
}
