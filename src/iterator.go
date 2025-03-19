package main

type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookCollection struct {
	books []string
}

func (b *BookCollection) CreateIterator() Iterator {
	return &BookIterator{books: b.books, index: 0}
}

type BookIterator struct {
	books []string
	index int
}

func (b *BookIterator) HasNext() bool {
	return b.index < len(b.books)
}

func (b *BookIterator) Next() interface{} {
	book := b.books[b.index]
	b.index++
	return book
}
