package elon

import "fmt"

func (c *Car) Drive() {
	// If there is not enough battery power left we
	// can no longer drive.
	if c.battery < c.batteryDrain {
		return
	}

	c.battery -= c.batteryDrain
	c.distance += c.speed
}

func (c *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", c.distance)
}

func (c *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", c.battery)
}

func (c *Car) CanFinish(trackDistance int) bool {
	// This will truncate as integer division which is what we want
	// as this car drives in discrete steps.
	numTimes := c.battery / c.batteryDrain
	return c.speed*numTimes >= trackDistance
}
