package main

import "testing"

func TestCourseTopics(t *testing.T) {
	topics := courseTopics()

	if len(topics) == 0 {
		t.Fatal("courseTopics() returned no topics")
	}

	expected := []string{"RESTful Services", "WebSocket Protocols", "Generics", "Practical Utility Development"}
	if len(topics) != len(expected) {
		t.Fatalf("expected %d topics, got %d", len(expected), len(topics))
	}

	for i, want := range expected {
		if topics[i].name != want {
			t.Errorf("topic[%d]: expected %q, got %q", i, want, topics[i].name)
		}
		if topics[i].description == "" {
			t.Errorf("topic[%d] (%s) has an empty description", i, topics[i].name)
		}
	}
}
