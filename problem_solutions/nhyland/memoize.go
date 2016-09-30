package main

import (
	"image/gif"
	"log"
	"reflect"
)

type ulamCreator func(positions []position) *gif.GIF

// memoize the createUlam function since it takes forever to compute
// at 3000 calculations, it cuts page load time in half
func memoize(fn ulamCreator) ulamCreator {
	var res struct {
		positions []position
		image     *gif.GIF
	}

	return func(p []position) *gif.GIF {
		if reflect.DeepEqual(p, res.positions) {
			log.Println("Reading from cache")
			return res.image
		}
		log.Println("Creating Ulam")
		res.positions = p
		res.image = fn(p)
		return res.image
	}
}
