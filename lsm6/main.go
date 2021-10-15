package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/cgxeiji/lsm6"
)

func main() {
	imu, err := lsm6.New("", 0)
	if err != nil {
		log.Fatal(err)
	}
	defer imu.Close()

	done := make(chan struct{})
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Printf("\n\n")
		for {
			select {
			case <-done:
				return
			default:
			}
			ax, ay, az, err := imu.Acc()
			if err != nil {
				log.Fatal(err)
			}

			gx, gy, gz, err := imu.Gyro()
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("\033[2F")
			fmt.Printf("ax, ay, az = %4.2f, %4.2f, %4.2f          \n", ax, ay, az)
			fmt.Printf("gx, gy, gz = %4.2f, %4.2f, %4.2f          \n", gx, gy, gz)
		}
	}()

	bufio.NewReader(os.Stdin).ReadString('\n')
	close(done)
	wg.Wait()
}
