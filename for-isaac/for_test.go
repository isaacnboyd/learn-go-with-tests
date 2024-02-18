package iteration

import (
	"fmt"
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("test repeat ", func(t *testing.T) {
		repeated := strings.Repeat("a", 10)
		expected := "aaaaaaaaaa"
		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})

}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Repeat("a", 20)
	}
}

func ExampleRepeat() {
	result := strings.Repeat("isaac", 5)
	fmt.Println(result)
	// Output: isaacisaacisaacisaacisaac
}
