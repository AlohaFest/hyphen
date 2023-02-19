
//
// stack.go
// Copyright (C) 2020 Toran Sahu <toran.sahu@yahoo.com>
//
// Distributed under terms of the MIT license.
//

package stack

type Stack []float32

func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Push(input float32) {
	*s = append(*s, input)