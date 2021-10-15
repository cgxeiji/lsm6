package lsm6

import (
	"fmt"

	"github.com/cgxeiji/serial"
)

// Device defines a LSM6 device.
type Device struct {
	i2c *serial.I2C

	accScale  float64
	gyroScale float64
}

// New returns a new LSM6 device.
func New(bus string, addr uint16) (*Device, error) {
	if addr == 0 {
		addr = Addr
	}

	i2c, err := serial.NewI2C(bus, addr)
	if err != nil {
		return nil, fmt.Errorf("lsm6: could not initialize I2C: %w", err)
	}

	d := &Device{
		i2c: i2c,
	}

	if err := d.i2c.Write(Ctrl1XL, 0x80); err != nil {
		return nil, fmt.Errorf("lsm6: could not configure device: %w", err)
	}
	d.accScale = 2

	if err := d.i2c.Write(Ctrl2G, 0x80); err != nil {
		return nil, fmt.Errorf("lsm6: could not configure device: %w", err)
	}
	d.gyroScale = 245

	if err := d.i2c.Write(Ctrl3C, 0x04); err != nil {
		return nil, fmt.Errorf("lsm6: could not configure device: %w", err)
	}

	return d, nil
}

// Close closes the device and cleans after itself.
func (d *Device) Close() {
	d.i2c.Close()
}

// Acc returns the values of the accelerometer.
func (d *Device) Acc() (x, y, z float64, err error) {
	x, y, z, err = d.readPkg6(OutXLXL)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("lsm6: could not read accelerometer: %w", err)
	}

	x *= d.accScale
	y *= d.accScale
	z *= d.accScale

	return x, y, z, nil
}

// Gyro returs the values of the gyroscope.
func (d *Device) Gyro() (x, y, z float64, err error) {
	x, y, z, err = d.readPkg6(OutXLG)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("lsm6: could not read gyroscope: %w", err)
	}

	x *= d.gyroScale
	y *= d.gyroScale
	z *= d.gyroScale

	return x, y, z, nil
}

func (d *Device) readPkg6(reg byte) (x, y, z float64, err error) {
	b, err := d.i2c.ReadBytes(reg, 6)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("lsm6: could not read %#x: %w", reg, err)
	}

	x = float64(
		int16(b[1])<<8|
			int16(b[0])) / halfADC
	y = float64(
		int16(b[3])<<8|
			int16(b[2])) / halfADC
	z = float64(
		int16(b[5])<<8|
			int16(b[4])) / halfADC

	return x, y, z, nil
}
