package main

import "fmt"

type command interface {
	execute()
}

type Lock struct {
	command command
}

func (l *Lock) connect() {
	l.command.execute()
}

type remoteDevice interface {
	lock()
	unlock()
}

type unlockCommand struct {
	remoteDevice remoteDevice
}

func (u *unlockCommand) execute() {
	u.remoteDevice.unlock()
}

type lockCommand struct {
	remoteDevice remoteDevice
}

func (l *lockCommand) execute() {
	l.remoteDevice.lock()
}

type remoteServer struct {
	isAvailable bool
}

func (rmt *remoteServer) unlock() {
	rmt.isAvailable = true
	fmt.Println("Unlocking server")
}

func (rmt *remoteServer) lock() {
	rmt.isAvailable = false
	fmt.Println("Locking server")
}

func main() {
	srv := &remoteServer{}
	lockCommand := &lockCommand{
		remoteDevice: srv,
	}
	unlockCommand := &unlockCommand{
		remoteDevice: srv,
	}

	lock := &Lock{
		command: lockCommand,
	}
	lock.connect()

	lock = &Lock{
		command: unlockCommand,
	}
	lock.connect()
}
