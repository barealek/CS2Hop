package main

import (
	"time"

	cs2hop "github.com/barealek/cs2hop/internal"
	"github.com/jamesmoriarty/gomem"
)

var (
	offsets  = cs2hop.Offsets{}
	VK_SPACE = 0x20
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
		if gomem.IsKeyDown(VK_SPACE) {
			flags, err := client.GetFlags()
			if err != nil {
				panic(err)
			}

			onGround := flags&(cs2hop.FL_ONGROUND) != 0

			if onGround {
				client.ForceJump()
			}

		}
		time.Sleep(1 * time.Millisecond)
	}
}
