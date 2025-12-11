package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var content string

type Device struct {
	Name string
	Outs []*Device
}

type Key struct {
	Name   string
	HasDAC bool
	HasFFT bool
}

func main() {
	lines := strings.Split(content, "\n")

	sum := 0

	devices := make(map[string]*Device)
	devices["out"] = &Device{"out", make([]*Device, 0)}
	for _, line := range lines {
		parts := strings.Split(line, ": ")

		device := parts[0]
		outs := make([]*Device, 0)
		devices[device] = &Device{device, outs}
	}

	for _, line := range lines {
		parts := strings.Split(line, ": ")
		current := devices[parts[0]]

		outs := make([]*Device, 0)
		for _, device := range strings.Split(parts[1], " ") {
			newDevice := devices[device]

			outs = append(outs, newDevice)
		}
		current.Outs = outs
	}

	memo := map[Key]int{}
	sum = dfs(devices["svr"], false, false, memo)
	fmt.Println(sum)
}

func dfs(device *Device, hasDAC, hasFFT bool, memo map[Key]int) int {
	key := Key{device.Name, hasDAC, hasFFT}

	if v, ok := memo[key]; ok {
		return v
	}

	if device.Name == "out" {
		if hasDAC && hasFFT {
			memo[key] = 1
			return 1
		}
		memo[key] = 0
		return 0
	}

	total := 0
	for _, conn := range device.Outs {
		total += dfs(conn, hasDAC || conn.Name == "dac", hasFFT || conn.Name == "fft", memo)
	}

	memo[key] = total
	return total
}
