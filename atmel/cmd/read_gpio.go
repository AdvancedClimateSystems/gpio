package main

import (
	"flag"
	"fmt"
	"github.com/advancedclimatesystems/gpio/atmel"
	"github.com/orangetux/gpio"
	"os"
	"strconv"
)

func usage() {
	fmt.Printf("Read state of GPIO pin.\n\nUsage: %s pin_number\n", os.Args[0])
	flag.PrintDefaults()
}
func main() {
	flag.Usage = usage
	flag.Parse()

	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	number, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Printf("%s is not a valid number.\n", flag.Arg(0))
		os.Exit(1)
	}
	p, err := atmel.OpenPin(number, gpio.ModeInput)
	if err != nil {
		fmt.Printf("Can't read pin: %s\n", err)
		os.Exit(1)
	}
	defer p.Close()

	fmt.Println(p.Get())
}
