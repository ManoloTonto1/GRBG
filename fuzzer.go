package main

type FuzzExploiter struct{}

func NewFuzzExploiter() *FuzzExploiter {
	return &FuzzExploiter{}
}
func (f *FuzzExploiter) Exploit() bool {

	return true
}
