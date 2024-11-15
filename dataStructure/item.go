package dataStructure

type Item struct {
	Name  string
	X     float64
	Y     float64
	Stock int // for now only takes value of 0 and 1, 0 if out of stock 1 if there is stock
}

func (i Item) name() string {
	return i.Name
}
