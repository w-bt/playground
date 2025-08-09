package main

import (
	"errors"
	"fmt"
	"strings"
)

type ChristmasLights struct {
	lightCount int
	innerLen   int
	idx        int
	dir        int
}

func NewChristmasLights(lightCount int) (*ChristmasLights, error) {
	if lightCount < 4 {
		return nil, errors.New("lightCount minimal 4")
	}
	return &ChristmasLights{
		lightCount: lightCount,
		innerLen:   lightCount - 2,
		idx:        0,
		dir:        1,
	}, nil
}

func (c *ChristmasLights) Next() []bool {
	lamps := make([]bool, 0, c.lightCount)

	lamps = append(lamps, true)

	mid := make([]bool, c.innerLen)
	mid[c.idx] = true
	lamps = append(lamps, mid...)

	lamps = append(lamps, true)

	if c.idx == 0 {
		c.dir = 1
	} else if c.idx == c.innerLen-1 {
		c.dir = -1
	}
	c.idx += c.dir

	return lamps
}

func main() {
	lights, err := NewChristmasLights(8)
	if err != nil {
		panic(err)
	}

	for i := 0; i < 11; i++ {
		state := lights.Next()

		var parts []string
		for _, b := range state {
			if b {
				parts = append(parts, "true")
			} else {
				parts = append(parts, "false")
			}
		}
		fmt.Printf("Step %d: %s\n", i+1, strings.Join(parts, ", "))
	}
}
