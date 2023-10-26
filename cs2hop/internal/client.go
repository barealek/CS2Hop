package cs2hop

import (
	"errors"
	"syscall"
	"time"
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

func (c *Client) GetLocalPlayerController() (uintptr, error) {

	ptr, err := c.Process.ReadUInt32(c.Address + uintptr(c.Offsets.DwLocalPlayerController.Value))
	if err != nil {
		return 0, errors.New("failed to read localplayer: " + err.Error())
	}
	return (uintptr)(ptr), nil
}

func (c *Client) GetLocalPlayerPawn() (uintptr, error) {

	ptr, err := c.Process.ReadUInt64(c.Address + uintptr(c.Offsets.DwLocalPlayerPawn.Value))
	if err != nil {
		return 0, errors.New("failed to read localplayer: " + err.Error())
	}
	return (uintptr)(ptr), nil
}

func (c *Client) GetFlags() (uint32, error) {
	lpPawn, _ := c.GetLocalPlayerPawn()
	fFlags, err := c.Process.ReadUInt32(lpPawn + uintptr(968))
	if err != nil {
		return 0, err
	}
	return fFlags, nil
}

func (c *Client) ForceJump() error {
	processHandle := c.Process.Handle

	client := uintptr(c.Address)
	var bytesRead uintptr
	var buffer [4]byte
	procReadProcessMemory.Call(processHandle, client+uintptr(c.Offsets.DwLocalPlayerPawn.Value), uintptr(unsafe.Pointer(&buffer[0])), uintptr(unsafe.Sizeof(buffer)), uintptr(unsafe.Pointer(&bytesRead)))

	player, _ := c.GetLocalPlayerController()
	forceJump := client + uintptr(c.Offsets.DwForceJump.Value)
	var bytesWritten uintptr
	procWriteProcessMemory.Call(processHandle, forceJump, uintptr(unsafe.Pointer(&player)), uintptr(unsafe.Sizeof(player)), uintptr(unsafe.Pointer(&bytesWritten)))

	time.Sleep(3 * time.Millisecond)

	newForceJumpValue := uint32(256)
	procWriteProcessMemory.Call(processHandle, forceJump, uintptr(unsafe.Pointer(&newForceJumpValue)), uintptr(unsafe.Sizeof(newForceJumpValue)), uintptr(unsafe.Pointer(&bytesWritten)))

	time.Sleep(3 * time.Millisecond)

	return nil
}
