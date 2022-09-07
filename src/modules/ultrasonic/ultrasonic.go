package ultrasonic

import (
	"fmt"
	"time"

	"github.com/robotmox/oscar/drivers"
	"github.com/robotmox/oscar/modules/common"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

const (
	ChannelName    = "Ultrasonic"
	Left           = 0
	Right          = 1
	LeftTrigger    = "10"
	LeftEcho       = "11"
	RightTrigger   = "12"
	RightEcho      = "13"
	DefaultTimeout = 500
)

func New(rpi *raspi.Adaptor) (*gobot.Robot, *chan common.Messager, error) {
	devices := make([]*drivers.HCSR04, 2)
	devices[Left] = drivers.NewHCSR04Driver(rpi, LeftTrigger, LeftEcho, DefaultTimeout)
	devices[Right] = drivers.NewHCSR04Driver(rpi, RightTrigger, RightEcho, DefaultTimeout)
	c := make(chan common.Messager)
	work := func() {
		var distances [2]float64
		for {
			for i, d := range devices {
				distances[i] = d.Measure()
			}
			c <- common.NewMessage("whiskers", fmt.Sprintf("%v,%v", distances[0], distances[1]))
			time.Sleep(time.Millisecond)
		}

	}
	robot := gobot.NewRobot("whiskers",
		[]gobot.Connection{rpi},
		devices,
		work,
	)

	return robot, &c, nil
}
