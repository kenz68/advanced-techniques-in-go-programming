// Lesson 01 — Course Introduction & Objectives
//
// This program introduces the Advanced Techniques in Go Programming course.
// It demonstrates that your Go environment is configured correctly and gives
// a brief preview of each topic the course will cover.
package main

import (
	"fmt"
	"runtime"
)

// topic describes a single course topic.
type topic struct {
	name        string
	description string
}

// courseTopics returns the list of topics covered in this course.
func courseTopics() []topic {
	return []topic{
		{
			name:        "RESTful Services",
			description: "Design and build HTTP APIs using net/http, routing, middleware, and JSON handling.",
		},
		{
			name:        "WebSocket Protocols",
			description: "Enable real-time, bidirectional communication between clients and servers.",
		},
		{
			name:        "Generics",
			description: "Write reusable, type-safe code with Go type parameters and constraints.",
		},
		{
			name:        "Practical Utility Development",
			description: "Build command-line tools and small services that solve real-world problems.",
		},
	}
}

func main() {
	fmt.Println("================================================")
	fmt.Println("  Advanced Techniques in Go Programming")
	fmt.Println("  Lesson 01 — Course Introduction & Objectives")
	fmt.Println("================================================")
	fmt.Printf("\nGo version : %s\n", runtime.Version())
	fmt.Printf("OS/Arch    : %s/%s\n", runtime.GOOS, runtime.GOARCH)

	fmt.Println("\n--- Who is this course for? ---")
	fmt.Println("  • Intermediate Go developers who know the basics and want to go further.")
	fmt.Println("  • Experienced programmers from other languages learning advanced Go patterns.")

	fmt.Println("\n--- Topics covered ---")
	for i, t := range courseTopics() {
		fmt.Printf("  %d. %s\n     %s\n\n", i+1, t.name, t.description)
	}

	fmt.Println("--- What to expect ---")
	fmt.Println("  Each lesson combines concise theory with hands-on exercises.")
	fmt.Println("  Run `go run main.go` inside any lesson directory to follow along.")
	fmt.Println("\nYour environment is ready. Let's get started!")
}
