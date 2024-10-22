package tempconv

import "fmt"

type Celsius float64
type Fahrenheit float64
type Kelvin float64

const (
	AbsZeroC = -273.15
	AbsZeroK = 0
)

func (c Celsius) String() string {return fmt.Sprintf("%f ˚C", c)}
func (f Fahrenheit) String() string {return fmt.Sprintf("%f ˚F", f)}
func (k Kelvin) String() string {return fmt.Sprintf("%f K", k)}
