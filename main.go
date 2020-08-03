package main

import (
	"fmt"
	"os"
)

func main() {
	fgColorFmt := "38;5;%d"
	bgColorFmt := "48;5;%d"

	// 8-bit colorの設定
	fg := fmt.Sprintf(fgColorFmt, 242)
	bg := fmt.Sprintf(bgColorFmt, 208)

	text := "hoge"
	colorizeText := fmt.Sprintf("\x1b[%s;%sm%s", fg, bg, text)

	fmt.Fprintf(os.Stdout, colorizeText)
}
