package main

import (
	"bufio"
	"os"
	"strings"
)

type BranchStack struct {
	file  *os.File
	stack []string
}

func NewBranchStack(path string) (*BranchStack, error) {
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		return nil, err
	}

	stack := []string{}
	for {
		reader := bufio.NewReader(f)
		line, _, err := reader.ReadLine()
		if err != nil || len(line) == 0 {
			break
		}

		stack = append(stack, string(line))
	}

	return &BranchStack{
		file:  f,
		stack: stack,
	}, nil
}

func (b *BranchStack) Push(branch string) error {
	b.stack = append(b.stack, branch)

	if _, err := b.file.WriteString(branch + "\n"); err != nil {
		return err
	}

	if err := b.file.Sync(); err != nil {
		return err
	}

	return nil
}

type StackIsEmptyError struct{}

func (e *StackIsEmptyError) Error() string {
	return "stack is empty"
}

func (b *BranchStack) Pop() (string, error) {
	if len(b.stack) == 0 {
		return "", &StackIsEmptyError{}
	}

	last := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]

	if err := b.file.Truncate(0); err != nil {
		return "", err
	}

	for _, branch := range b.stack {
		if _, err := b.file.WriteString(branch + "\n"); err != nil {
			return "", err
		}
	}

	if err := b.file.Sync(); err != nil {
		return "", err
	}

	return last, nil
}

func (b *BranchStack) Close() error {
	return b.file.Close()
}

func (b *BranchStack) Len() int {
	return len(b.stack)
}

func (b *BranchStack) Top() (string, bool) {
	if len(b.stack) == 0 {
		return "", false
	}

	return b.stack[len(b.stack)-1], true
}

func (b *BranchStack) String() string {
	sb := strings.Builder{}

	for _, branch := range b.stack {
		sb.WriteString(branch)
		sb.WriteString(" -> ")
	}

	return sb.String()
}
