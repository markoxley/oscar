package camera

import (
	"github.com/robotmox/oscar/modules/common"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

const ChannelName = "Camera"

func New(rpi *raspi.Adaptor) (*gobot.Robot, *chan common.Messager, error) {
	c := make(chan common.Messager)
	return nil, &c, nil
}
