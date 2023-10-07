package iterator

// Iterator is the interface that wraps the basic methods for an iterator.
type Iterator[T any] interface {
	hasNext() bool
	getNext() T
}

// SliceIterator is the implementation of Iterator for a slice.
type SliceIterator[T any] struct {
	index int
	items []T
}

// NewIterator returns a new instance of SliceIterator.
func NewIterator[T any](items []T) *SliceIterator[T] {
	return &SliceIterator[T]{
		index: 0,
		items: items,
	}
}

// HasNext returns true if there is a next item to retrieve.
func (u *SliceIterator[T]) HasNext() bool {
	return u.index < len(u.items)
}

// GetNext returns the next item in the slice.
func (u *SliceIterator[T]) GetNext() (T, bool) {
	if u.HasNext() {
		item := u.items[u.index]
		u.index++
		return item, true
	}

	return *new(T), false
}

// GetNextOrDefault returns the next item in the slice or the default value if there is no next item.
func (u *SliceIterator[T]) GetNextOrDefault(d T) T {
	item, ok := u.GetNext()
	if !ok {
		return d
	}

	return item
}
