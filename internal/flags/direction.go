package flags

import (
	"errors"
	"fmt"
)

type Direction string

const (
	Straight Direction = "straight"
	Inverse  Direction = "inverse"
	Random   Direction = "random"
)

func (flag *Direction) Name() string {
	return "direc"
}

func (flag *Direction) Values() []Direction {
	return []Direction{Straight, Inverse, Random}
}

func (flag *Direction) String() string {
	return string(*flag)
}

func (flag *Direction) Set(value string) error {
	allowed := flag.Values()
	directionValue := Direction(value)
	for _, it := range allowed {
		if it == directionValue {
			*flag = directionValue
			return nil
		}
	}
	return errors.New(fmt.Sprintf("direction flag value must be one of: %v", allowed))
}

func (flag *Direction) Type() string {
	return "Direction"
}

func (flag *Direction) HelpMsg() string {
	return fmt.Sprintf("Cards direction. One of: %v", flag.Values())
}
