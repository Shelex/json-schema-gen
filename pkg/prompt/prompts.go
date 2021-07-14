package prompt

import (
	"github.com/manifoldco/promptui"
)

// Executor interface for `prompt` package functions
type Executor interface {
	Confirm(label string) (bool, error)
	Input(label string) (string, error)
	Select(label string, items []string, size int) (string, error)
}

// UI - implementation of the promptExecutor interface
type UI struct {
	Validate func(string) error
}

// Confirm - prompt with boolean answer
func (p UI) Confirm(label string) (bool, error) {
	confirmOverwrite := promptui.Prompt{
		Label:     label,
		IsConfirm: true,
	}
	if _, err := confirmOverwrite.Run(); err != nil {
		return false, nil
	}
	return true, nil
}

// Input - prompt with string answer accepted
func (p UI) Input(label string) (string, error) {
	input := promptui.Prompt{
		Label:    label,
		Validate: p.Validate,
	}
	result, err := input.Run()
	return result, err
}

// Select - prompt to select one from multiple options
func (p UI) Select(label string, items []string, size int) (string, error) {
	selector := promptui.Select{
		Label: label,
		Items: items,
		Size:  size,
	}
	_, result, err := selector.Run()
	return result, err
}
