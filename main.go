package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"os"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Please provide a valid file path to an image.")
	}

	reader, err := os.Open(args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer reader.Close()

	// Decode image dimensions and format
	cfg, format, err := image.DecodeConfig(reader)
	if err != nil {
		log.Fatal(err)
	}

	// Rewind to start of file
	_, err = reader.Seek(0, io.SeekStart)
	if err != nil {
		log.Fatal(err)
	}

	// Decode image data
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()

	fmt.Println("Width:", cfg.Width, "Height:", cfg.Height, "Format:", format)

	var histogram [16][4]int
	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, a := img.At(x, y).RGBA()
			// A color's RGBA method returns values in the range [0, 65535].
			// Shifting by 12 reduces this to the range [0, 15].
			histogram[r>>12][0]++
			histogram[g>>12][1]++
			histogram[b>>12][2]++
			histogram[a>>12][3]++
		}
	}

	// Print the results.
	fmt.Printf("%-14s %6s %6s %6s %6s\n", "bin", "red", "green", "blue", "alpha")
	for i, x := range histogram {
		fmt.Printf("0x%04x-0x%04x: %6d %6d %6d %6d\n", i<<12, (i+1)<<12-1, x[0], x[1], x[2], x[3])
	}
}
