package main

import (
	"fmt"
	"io"
	"net"
	"os"
)
import "os/exec"

func main() {
	fmt.Println("Hello")

	cmd := exec.Command("raspivid", "-t", "120000", "-l", "-o", "tcp://0.0.0.0:3333", "-w", "400", "-h", "300")

	err := cmd.Start()
	if err != nil {
		panic(err)
	}

	conn, err := net.Dial("tcp", "raspberrypi:3333")
	defer conn.Close()
	if err != nil {
		panic(err)
	}
	f, err := os.Create("stream.h264")
	defer f.Close()
	if err != nil {
		panic(err)
	}
	err = cmd.Wait()
	if err != nil {
		panic(err)
	}
	written, err := io.Copy(f, conn)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Wrote %d\n", written)
}
