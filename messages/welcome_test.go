package messages

import (
	"go-hwapp/messages"
	"testing"
)

func TestWelcomeMessage(t *testing.T) {
	welcome := messages.Welcome{Message: "hi"}

	if welcome != "hi" {
		t.Fatalf("Expected hi")
	}
}
