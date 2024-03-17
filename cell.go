package main

type Cell struct {
	Value int
}

func (c *Cell) IsEqual(other Cell) bool {
	return c.Value == other.Value
}

func (c *Cell) IsEmpty() bool {
	return c.Value == 0
}
