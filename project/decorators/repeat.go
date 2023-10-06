package decorators

import (
	"fmt"
)

type RepeatDecorator struct {
	RepeatTimes int

}

func (r *RepeatDecorator) Decorate(playlist string) string {
	return fmt.Sprintf("Repeating %s %d times", playlist, r.RepeatTimes)
}