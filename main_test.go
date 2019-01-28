package coat

import "testing"

func TestPrint(t *testing.T) {
	console := NewConsole("test", false)
	console.Log("hello")
}
