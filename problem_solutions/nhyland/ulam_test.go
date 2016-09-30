package main

import "testing"

func TestUlam(t *testing.T) {

	tCase := []struct {
		number   int      //The position.number to test on
		position position //the expected
	}{
		{
			number: 2,
			position: position{
				number:    2,
				index:     1,
				rowCap:    1,
				direction: right,
				addCt:     false,
			},
		},
		{
			number: 3,
			position: position{
				number:    3,
				index:     1,
				rowCap:    1,
				direction: up,
				addCt:     true,
			},
		},
		{
			number: 4,
			position: position{
				number:    4,
				index:     1,
				rowCap:    2,
				direction: left,
				addCt:     false,
			},
		},
		{
			number: 5,
			position: position{
				number:    5,
				index:     2,
				rowCap:    2,
				direction: left,
				addCt:     false,
			},
		},
		{
			number: 6,
			position: position{
				number:    6,
				index:     1,
				rowCap:    2,
				direction: down,
				addCt:     true,
			},
		},
		{
			number: 7,
			position: position{
				number:    7,
				index:     2,
				rowCap:    2,
				direction: down,
				addCt:     true,
			},
		},
		{
			number: 8,
			position: position{
				number:    8,
				index:     1,
				rowCap:    3,
				direction: right,
				addCt:     false,
			},
		},
		{
			number: 9,
			position: position{
				number:    9,
				index:     2,
				rowCap:    3,
				direction: right,
				addCt:     false,
			},
		},
		{
			number: 10,
			position: position{
				number:    10,
				index:     3,
				rowCap:    3,
				direction: right,
				addCt:     false,
			},
		},
		{
			number: 11,
			position: position{
				number:    11,
				index:     1,
				rowCap:    3,
				direction: up,
				addCt:     true,
			},
		},
		{
			number: 12,
			position: position{
				number:    12,
				index:     2,
				rowCap:    3,
				direction: up,
				addCt:     true,
			},
		},
		{
			number: 14,
			position: position{
				number:    14,
				index:     1,
				rowCap:    4,
				direction: left,
				addCt:     false,
			},
		},
		{
			number: 44,
			position: position{
				number:    44,
				index:     1,
				rowCap:    7,
				direction: right,
				addCt:     false,
			},
		},
	}

	//Get the max test case number, to run pos.next() for
	maxCt := func() int {
		var c int
		for _, v := range tCase {
			if v.number > c {
				c = v.number
			}
		}
		return c
	}()

	pos := New()

	for i := 0; i < maxCt; i++ {
		p := pos.next()

		for _, v := range tCase {
			if p.number == v.number && p != v.position {
				t.Errorf("\nExpected: \n%#v\nActual:\n%#v", v.position, p)
			}
		}

	}

}
