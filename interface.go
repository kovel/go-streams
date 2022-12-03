package main

type Stringer interface {
	String() string
}

type CollectionInterface interface {
	comparable
	String() string
}
