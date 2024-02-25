package main

import (
	"fmt"
	"log"
	"math"
)

func GenDisplaceFn(acel, v_o, s_o float64) func(float64) float64 {
	disp := func(time float64) float64 {
		return ((0.5*acel)*(math.Pow(time, 2)) + (v_o * time) + s_o)
	}
	return disp
}

func main() {
	vars := make([]float64, 4, 4)
	str_val := []string{"acceleration", "initial velocity", "initial displacement", "time"}
	for i, v := range str_val {
		fmt.Printf("Enter a value for %s\n", v)
		_, err := fmt.Scan(&vars[i])
		if err != nil {
			log.Fatal(err)
		}
	}
	var a, v, s, t = vars[0], vars[1], vars[2], vars[3]

	disp := GenDisplaceFn(a, v, s)
	fmt.Println("The displacement is:  \n", disp(t))
}
