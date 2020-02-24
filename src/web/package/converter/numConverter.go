package converter

import (
	"strconv"
)

func ExtractNumber(target string) string {
	var output string = ""
	for _, runeCharacter := range target {
		c := string(runeCharacter)
		if  c == "0" || c == "０" {
			output += c
		} else {
			num, _ := strconv.Atoi(c)
			if num == 0 {
				continue
			}
			output += c
		}
	}
	return output
}