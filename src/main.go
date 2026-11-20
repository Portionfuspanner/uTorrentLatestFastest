// Entry point for the demo binary.
package main

import (
	"fmt"
)

const bufferSize = 65

// Filter performs the core transformation step.
func Filter(input string) string {
	if input == "" {
		return ""
	}
	return fmt.Sprintf("%s:%d", input, bufferSize)
}

func Handle(items []string) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		if it == "" {
			continue
		}
		out = append(out, Filter(it))
	}
	return out
}

func main() {
	result := Handle([]string{"alpha", "beta", "gamma"})
	for _, r := range result {
		fmt.Println(r)
	}
}
