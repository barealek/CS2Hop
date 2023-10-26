package main

import (
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

	client, err := cs2hop.GetClientFromProcessName("cs2.exe", &offsets)
	if err != nil {
		panic(err)
	}

	for {
		flags, err := client.GetFlags()
		if err != nil {
			panic(err)
		}
		onGround := flags&(cs2hop.FL_ONGROUND) != 0
		if onGround {
			client.ForceJump()
		}
	}
}
