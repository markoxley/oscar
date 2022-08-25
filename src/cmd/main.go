package main

import (
	"log"

	"github.com/robotmox/oscar/modules/camera"
	"github.com/robotmox/oscar/modules/common"
	"github.com/robotmox/oscar/modules/control"
	"github.com/robotmox/oscar/modules/imu"
	"github.com/robotmox/oscar/modules/legs"
	"github.com/robotmox/oscar/modules/ultrasonic"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/platforms/raspi"
)

func main() {
	rpi := raspi.NewAdaptor()
	channels := make(map[string]*chan common.Messager)
	master := gobot.NewMaster()

	lr, lc, err := legs.New(rpi)
	if err != nil {
		log.Fatalf("unable to initialise legs: %s", err.Error())
	}
	channels[legs.ChannelName] = lc
	master.AddRobot(lr)

	cr, cc, err := camera.New(rpi)
	if err != nil {
		log.Fatalf("unable to initialise camera: %s", err.Error())
	}
	channels[camera.ChannelName] = cc
	master.AddRobot(cr)

	ir, ic, err := imu.New(rpi)
	if err != nil {
		log.Fatalf("unable to initialise IMU: %s", err.Error())
	}
	channels[imu.ChannelName] = ic
	master.AddRobot(ir)

	ur, uc, err := ultrasonic.New(rpi)
	if err != nil {
		log.Fatalf("unable to initialise ultrasonic: %s", err.Error())
	}
	channels[ultrasonic.ChannelName] = uc
	master.AddRobot(ur)

	cnr, cnc, err := control.New(rpi)
	if err != nil {
		log.Fatalf("unable to initialise control: %s", err.Error())
	}
	channels[ultrasonic.ChannelName] = cnc
	master.AddRobot(cnr)

	err = master.Start()
	if err != nil {
		log.Fatalf("unable to execute master: %s", err.Error())
	}
}
