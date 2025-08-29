package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"math"
	"os"
)

func get_brightness(r int, g int, b int) int {
	dr := float64(r) / 256
	dg := float64(g) / 256
	db := float64(b) / 256
	red := 0.299 * dr * dr
	green := 0.587 * dg * dg
	blue := 0.114 * db * db
	bright := math.Sqrt(red + green + blue)
	return int(bright) >> 8
}

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

	/* 	// Decode image dimensions and format
	   	cfg, format, err := image.DecodeConfig(reader)
	   	if err != nil {
	   		log.Fatal(err)
	   	}

	   	// Rewind to start of file
	   	_, err = reader.Seek(0, io.SeekStart)
	   	if err != nil {
	   		log.Fatal(err)
	   	} */

	// Decode image data
	img, _, err := image.Decode(reader)
	if err != nil {
		log.Fatal(err)
	}
	bounds := img.Bounds()

	img_width := bounds.Max.X
	img_height := bounds.Max.Y

	fmt.Printf("Width: %d Height: %d", img_width, img_height)

	brightness_map := make([][]int, img_width)
	for i := range brightness_map {
		brightness_map[i] = make([]int, img_height)
	}

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()
			red := int(r << 8)
			green := int(g << 8)
			blue := int(b << 8)
			brightness := get_brightness(red, green, blue)
			brightness_map[x][y] = brightness
		}
	}

	for i := range brightness_map {
		for j := range brightness_map[i] {
			fmt.Printf("Brightness at (%d, %d): %d\n", i, j, brightness_map[i][j])
		}
	}

}
