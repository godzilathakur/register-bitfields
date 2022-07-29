package builder

import "log"

func WhiteSpace(tabWidth int, numTabs int) string {
	if tabWidth < 0 || numTabs < 0 {
		log.Fatalln("Invalid tab config")
	}
	result := ""
	for i := 0; i < tabWidth*numTabs; i += 2 {
		result += "  "
	}
	return result
}
