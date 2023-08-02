package helper

import "fmt"

func HtmlColorTag(text string, r, g, b uint8) string {
	return fmt.Sprintf(`<nobr style="color:rgb(%d,%d,%d)">%s</nobr>`, r, g, b, text)
}

func RedTag(text string) string {
	return HtmlColorTag(text, 252, 98, 98)
}

func GreenTag(text string) string {
	return HtmlColorTag(text, 82, 184, 64)
}

func OrangeTag(text string) string {
	return HtmlColorTag(text, 232, 151, 63)
}
