package decorators

import (
	"fmt"
)

type VolumeDecorator struct{
	VolumeLevel int
}

func (v *VolumeDecorator) Decorate(playlist string) string {
	return fmt.Sprintf("Adjusting volume of %s to %d", playlist, v.VolumeLevel)
}
