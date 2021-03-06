package Solution

import "testing"

func TestSolution(t *testing.T) {
	t.Run("Test-1", func(t *testing.T) {
		got := climbStairs(2)
		want := 2
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})

	t.Run("Test-2", func(t *testing.T) {
		got := climbStairs(6)
		want := 1
		if got != want {
			t.Errorf("GOT:", got, "WANT:", want)
		}
	})
}
