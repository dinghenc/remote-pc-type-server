package robot

import (
	"fmt"

	"github.com/go-vgo/robotgo"
)

type Operator struct {
}

func (o *Operator) PasteString(text string) error {
	robotgo.Click()
	if err := robotgo.KeyTap(robotgo.KeyA, robotgo.CmdCtrl()); err != nil {
		return fmt.Errorf("key tap ctrl+a failed: %w", err)
	}
	if err := robotgo.KeyTap(robotgo.Backspace); err != nil {
		return fmt.Errorf("key tap backspace failed: %w", err)
	}
	robotgo.TypeStr(text)
	return nil
}

func (o *Operator) Enter() error {
	if err := robotgo.KeyTap(robotgo.Enter); err != nil {
		return fmt.Errorf("key tap enter failed: %w", err)
	}
	return nil
}
