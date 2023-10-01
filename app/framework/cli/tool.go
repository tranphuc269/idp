package cli

import (
	"idp_system/app/framework/event/eventbus"
	"idp_system/app/framework/terminal"
	"os"
	"os/signal"
	"syscall"
)

type GUI struct {
	term *terminal.Terminal
}

func (c *GUI) EnterMainLoop(handleEvents func()) {
	// Prevent process from being killed by pressing Ctrl + C
	signal.Ignore(syscall.SIGINT)
	c.term.StopWaitForEnter()
	c.term.StartEventLoop()
	c.term.TurnOffEcho()

	for {
		handleEvents()
	}
}

func (c *GUI) OnKeyPress(keyName terminal.KeyName, callback eventbus.DataChannel) {
	c.term.OnKeyPress(keyName, callback)
}

func (c *GUI) Exit() {
	c.term.TurnOnEcho()
	os.Exit(0)
}

func NewGUI(term *terminal.Terminal) GUI {
	return GUI{
		term: term,
	}
}
