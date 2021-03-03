package main

import (
	"fmt"
)

type ClassElem struct {
	Name  string
	Value bool
}

type ClassCode struct {
	elems []string
	Code  uint64
}

func NewClassCode(elems ...ClassElem) (ec *ClassCode, err error) {
	if len(elems) > 64 {
		return nil, fmt.Errorf("input elements is greater than 64")
	}

	var i, j int
	ec = new(ClassCode)
	ec.elems = make([]string, len(elems))

	for i = 0; i < len(elems); i++ {
		for j = i + 1; j < len(elems); j++ {
			if elems[i].Name == elems[j].Name {
				return nil, fmt.Errorf("duplicate found at %d and %d", i, j)
			}
		}
		ec.elems[i] = elems[i].Name

		if elems[i].Value {
			ec.Code += 1 << i
		}
	}

	return
}

func (ec *ClassCode) Elems() (names []string) {
	names = make([]string, len(ec.elems))
	copy(names, ec.elems)
	return
}

func (ec *ClassCode) Bools() (bools []bool) {
	bools = make([]bool, len(ec.elems))
	for i := 0; i < len(ec.elems); i++ {
		bools[i] = (ec.Code>>i)%2 == 1
	}

	return
}

func (ec *ClassCode) Value(key string) (bool, error) {
	for i := range ec.elems {
		if key == ec.elems[i] {
			return (ec.Code>>i)%2 == 1, nil
		}
	}

	return false, fmt.Errorf("key not found")
}

func (ec *ClassCode) Num() int {
	return len(ec.elems)
}

func (ec *ClassCode) ClearAt(n int) (err error) {
	if n < 1 || n > len(ec.elems) {
		return fmt.Errorf("invalid number")
	}

	n--
	ec.Code &= (ec.Code ^ (1 << n))
	return
}

func (ec *ClassCode) ToggleAt(n int) (err error) {
	if n < 1 || n > len(ec.elems) {
		return fmt.Errorf("invalid number")
	}

	n--
	ec.Code |= (1 << n)
	return
}
