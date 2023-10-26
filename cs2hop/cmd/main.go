package main

import (
	"fmt"
	"time"

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
		time.Sleep(1000 * time.Millisecond)
		fmt.Println("ForceJump Error:", client.ForceJump())
	}

}
