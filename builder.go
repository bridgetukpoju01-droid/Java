package main

import "strings"

type part struct {
	text  string
	style string
}

type ArtBuilder struct {
	parts []part
	last  int
}

func NewArtBuilder() *ArtBuilder {
	return &ArtBuilder{last: -1}
}

func (a *ArtBuilder) AddText(text string) *ArtBuilder {
	a.parts = append(a.parts, part{text: text, style: "normal"})
	a.last = len(a.parts) - 1
	return a
}

func (w *ArtBuilder) SetStyle(style string) *ArtBuilder {
	if style != "normal" && style != "bold" && style != "italic" && style != "outline" {
		panic("invalid style")
	}

	if w.last >= 0 {
		w.parts[w.last].style = style
	}
	return w
}

func (a *ArtBuilder) Build() string {
	lines := make([]string, 8)

	for _, p := range a.parts {
		text := p.text

		switch p.style {
		case "bold":
			text = "##" + text + "##"
		case "italic":
			text = "/" + text
		case "outline":
			text = "|" + text + "|"
		}

		for i := 0; i < 8; i++ {
			lines[i] += text
		}
	}

	return strings.Join(lines, "\n") + "\n"
}
