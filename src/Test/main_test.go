package Test

import "testing"
import "Hello/HelloWorld"

func TestSum(t *testing.T) {
	sub := HelloWorld.New("Test Sub Package")
	sub.Print()
	total := HelloWorld.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
