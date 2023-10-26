package cs2hop

import (
	"errors"
	"fmt"
	"syscall"
	"unsafe"

	"github.com/jamesmoriarty/gomem"
)

var (
	kernel32               = syscall.NewLazyDLL("kernel32.dll")
	procOpenProcess        = kernel32.NewProc("OpenProcess")
	procReadProcessMemory  = kernel32.NewProc("ReadProcessMemory")
	procWriteProcessMemory = kernel32.NewProc("WriteProcessMemory")
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

	return &Client{Process: process, Address: address, Offsets: offsets}, nil
}

func (c *Client) GetLocalPlayer() (uintptr, error) {

	ptr, err := c.Process.ReadUInt32(c.Address + 25063768)
	if err != nil {
		return 0, errors.New("failed to read localplayer: " + err.Error())
	}
	return (uintptr)(ptr), nil
}

func (c *Client) GetFlags() error {
	fFlags, err := c.Process.ReadUInt32(c.Address + 25713704)

	// ...
	fmt.Println("Flags: ", fFlags)
	state := fFlags&(FL_ONGROUND) != 0
	fmt.Println(state)

	return err
}

func (c *Client) ForceJump() error {
	processHandle := c.Process.Handle

	client := uintptr(c.Address)
	var bytesRead uintptr
	var buffer [4]byte
	procReadProcessMemory.Call(processHandle, client+uintptr(25713704), uintptr(unsafe.Pointer(&buffer[0])), uintptr(unsafe.Sizeof(buffer)), uintptr(unsafe.Pointer(&bytesRead)))

	player := *(*uint32)(unsafe.Pointer(&buffer[0]))
	forceJump := client + uintptr(23716704)
	var bytesWritten uintptr
	procWriteProcessMemory.Call(processHandle, forceJump, uintptr(unsafe.Pointer(&player)), uintptr(unsafe.Sizeof(player)), uintptr(unsafe.Pointer(&bytesWritten)))

	return nil
}
