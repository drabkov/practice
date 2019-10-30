package mylist

import (
	"fmt"
)

type Collection struct {
	first *Element
	last  *Element
	len   int
}

func (c *Collection) First() *Element {
	return c.first
}

func (c *Collection) Last() *Element {
	return c.last
}

func (c *Collection) Length() int {
	return c.len
}

func (c *Collection) Add(element int) {
	e := &Element{
		value: element,
	}

	if c.first == nil {
		c.first = e
	} else {
		e.prev = c.last
		c.last.next = e
	}
	c.last = e
	c.len++
}

func (c *Collection) Get(index int) *Element {
	e := c.first
	for i := 0; i < index; i++ {
		e = e.Next()
	}
	return e
}

func (c *Collection) Remove(index int) *Element {
	e := c.Get(index)
	if e != nil {
		e.prev.next = e.next
		e.next.prev = e.prev
		e.next = nil
		e.prev = nil
		c.len--
	}
	return e
}

func (c *Collection) Print() {
	fmt.Print("[")
	for e := c.First(); e != nil; e = e.Next() {
		fmt.Printf("%d", e.Value())
		if e.Next() != nil {
			fmt.Print(", ")
		}
	}
	fmt.Print("]\n")
}
