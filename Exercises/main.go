package main

import (
	"./Exercise5"
)

func main(){
	driver.Init()
	driver.SetMotorDir(1)
	
	for {
		if driver.GetFloorSensor() == 2 {
			driver.SetMotorDir(0)
		}
	}

}