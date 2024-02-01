package uniquepaths

// Not accepted
// For any given m, n the number of moves will be (m-1) + (n-1)
// I'm imagining a bucket of balls with m "down" balls and n "right" balls
// The answer is "how many different ways can you pull out these balls 1 at a time without replacement?"
// Good thing I totally remember basic probability stuff really well!
func uniquePaths(m int, n int) int {
    total_steps := (m - 1) + (n - 1) - 1
	paths := 0
	for i := total_steps; i > 0; i -= 1 {
		paths += i
	}

	return paths
}