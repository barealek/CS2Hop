package cs2hop

import "github.com/jamesmoriarty/gomem"

type Client struct {
	Process *gomem.Process
	Address uintptr
	Offsets *Offsets
}

// Needs work
func GetClientFrom(process *gomem.Process, offsets *Offsets) (*Client, error) {
	address, err := process.GetModule("client.dll")
	if err != nil {
		return nil, err
	}

	client := &Client{Process: process, Address: address, Offsets: offsets}
	return client, nil
}

func (c *Client) GetLocalPlayer() (uintptr, error) {
	ptr, _ := c.Process.ReadUInt32(c.Address + uintptr(c.Offsets.ClientDll.DwLocalPlayer))
	return (uintptr)(ptr), nil

}

func (c *Client) ForceJump() error {
	ptr, _ := c.Process.ReadUInt32(c.Address + uintptr(c.Offsets.ClientDll.DwForceJump))
	return c.Process.WriteByte(uintptr(ptr), 0x6)
}
