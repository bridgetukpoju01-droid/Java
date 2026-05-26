package main

func GenerateFont() map[rune][]string {
	font := make(map[rune][]string)

	for c := rune(32); c <= 126; c++ {
		lines := make([]string, 8)
		for row := 0; row < 8; row++ {
			line := ""
			for col := 0; col < 8; col++ {
				if c == ' ' {
					line += " "
				} else if row == col || row == int(c)%8 {
					line += string(c)
				} else {
					line += "."
				}
			}
			lines[row] = line
		}
		font[c] = lines
	}

	return font
}
