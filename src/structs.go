package main

import "fmt"

type car struct {
	gas_pedal       uint16
	break_pedal     uint16
	streering_wheel int16
	top_speed_kmh   float64
}

func main() {
	a_car := car{
		gas_pedal:       22314,
		break_pedal:     0,
		streering_wheel: 12561,
		top_speed_kmh:   225.0}
	fmt.Println(a_car.gas_pedal)
}
