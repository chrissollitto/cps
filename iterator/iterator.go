package iterator

// CPSIterator is a generic implementation of an iterator for CPS recursive functions
type CPSIterator[S any, T any] struct {
	hasNext  bool
	head     T
	nextCont func() T
	cpsFunc  func(s *S, cont func() T) T
}

func (it *CPSIterator[S, T]) Init(start *S, f func(s *S, cont func() T) T) {
	it.cpsFunc = f
	it.hasNext = true
	it.head = it.Run(start, nil)
}

func (it *CPSIterator[S, T]) Next() T {
	if !it.hasNext {
		panic("Empty Iterator")
	}
	val := it.head
	if it.nextCont != nil {
		it.head = it.nextCont()
	} else {
		it.hasNext = false
	}
	return val
}

func (it *CPSIterator[S, T]) HasNext() bool {
	return it.hasNext
}

// run is the flow control function
func (it *CPSIterator[S, T]) Run(s *S, cont func() T) T {
	if s != nil {
		return it.cpsFunc(s, cont)
	} else if cont != nil {
		return cont()
	} else {
		it.nextCont = nil
		it.hasNext = false
		return *new(T)
	}
}

// suspend suspends the current computation and sets the continuation
func (it *CPSIterator[S, T]) Suspend(emitVal T, cont func() T) T {
	it.nextCont = cont
	return emitVal
}
