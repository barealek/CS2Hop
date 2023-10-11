package cs2hop

import (
	"errors"
	"fmt"

	"github.com/jamesmoriarty/gomem"
)

type Client struct {
	Process *gomem.Process
	Address uintptr
	Offsets *Offsets
}

func GetClientFromProcess(process *gomem.Process, offsets *Offsets) (*Client, error) {
	address, err := process.GetModule("client.dll")
	if err != nil {
		return nil, err
	}

	client := &Client{Process: process, Address: address, Offsets: offsets}
	return client, nil
}

func (c *Client) GetLocalPlayer() (uintptr, error) {
	ptr, err := c.Process.ReadUInt32(c.Address + uintptr(c.Offsets.ClientDll.DwLocalPlayerPawn))
	if err != nil {
		return 0, errors.New("failed to read localplayer: " + err.Error())
	}
	return (uintptr)(ptr), nil
}

func (c *Client) PlayerIsInAir() (bool, error) {
	pawn, err := c.GetLocalPlayer()
	if err != nil {
		return false, err
	}
	fmt.Println("Pawn address: ", pawn)

	// ...

	return false, nil
}
