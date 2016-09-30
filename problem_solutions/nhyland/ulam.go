package main

/*
	ulam is responsible for creating and incrementing the ulam positions for the spiral.
*/

const (
	right = iota
	up
	down
	left
)

//Example for package
func Example() {
	positions := []position{}
	ulam := New()

	for i := 0; i < 3; i++ {
		positions = append(positions, ulam.next())
	}

}

//position is the positional coordinates in an ulam spiral
type position struct {
	number    int  //the current number
	index     int  //the current step in the line
	rowCap    int  //the cap of index allowed before turning.
	direction int  //the current direction
	addCt     bool //if true, add one to rowCt.
}

//New returns a new ulam position
func New() *position {
	return &position{
		number:    1,
		index:     0,
		rowCap:    1,
		direction: right,
		addCt:     false}
}

//next will return the new coordinates of the next drawing point.
func (p *position) next() position {
	p.number++

	if p.index != p.rowCap {
		p.index++
		return *p
	}

	if p.addCt {
		p.rowCap++
	}

	if p.number > 1 {
		p.addCt = !p.addCt
	}

	p.rotate()
	p.index = 1
	return *p
}

//rotate rotates the drawing direction 90 degrees, counter clockwise
func (p *position) rotate() {
	switch p.direction {
	case right:
		p.direction = up
	case up:
		p.direction = left
	case left:
		p.direction = down
	case down:
		p.direction = right
	}
}
