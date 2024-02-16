package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"log"
	"os"
)

func main() {
	// parse commandline args
	percent := flag.Float64("p", 0.05, "percent for automatic bounds")
	sigma := flag.Float64("s", 1.00, "percent for automatic bounds")
	inp := flag.String("inp", "", "input file name")
	out := flag.String("out", "out", "output file name")

	// did not deliver file
	if len(*inp) == 0 {
		log.Fatal("You didn't supply a file goof")
	}

	fmt.Println("reading input file")
	inp_file, err := os.Open(*inp)
	if err != nil {
		log.Fatal(err)
	}

	out_file, err := os.Create(*out)
	if err != nil {
		log.Fatal(err)
	}

	defer inp_file.Close()
	defer out_file.Close()

	fmt.Println("parsing input file")
	inp_img, _, err := image.Decode(inp_file)
	if err != nil {
		log.Fatal(err)
	}

	out_img := canny(inp_img, *sigma, *percent)

	png.Encode(out_file, out_img)
}

func canny(img image.Image, sigma, percent float64) image.Image {
	bounds := img.Bounds()
	peaks := image.NewGray(bounds)
	mag := image.NewGray(bounds)
	outimg := image.NewGray(bounds)
	mr := (int)(sigma * 3)

	gauss := make([][]float64, (mr*2 + 1))
	rows := make([]float64, (mr*2+1)*(mr*2+1))

	for i := 0; i < 2*mr+1; i++ {
		gauss[i] = rows[i*(2*mr+1) : (i+1)*(2*mr+1)]
	}

	return nil
}
