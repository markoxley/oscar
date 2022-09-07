package legs

import (
	"fmt"
	"time"

	"github.com/robotmox/oscar/modules/common"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
)

const ChannelName = "Legs"

func New(rpi *raspi.Adaptor) (*gobot.Robot, *chan common.Messager, error) {
	pca9685 := i2c.NewPCA9685Driver(rpi)
	devices := make([]gobot.Device, 1+(legCount*jointCount))
	devices[0] = pca9685
	legs := make([]*Leg, legCount)
	jId := 0
	for l := range legs {
		legs[l].joints = make([]*gpio.ServoDriver, jointCount)
		for j := range legs[l].joints {
			legs[j].joints[j] = gpio.NewServoDriver(pca9685, fmt.Sprintf("%v", jId))
			devices[jId+1] = legs[j].joints[j]
		}
	}

	work := func() {
		pca9685.SetPWMFreq(60)
		gobot.Every(1*time.Millisecond, func() {

		})
	}
	robot := gobot.NewRobot("legs",
		[]gobot.Connection{rpi},
		devices,
		work,
	)
	c := make(chan common.Messager)
	return robot, &c, nil
}
