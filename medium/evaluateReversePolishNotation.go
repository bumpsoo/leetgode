package medium

import "strconv"

// https://leetcode.com/problems/evaluate-reverse-polish-notation
type stack[T any] []T

func (s *stack[T]) Pop() T {
	l := len(*s)
	ret := (*s)[l-1]
	*s = (*s)[:l-1]
	return ret
}

func (s *stack[T]) Push(val T) {
	*s = append(*s, val)
}

func eval(s string) func(int, int) int {
	switch s {
	case "+":
		return func(left, right int) int { return left + right }
	case "-":
		return func(left, right int) int { return left - right }
	case "*":
		return func(left, right int) int { return left * right }
	case "/":
		return func(left, right int) int { return left / right }
	}
	return nil
}

func evalRPN(tokens []string) int {
	s := stack[int]{}
	for _, token := range tokens {
		f := eval(token)
		if f == nil {
			t, _ := strconv.Atoi(token)
			s.Push(t)
		} else {
			right, left := s.Pop(), s.Pop()
			s.Push(f(left, right))
		}
	}
	return s.Pop()
}
