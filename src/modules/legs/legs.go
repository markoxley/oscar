package legs

import (
	"gobot.io/x/gobot/drivers/gpio"
)

const (
	jointCount = 3
	legCount   = 4
)

const (
	FrontLeft = iota
	FrontRight
	RearRight
	RearLeft
)

type Leg struct {
	joints []*gpio.ServoDriver
}
