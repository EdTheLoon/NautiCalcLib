// NautiCalc (nauticalc.go)
// Copyright 11/05/2017
// Author: Adrian Reid

// A tool for Nautical calculations
package nauticalclib

import (
	"math"
)

// INTERNAL FUNCTIONS

// Determines whether the passed in float is East or West then removes any negative
// sign. Returns a float and string
func getDirectionEW(v float64) (float64, string) {
	var d string
	if v < 0 {
		d = "W"
		v = v - (v * 2)
	} else {
		d = "E"
	}
	return v, d
}

// Determines whether the passed in float is North or South then removes any negative
// sign. Returns a float and string
func getDirectionNS(v float64) (float64, string) {
	var d string
	if v < 0 {
		d = "S"
		v = v - (v * 2)
	} else {
		d = "N"
	}
	return v, d
}

// Converts north into south and vice versa
func opposite(d string) string {
	if d == "N" {
		d = "S"
		return d
	} else if d == "S" {
		d = "N"
		return d
	} else if d == "E" {
		d = "W"
		return d
	} else if d == "W" {
		d = "E"
		return d
	} else {
		return "error"
	}
}

// Converts a value from radians into degrees
func degs(r float64) float64 {
	d := r * 180 / math.Pi
	return d
}

// Converts a value from degrees into radians
func rads(d float64) float64 {
	r := d * math.Pi / 180
	return r
}

// STRUCTS

type CompassError struct {
	Magnetic  float64
	Gyro      float64
	Variation float64
	VarDir    string
	Deviation float64
	DevDir    string
	Corrected float64
	ComErr    float64
	ErrDir    string
}

type GyroError struct {
	Gyro        float64
	Latitude    float64
	LatDir      string
	LHA         float64
	Declination float64
	DeclDir     string
	A           float64
	ADir        string
	B           float64
	BDir        string
	C           float64
	CDir        string
	A3          float64
	Azimuth     float64
	AzimuthDir  string
	GyroErr     float64
	ErrDir      string
}

// METHODS

// Gets user input, calculates the compass error and deviation and then prints it
func (c *CompassError) Calculate() {
	// Start calculating
	if c.VarDir == "W" {
		c.Corrected = c.Gyro + c.Variation
	} else {
		c.Corrected = c.Gyro - c.Variation
	}
	c.Deviation = c.Corrected - c.Magnetic
	c.ComErr = c.Magnetic - c.Gyro

	// Determine east or west and convert negative number into positive number
	c.Deviation, c.DevDir = getDirectionEW(c.Deviation)
	c.ComErr, c.ErrDir = getDirectionEW(c.ComErr)
}

func (g *GyroError) Calculate() {
	// Start calculating A
	g.A = math.Abs(degs(math.Tan(rads(g.Latitude))) / degs(math.Tan(rads(g.LHA))))
	if g.LHA > 90.0 && g.LHA < 270.0 {
		g.ADir = g.LatDir
	} else {
		g.ADir = opposite(g.LatDir)
	}

	// Start calculating B
	g.B = math.Abs(degs(math.Tan(rads(g.Declination))) / degs(math.Sin(rads(g.LHA))))
	g.BDir = g.DeclDir

	// Start calculating C
	if g.ADir == g.BDir {
		g.C = g.A + g.B
		g.CDir = g.ADir
	} else {
		g.C = math.Abs(g.A - g.B)
		highest := math.Max(g.A, g.B)
		if highest == g.A {
			g.CDir = g.ADir
		} else {
			g.CDir = g.BDir
		}
	}

	// Start calculating g.A3
	cCosLat := rads(g.C) * math.Cos(rads(g.Latitude))
	g.A3 = rads(1) / cCosLat
	g.A3 = degs(math.Atan(g.A3))

	// Remove negative sign
	if g.A3 < 0 {
		g.A3 = g.A3 - (g.A3 * 2)
	}

	// East or west
	if g.LHA > 180.0 && g.LHA < 360.0 {
		g.AzimuthDir = "E"
	} else {
		g.AzimuthDir = "W"
	}

	// True bearing
	if g.CDir == "N" {
		if g.AzimuthDir == "E" {
			g.Azimuth = 0 + g.A3
		} else {
			g.Azimuth = 360 - g.A3
		}
	} else {
		if g.AzimuthDir == "E" {
			g.Azimuth = 180 - g.A3
		} else {
			g.Azimuth = 180 + g.A3
		}
	}

	// Start calculating the error
	g.GyroErr = g.Azimuth - g.Gyro
	g.GyroErr, g.ErrDir = getDirectionEW(g.GyroErr)

}
