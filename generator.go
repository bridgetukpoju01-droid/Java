package main

func GeneratePattern(input rune) []string {
	if input < 'A' || input > 'Z' {
		return []string{}
	}

	fontMap := map[rune][]string{
		'A': {
			"  ##  ",
			" #  # ",
			" #  # ",
			" #### ",
			" #  # ",
			" #  # ",
			" #  # ",
			"      ",
		},
		'Z': {
			" #### ",
			"    # ",
			"   #  ",
			"  #   ",
			" #    ",
			" #    ",
			" #### ",
			"      ",
		},
	}

	if pattern, exists := fontMap[input]; exists {
		return pattern
	}

	defaultPattern := make([]string, 8)
	for i := 0; i < 8; i++ {
		if i == 0 || i == 6 {
			defaultPattern[i] = " #### "
		} else if i == 7 {
			defaultPattern[i] = "      "
		} else {
			defaultPattern[i] = " #  # "
		}
	}
	return defaultPattern
}
