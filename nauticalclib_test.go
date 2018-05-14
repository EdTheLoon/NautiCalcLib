package nauticalclib

import (
	"fmt"
	"testing"
)

func printRes(label string, exp interface{}, act interface{}) {
	fmt.Printf("%s\tExpected: %v\tActual: %v\n", label, exp, act)
}

func TestGyro(*testing.T) {
	// Set up the gyro error calculations
	var ge GyroError
	ge.Latitude = 34.495
	ge.LatDir = "S"
	ge.Declination = 19.37166667
	ge.DeclDir = "S"
	ge.LHA = 285.67
	ge.Gyro = 97.5

	// Calculate the gyro error
	ge.Calculate()

	// Start outputting results
	printRes("A", 0.19276168859404352, ge.A)
	printRes("A dir", "N", ge.ADir)
	printRes("B", 0.3651720926190195, ge.B)
	printRes("B dir", "S", ge.BDir)
	printRes("C", 0.17241040402497598, ge.C)
	printRes("C dir", "S", ge.CDir)
	printRes("A3", 81.91261495646594, ge.A3)
	printRes("Azimuth", 98.08738504, ge.Azimuth)
	printRes("Azi. Dir", "E", ge.AzimuthDir)
	printRes("Error", 0.5873850435340557, ge.GyroErr)
	printRes("Err Dir", "E", ge.ErrDir)
}
