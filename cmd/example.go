package main

import (
	"fmt"
	"image"
	"log"
	"os"

	_ "image/jpeg"

	"github.com/andybons/ascii"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s src.jpg\n", os.Args[0])
		os.Exit(2)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatal(err)
	}
	art := ascii.Thumbnail(img, 85, 85)
	fmt.Printf("%+v\n", img.Bounds())
	fmt.Printf("%s\n", art)
}
