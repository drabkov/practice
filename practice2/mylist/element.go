package mylist

type Element struct {
	value int
	next  *Element
	prev  *Element
}

func (e *Element) Next() *Element {
	if e != nil {
		return e.next
	} else {
		return nil
	}
}

func (e *Element) Prev() *Element {
	if e != nil {
		return e.prev
	} else {
		return nil
	}
}

func (e *Element) Value() int {
	return e.value
}
