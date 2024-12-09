package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var content string

type Space struct {
	Free  bool
	Moved bool
	ID    int
	Size  int
}

func main() {
	numbers := strings.Split(content, "")

	memory := make([]int, 0)
	memory2 := make([]Space, 0)
	for i, n := range numbers {
		size, _ := strconv.Atoi(n)
		if i%2 == 0 {
			memory2 = append(memory2, Space{false, false, i / 2, size})
			for j := 0; j < size; j++ {
				memory = append(memory, i/2)
			}
		} else {
			memory2 = append(memory2, Space{true, false, -1, size})
			for j := 0; j < size; j++ {
				memory = append(memory, -1)
			}
		}
	}

	i := 0
	for {
		if i >= len(memory) {
			break
		}
		n := memory[i]

		if n == -1 {
			memory[i] = memory[len(memory)-1]
			memory = memory[:len(memory)-1]

			for memory[len(memory)-1] == -1 {
				memory = memory[:len(memory)-1]
			}

		}
		i++
	}

	i = len(memory2) - 1
	for {
		if i < 0 {
			break
		}
		space := memory2[i]
		if space.Free || space.Moved {
			i--
			continue
		}

		for j := 0; j < i; j++ {
			if memory2[j].Free && memory2[j].Size >= space.Size {
				freeSpace := memory2[j]
				freeSpace.Size = freeSpace.Size - space.Size
				space.Moved = true

				newMemory := make([]Space, 0)
				newMemory = append(newMemory, memory2[:j]...)
				newMemory = append(newMemory, space)
				newMemory = append(newMemory, freeSpace)
				newMemory = append(newMemory, memory2[j+1:i]...)
				newMemory = append(newMemory, Space{true, false, -1, space.Size})
				newMemory = append(newMemory, memory2[i+1:]...)
				memory2 = newMemory
				i++
				break
			}
		}
		i--
	}

	checksum := 0
	for i, n := range memory {
		checksum += n * i
	}

	flatMemory := make([]int, 0)

	for _, space := range memory2 {
		if space.Free {
			for j := 0; j < space.Size; j++ {
				flatMemory = append(flatMemory, -1)
			}
			continue
		}
		for j := 0; j < space.Size; j++ {
			flatMemory = append(flatMemory, space.ID)
		}
	}
	checksum2 := 0
	for i, n := range flatMemory {
		if n == -1 {
			continue
		}
		checksum2 += n * i
	}

	fmt.Println(checksum)
	fmt.Println(checksum2)
}
