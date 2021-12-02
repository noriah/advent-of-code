package day17

func Part1(data []byte) int {
	grid := &Grid{}

	numCycles := 6

	grid.Init(data, numCycles, false)

	for i := 0; i < numCycles; i++ {
		grid.Cycle()
	}

	return grid.GetNumActive()
}
