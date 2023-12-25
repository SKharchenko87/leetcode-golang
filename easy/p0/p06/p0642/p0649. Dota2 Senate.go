package main

import (
	"fmt"
	"strings"
)

func main() {
	//senate := "D"
	//senate := "DRRD"
	//senate := "DR"
	//senate := "RDD"
	//senate := "DRRDDR"
	//senate := "DRRDDDDDDR"
	//senate := "DRRDRRRRDR"
	//senate := "D R R RRRR"
	//senate := "RR"
	senate := "DDRRR"
	fmt.Println(predictPartyVictory(senate))

	//fmt.Println(predictPartyVictory(predictPartyVictory(senate)))
	//
	//fmt.Println(predictPartyVictory(predictPartyVictory(predictPartyVictory(senate))))
}

func predictPartyVictory(senate string) string {
	if recurs(senate) {
		return "Radiant"
	} else {
		return "Dire"
	}
}

func lastVoiceRound(str string, sb strings.Builder, senatorNotVoice int, res string, x int32) string {
	for _, v := range str {
		if v == x && senatorNotVoice > 0 {
			senatorNotVoice--
		} else {
			sb.WriteString(string(v))
		}
	}
	if senatorNotVoice > 0 {
		return res
	}
	return sb.String()
}
func recurs(senate string) bool {
	if len(senate) == 1 {
		if senate[0] == 'R' {
			return true
		} else {
			return false
		}
	}
	senateD, senateDVoice := 0, 0
	senateR, senateRVoice := 0, 0
	sb := strings.Builder{}

	for _, v := range senate {
		if v == 'D' {
			senateD++
			if senateD > senateR {
				senateDVoice++
				sb.WriteString("D")
			}
		} else {
			senateR++
			if senateR > senateD {
				senateRVoice++
				sb.WriteString("R")
			}
		}
	}
	str := sb.String()
	sd := senateDVoice - (senateR - senateRVoice) // Сколько сенаторов не голосовало у D
	sr := senateRVoice - (senateD - senateDVoice) // Сколько сенаторов не голосовало у R

	sb.Reset()
	if sd > 0 {
		str = lastVoiceRound(str, sb, sd, "D", 'R')
	} else if sr > 0 {
		str = lastVoiceRound(str, sb, sr, "R", 'D')
	}

	return recurs(str)
}
