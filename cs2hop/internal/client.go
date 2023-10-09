package cs2hop

import "github.com/jamesmoriarty/gomem"

type Client struct {
	Process *gomem.Process
	Address uintptr
	Offsets *Offsets
}
