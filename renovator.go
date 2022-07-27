// This go program takes an image as input and outputs a base16 color theme.
// The output should be a string of the form:
// #RRGGBB
// where each pair of hex digits is a color component.
// The input image should be a PNG.
// There should be 16 colors grabbed or generated from the image.
// The colors should be chosen in a way that is visually pleasing.
// The colors should be chosen so that the most important colors are used.

package main


import (
	"image/color"
	"image/png"
	"os"
	"fmt"
)


func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	bounds := img.Bounds()
	width := bounds.Max.X
	height := bounds.Max.Y

	var colors [16]color.Color

	for i := 0; i < 16; i++ {
		colors[i] = img.At(i * width / 16, height / 2)
	}


	for i := 0; i < 16; i++ {
		r, g, b, _ := colors[i].RGBA()
		fmt.Printf("#%02x%02x%02x\n", r>>8, g>>8, b>>8)
	}


	// print hex colors as true color in the terminal
	for i := 0; i < 16; i++ {
		r, g, b, _ := colors[i].RGBA()
		fmt.Printf("\x1b[48;2;%d;%d;%dm  \x1b[0m", r>>8, g>>8, b>>8)
	}
	fmt.Println()
}
