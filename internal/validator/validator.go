package validator

import (
	"context"
	"strings"
)

type Validator interface {
	Valid(context.Context) Evaluator
}

type Evaluator map[string]string

func (e *Evaluator) AddFieldError(key, message string) {
	if *e == nil {
		*e = make(map[string]string)

	}
	if _, exists := (*e)[key]; !exists {
		(*e)[key] = message
	}
}

func (e *Evaluator) CheckField(ok bool, key, message string) {
	if !ok {
		e.AddFieldError(key, message)
	}
}

func NotBlank(value string) bool {
	return strings.TrimSpace(value) != ""
}

func NotBlankNumber(value int) bool {
	return value != 0
}

func MinChar(value string, length int) bool {
	return len(value) >= length
}

func MaxChar(value string, length int) bool {
	return len(value) <= length
}

func NonNegativeValue(value int, n int) bool {
	return value > n
}

func MinValue(value int, min int) bool {
	return value >= min
}

func MaxValue(value int, max int) bool {
	return value <= max
}
