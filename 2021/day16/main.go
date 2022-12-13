package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	textFile, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := ioutil.ReadAll(textFile)
	data := string(byteValue)

	equivalences := map[string]string{
		"0": "0000",
		"1": "0001",
		"2": "0010",
		"3": "0011",
		"4": "0100",
		"5": "0101",
		"6": "0110",
		"7": "0111",
		"8": "1000",
		"9": "1001",
		"A": "1010",
		"B": "1011",
		"C": "1100",
		"D": "1101",
		"E": "1110",
		"F": "1111",
	}

	base := ""
	for _, c := range data {
		base += equivalences[string(c)]
	}

	packet := decodePacket(base)

	fmt.Println(sumVersions(packet))
	fmt.Println(evaluatePacket(packet))
}

func sumVersions(packet Packet) int {
	total := packet.version
	for _, p := range packet.packets {
		total += sumVersions(p)
	}
	return total
}

func evaluatePacket(packet Packet) int {
	var value int

	if packet.typeID == 0 {
		value = 0
		for _, p := range packet.packets {
			value += evaluatePacket(p)
		}
	}

	if packet.typeID == 1 {
		value = 1
		for _, p := range packet.packets {
			value *= evaluatePacket(p)
		}
	}

	if packet.typeID == 2 {
		value = -1
		for _, p := range packet.packets {
			packetValue := evaluatePacket(p)
			if packetValue < value || value == -1 {
				value = packetValue
			}
		}
	}

	if packet.typeID == 3 {
		value = -1
		for _, p := range packet.packets {
			packetValue := evaluatePacket(p)
			if packetValue > value || value == -1 {
				value = packetValue
			}
		}
	}

	if packet.typeID == 4 {
		value = packet.number
	}

	if packet.typeID == 5 {
		p1 := evaluatePacket(packet.packets[0])
		p2 := evaluatePacket(packet.packets[1])
		value = 0
		if p1 > p2 {
			value = 1
		}
	}

	if packet.typeID == 6 {
		p1 := evaluatePacket(packet.packets[0])
		p2 := evaluatePacket(packet.packets[1])
		value = 0
		if p1 < p2 {
			value = 1
		}
	}

	if packet.typeID == 7 {
		p1 := evaluatePacket(packet.packets[0])
		p2 := evaluatePacket(packet.packets[1])
		value = 0
		if p1 == p2 {
			value = 1
		}
	}

	return value
}

type Packet struct {
	version  int
	consumed int
	typeID   int
	packets  []Packet
	number   int
}

func decodePacket(bits string) Packet {
	version, _ := strconv.ParseInt(bits[0:3], 2, 64)
	typeID, _ := strconv.ParseInt(bits[3:6], 2, 64)

	bits = bits[6:]
	consumed := 6
	if typeID == 4 {
		content := ""
		for {
			prefix := bits[0:1]
			bit := bits[1:5]
			content += bit
			bits = bits[5:]
			consumed += 5
			if prefix == "0" {
				break
			}
		}
		number, _ := strconv.ParseInt(content, 2, 64)
		return Packet{version: int(version), consumed: consumed, typeID: int(typeID), packets: []Packet{}, number: int(number)}
	} else {
		lengthTypeID := bits[0:1]
		if lengthTypeID == "0" {
			data := bits[1:16]
			bits = bits[16:]
			consumed += 16
			subPacketsLength, _ := strconv.ParseInt(data, 2, 64)
			packets := decodePackets(bits[0:subPacketsLength])
			consumed += int(subPacketsLength)
			return Packet{version: int(version), consumed: consumed, typeID: int(typeID), packets: packets}
		} else if lengthTypeID == "1" {
			data := bits[1:12]
			bits = bits[12:]
			consumed += 12
			subPacketsNumber, _ := strconv.ParseInt(data, 2, 64)
			packets := make([]Packet, 0)
			for i := 0; i < int(subPacketsNumber); i++ {
				packet := decodePacket(bits)
				bits = bits[packet.consumed:]
				consumed += packet.consumed
				packets = append(packets, packet)
			}
			return Packet{version: int(version), consumed: consumed, typeID: int(typeID), packets: packets}
		}
	}
	return Packet{}
}

func decodePackets(bits string) []Packet {
	packets := make([]Packet, 0)
	for len(bits) > 0 {
		packet := decodePacket(bits)
		packets = append(packets, packet)
		bits = bits[packet.consumed:]
	}

	return packets
}
