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

func GetClientFromProcessName(processName string, offsets *Offsets) (*Client, error) {
	process, err := gomem.GetProcessFromName(processName)
	if err != nil {
		return nil, errors.New("failed to get process: " + processName)
	}

	process.Open()

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
	lp, err := c.GetLocalPlayer()
	if err != nil {
		return false, err
	}

	flags, err := c.Process.ReadByte(lp + uintptr(968))
	if err != nil {
		fmt.Println("FEjl")
		return false, err
	}
	fmt.Println("Flags: ", flags)

	return false, nil
}

func (c *Client) ForceJump() error {
	lp, err := c.GetLocalPlayer()
	if err != nil {
		return err
	}

	flags, err := c.Process.ReadByte(lp + uintptr(c.Offsets.ClientDll.DwForceJump))
	fmt.Println("Flags: ", flags)
	// ...
	return err
}
