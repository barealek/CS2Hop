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

	time.Sleep(3000)
	for {
		fmt.Println(client.GetFlags())
		time.Sleep(5 * time.Second)
	}

}
