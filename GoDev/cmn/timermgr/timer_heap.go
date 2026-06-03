package timermgr

// TimerHeapType define a heap-based priority queue
type timerHeapType []*timerType

func (heap timerHeapType) getIndexByID(id int64) int {
	for _, t := range heap {
		if t.id == id {
			return t.index
		}
	}
	return -1
}

// Len return the length of TimerHeap
func (heap timerHeapType) Len() int {
	return len(heap)
}

// Less return less compare result for timer in timer heap
func (heap timerHeapType) Less(i, j int) bool {
	return heap[i].expiration.UnixNano() < heap[j].expiration.UnixNano()
}

// Swap will swap tow timers in timer heap
func (heap timerHeapType) Swap(i, j int) {
	heap[i], heap[j] = heap[j], heap[i]
	heap[i].index = i
	heap[j].index = j
}

// Push will insert a timer into timer heap
func (heap *timerHeapType) Push(x interface{}) {
	n := len(*heap)
	timer := x.(*timerType)
	timer.index = n
	*heap = append(*heap, timer)
}

// Pop will pop a timer from timer heap
func (heap *timerHeapType) Pop() interface{} {
	old := *heap
	n := len(old)
	timer := old[n-1]
	timer.index = -1
	*heap = old[0 : n-1]
	return timer
}
