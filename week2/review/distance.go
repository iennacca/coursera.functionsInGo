package main

import (
	"fmt"
)

func GenDisplaceFn(initS0 float64, vel float64, accel float64) func(float64) float64{
	var getDisplacement = func (t float64) float64{
		return ((0.5*accel*t*t) + vel*t + initS0)
	}
	
	return getDisplacement
}

func main(){
	var s0, vel, accel, t float64
	fmt.Printf("Please enter the value for initial distance, velocity, acceleartion and time: ");
	_ , err := fmt.Scanf("%f %f %f %f",&s0, &vel, &accel, &t)
	if err != nil {
		fmt.Println(err)
	}

	var getTheDistanceFromTime = GenDisplaceFn(s0, vel, accel)
	fmt.Printf("%f",(getTheDistanceFromTime(t)))
}