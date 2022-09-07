package drivers

import (
	"errors"
	"math"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
)

type HCSR04 struct {
	pin        string
	name       string
	connection gpio.DigitalWriter
	distance   float64
	echo       string
	timeout    uint16
	gobot.Commander
}

func NewHCSR04Driver(a gpio.DigitalWriter, trigger string, echo string, timeout uint16) *HCSR04 {
	h := &HCSR04{
		name:       gobot.DefaultName("HCSR04"),
		pin:        trigger,
		echo:       echo,
		connection: a,
		distance:   math.MaxFloat64,
		Commander:  gobot.NewCommander(),
		timeout:    timeout,
	}

	h.AddCommand("Measure", func(params map[string]interface{}) interface{} {
		return h.Measure()
	})

	return h
}

// Start implements the Driver interface
func (h *HCSR04) Start() (err error) { return }

// Halt implements the Driver interface
func (h *HCSR04) Halt() (err error) { return }

// Name returns the LedDrivers name
func (h *HCSR04) Name() string { return h.name }

// SetName sets the LedDrivers name
func (h *HCSR04) SetName(n string) { h.name = n }

// Pin returns the LedDrivers name
func (h *HCSR04) Pin() string { return h.pin }

// Pin returns the LedDrivers name
func (h *HCSR04) Echo() string { return h.echo }

// Connection returns the HCSR04 Connection
func (h *HCSR04) Connection() gobot.Connection {
	return h.connection.(gobot.Connection)
}

func (h *HCSR04) Distance() float64 {
	return h.distance
}

func (h *HCSR04) Measure() (dist float64) {
	t := time.Now()
	if err := h.triggerSet(false); err != nil {
		return
	}
	time.Sleep(2 * time.Microsecond)
	if err := h.triggerSet(true); err != nil {
		return
	}
	time.Sleep(10 * time.Microsecond)
	if err := h.triggerSet(false); err != nil {
		return
	}
	i := uint8(0)
	for {

		state, err := h.echoRead()
		if err != nil {
			return
		}
		if state {
			t = time.Now()
			break
		}
		i++
		if i > 10 {
			if time.Since(t).Microseconds() > int64(h.timeout) {
				return
			}
			i = 0
		}
	}
	i = 0
	for {
		state, err := h.echoRead()
		if err != nil {
			return
		}
		if !state {
			h.distance = float64(time.Since(t).Microseconds())
			return h.distance
		}
		i++
		if i > 10 {
			if time.Since(t).Microseconds() > int64(h.timeout) {
				return
			}
			i = 0
		}
	}
}

func (h *HCSR04) triggerSet(high bool) (err error) {
	var value byte
	if high {
		value = 1
	}
	if writer, ok := h.Connection().(gpio.DigitalWriter); ok {
		err = writer.DigitalWrite(h.Pin(), value)
		if err != nil {
			h.distance = 0
		}
		return
	}
	h.distance = 0
	return errors.New("unable to set trigger")
}

func (h *HCSR04) echoRead() (value bool, err error) {
	if reader, ok := h.Connection().(gpio.DigitalReader); ok {
		var val int
		val, err = reader.DigitalRead(h.Echo())
		if err != nil {
			h.distance = 0
		}
		value = val > 0
	}
	return
}
