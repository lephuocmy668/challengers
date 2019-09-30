package fm

// Complete the fibonacciModified function below.
func fibonacciModified(t1 int32, t2 int32, n int32) int32 {
    iTerm := int64(t1)
    iPlusOneTerm := int64(t2)

    for i := int64(3); i <= int64(n); i++ {
        newElement := iTerm + iPlusOneTerm * iPlusOneTerm
        if i == int64(n) {
            return int32(newElement)
        }
        iTerm = iPlusOneTerm
        iPlusOneTerm = newElement
    }
    return int32(iPlusOneTerm)
}
