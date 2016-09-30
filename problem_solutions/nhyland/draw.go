package main

import (
	"image"
	"image/color"
	"image/gif"
)

// a palette I stole online for the image.NewPaletted - required.
var palette = []color.Color{
	color.White,
	color.RGBA{250, 130, 0, 100},
	color.RGBA{0, 0, 250, 100},
	color.RGBA{0, 250, 0, 100},
	// color.RGBA{0xff, 0x00, 0x00, 0xff},
	// color.RGBA{0xff, 0x00, 0xff, 0xff},
	// color.RGBA{0xff, 0xff, 0x00, 0xff},
	// color.RGBA{0xff, 0xff, 0xff, 0xff},
}

// xy coordinates
type xy struct {
	x, y int
}

// squareParams are the parameters for square()
type squareParams struct {
	from, to xy
	c        color.Color
}

// square draws a square on a paletted image, taking in squareParams as a struct
func square(img *image.Paletted, p squareParams) {
	for x := p.from.x; x < p.to.x; x++ {
		for y := p.from.y; y < p.to.y; y++ {
			img.Set(x, y, p.c)
		}
	}
}

// lineParams are the parameters to call line
type lineParams struct {
	location  xy
	direction int
	length    int
	c         color.RGBA
}

// line takes an image, the start coordinates,
// the direction and distance of the line, and the color
// of the line
// func straightLine(img *image.Paletted, start xy, dir, length int, c color.Color) {
func line(img *image.Paletted, p lineParams) {
	switch p.direction {
	case left:
		for i := p.length; i >= 0; i-- {
			img.Set(p.location.x-i, p.location.y, p.c)
		}
	case right:
		for i := 0; i <= p.length; i++ {
			img.Set(p.location.x+i, p.location.y, p.c)
		}
	case up:
		for i := p.length; i >= 0; i-- {
			img.Set(p.location.x, p.location.y-i, p.c)
		}
	case down:
		for i := 0; i <= p.length; i++ {
			img.Set(p.location.x, p.location.y+i, p.c)
		}
	}
}

// createUlam creates an ulam gif and sends it to a writer
func createUlam(positions []position) *gif.GIF {

	var (
		images  []*image.Paletted    // images for the gif
		delays  []int                // delays to transition to the next gif
		lines   []lineParams         // parameters for each line creation call
		squares []squareParams       // parameters for each squares - defaulted to 0 for non primes
		lenLine = 10                 // length of each line
		wh      = 500                // to keep width and height equal
		w, h    = wh, wh             // width and height
		lastLoc = xy{wh / 2, wh / 2} // beginning position, also placeholder for each previous location
		delayCt = 0                  // delayCount for transition between images
	)

	// Loop over positions and create a new image, utilizing the previous straightLine calls
	for i, pos := range positions {

		img := image.NewPaletted(image.Rect(0, 0, w, h), palette)
		lines = append(lines, lineParams{
			location:  lastLoc,
			direction: pos.direction,
			length:    lenLine,
			c:         color.RGBA{}})

		for _, p := range lines {
			line(img, p)
		}

		for _, p := range squares {
			square(img, p)
		}

		// Change the lastLocation based on the current direction
		// to keep track of what the next location start needs to be
		switch pos.direction {
		case down:
			lastLoc = xy{lastLoc.x, lastLoc.y + lenLine}
		case up:
			lastLoc = xy{lastLoc.x, lastLoc.y - lenLine}
		case left:
			lastLoc = xy{lastLoc.x - lenLine, lastLoc.y}
		case right:
			lastLoc = xy{lastLoc.x + lenLine, lastLoc.y}
		}

		// Create a sqare for each prime
		// This is called after lastLoc updates, so it's one frame behind.
		if isPrime(pos.number) {
			s := 2 //size
			squares = append(squares, squareParams{
				from: xy{lastLoc.x - s, lastLoc.y - s},
				to:   xy{lastLoc.x + s, lastLoc.y + s},
				c:    color.RGBA{255, 150, 0, 255}})
		} else {
			squares = append(squares, squareParams{})
		}

		// Pause for effect at last iteration
		if i == len(positions)-1 {
			delayCt = 6000
		}

		delays = append(delays, delayCt)
		images = append(images, img)

	}

	return &gif.GIF{
		Image: images,
		Delay: delays,
	}
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 && i != n {
			return false
		}
	}
	return true
}
