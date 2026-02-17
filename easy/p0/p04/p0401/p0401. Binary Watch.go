package p0401

import (
	"fmt"
	"math/bits"
)

func readBinaryWatch(turnedOn int) []string {
	res := []string{}
	for i := uint(0); i < 0b1111111111; i++ {
		minute := i & 0b0000111111
		hour := i >> 6
		if minute < 60 && hour < 12 && bits.OnesCount(i) == turnedOn {
			res = append(res, fmt.Sprintf("%d:%02d", hour, minute))
		}
	}
	return res
}

func readBinaryWatch0(turnedOn int) []string {
	if turnedOn > 8 {
		return []string{}
	} else if turnedOn == 0 {
		return []string{"0:00"}
	}

	res := []string{}
	for hour := uint(0); hour < 12; hour++ {
		for minute := uint(0); minute < 60; minute++ {
			if bits.OnesCount(hour)+bits.OnesCount(minute) == turnedOn {
				res = append(res, fmt.Sprintf("%d:%02d", hour, minute))
			}
		}
	}

	return res
}
