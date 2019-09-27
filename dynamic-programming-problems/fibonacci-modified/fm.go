package fm

// Complete the fibonacciModified function below.
func fibonacciModified(t1 int32, t2 int32, n int32) int32 {
	iTerm := t1
	iPlusOneTerm := t2

	for i := int32(3); i < n; i++ {
		newElement := iTerm + (iPlusOneTerm+1)*(iPlusOneTerm+1)
		if i == n {
			return newElement
		}
		iTerm = iPlusOneTerm
		iPlusOneTerm = newElement
	}
	return 0
}
