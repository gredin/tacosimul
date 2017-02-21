package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPosition_ManhattanDistance(t *testing.T) {
	cases := []struct {
		p1                Position
		p2                Position
		expected_distance int
	}{
		{Position{1, 1}, Position{3, 3}, 4},
		{Position{0, 0}, Position{3, 3}, 6},
	}

	for _, c := range cases {
		/*if d := c.p1.ManhattanDistance(c.p2); d != c.d {
			t.Errorf("distance: %v (expected: %v)", d, c.d)
		}*/

		d := c.p1.ManhattanDistance(c.p2)
		assert.Equal(t, c.expected_distance, d)
	}
}

func TestPosition_MoveTo(t *testing.T) {
	cases := []struct {
		p                 Position
		destination       Position
		steps             int
		expected_position Position
	}{
		{Position{0, 0}, Position{3, 3}, 4, Position{3, 1}},
		{Position{0, 0}, Position{3, 3}, 6, Position{3, 3}},
		{Position{0, 0}, Position{3, 3}, 10, Position{3, 3}},
	}

	for _, c := range cases {
		assert.Equal(t, c.expected_position, c.p.MoveTo(c.destination, c.steps))
	}
}

func TestPosition_Equals(t *testing.T) {
	assert.Equal(t, true, Position{1, 1}.Equals(Position{1, 1}))
	assert.Equal(t, false, Position{1, 1}.Equals(Position{0, 1}))
	assert.Equal(t, false, Position{1, 1}.Equals(Position{1, 0}))
}
