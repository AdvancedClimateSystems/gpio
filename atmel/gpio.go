package atmel

import (
	"fmt"
	"github.com/orangetux/gpio"
	"os"
	"path/filepath"
)

const (
	gpiobase   = "/sys/class/gpio"
	exportPath = "/sys/class/gpio/export"
)

// OpenPin exports the pin, creating the virtual files necessary for interacting with the pin.
// It also sets the mode for the pin, making it ready for use.
func OpenPin(n int, mode gpio.Mode) (gpio.Pinner, error) {
	// export this pin to create the virtual files on the system
	pinBase, err := expose(n)
	if err != nil {
		return nil, err
	}
	value, err := os.OpenFile(filepath.Join(pinBase, "value"), os.O_RDWR, 0600)
	if err != nil {
		return nil, err
	}
	p := &gpio.Pin{
		Number:    n,
		ModePath:  filepath.Join(pinBase, "direction"),
		EdgePath:  filepath.Join(pinBase, "edge"),
		ValueFile: value,
		Initial:   true,
	}
	p.SetMode(mode)
	if p.Err() != nil {
		p.Close()
		return nil, p.Err()
	}
	return p, nil
}

func expose(pin int) (string, error) {
	var pinController string
	pinNumber := pin % 32
	var err error

	switch pin / 32 {
	case 1:
		pinController = "A"
	case 2:
		pinController = "B"
	case 3:
		pinController = "C"
	case 4:
		pinController = "D"
	case 5:
		pinController = "E"
	default:
		err = fmt.Errorf("GPIO %d does not exist", pin)
	}

	pinBase := filepath.Join(gpiobase, fmt.Sprintf("pio%s%d", pinController, pinNumber))

	if _, statErr := os.Stat(pinBase); os.IsNotExist(statErr) {
		err = gpio.WriteFile(filepath.Join(gpiobase, "export"), "%d", pin)
	}
	return pinBase, err
}
