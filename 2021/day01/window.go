package day01

type Window struct {
	index    int
	length   int
	sum      int
	Data     []int
	Capacity int
}

// Add a value to the window, shifting if needed. Returns new sum.
func (w *Window) Push(value int) int {

	if w.length < w.Capacity {
		w.length++
		w.sum += value
	} else {
		old := w.Data[w.index]
		w.sum += value - old
	}

	w.Data[w.index] = value

	if w.index++; w.index >= w.Capacity {
		w.index = 0
	}

	return w.sum
}

// Returns sum of window values.
func (w *Window) Sum() int {
	return w.sum
}
