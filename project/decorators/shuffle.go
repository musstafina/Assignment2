package decorators

import (
	"fmt"
)

type ShuffleDecorator struct{}

func (s *ShuffleDecorator) Decorate(playlist string) string {
	return fmt.Sprintf("Shuffling %s", playlist)
}