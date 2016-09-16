# GPIO

GPIO packag with support for [Atmel GPIO Controller][atmel]. This controller
is used by many SoC's such as the Atmel AT91SAM9x series. The package is only a
small wrapper around [orangetux/gpio][gpio].

## Usage

```go
package main

import (
    "fmt"

    "github.com/orangetux/gpio"
    "github.com/advancedclimatesystems/gpio/atmel"
)

func main() {
    pin, err := atmel.OpenPin(115, gpio.ModeInput)
    if err != nil {
        fmt.Printf("Error opening pin %s.\n", err)
        return
    }
    defer pin.Close()

    fmt.Printf("State of pin %v.\n", pin.Get())
}
```

## License

GPIO is licensed under [Mozilla Public License][mpl] © 2016 [Advanced Climate
System][acs].


[acs]: http://advancedclimate.nl
[mpl]: LICENSE
[atmel]: https://www.kernel.org/doc/Documentation/devicetree/bindings/gpio/gpio_atmel.txt
[gpio]: https://github.com/orangetux/gpio
