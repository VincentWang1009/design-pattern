package command

import "testing"

func TestInit(t *testing.T) {
	tv := &Tv{}

	onCommand := &OnCommand{
		device: tv,
	}

	OffCommand := &OffCommand{
		device: tv,
	}

	onCommandButton := &Button{
		command: onCommand,
	}

	onCommandButton.press()

	offCommandButton := &Button{
		command: OffCommand,
	}

	offCommandButton.press()
}
