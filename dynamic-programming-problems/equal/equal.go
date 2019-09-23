// equal : https://www.hackerrank.com/challenges/equal/problem

package equal

// Complete the equal function below.
func equal(arr []int32) int32 {
    if len(arr) < 1 {
        return 0
    }

    // find min element
    min := arr[0]
    for i := 0; i < len(arr); i++ {
        if arr[i] < min {
            min = arr[i]
        }
    }

    var result int32 = math.MaxInt32
    
    for i := 0; i < 4; i++ {
        var tmpAnswer int32 = 0
       
        for j := 0; j < len(arr); j++ {
            tmpAnswer += getNumAction(arr[j]-min+int32(i));
        }

        if tmpAnswer < result {
            result = tmpAnswer;
        }
    }

    return result
}

func getNumAction(delta int32) int32 {
    var result int32 = 0

    result += delta / 5
    delta %= 5

    result += delta/2;
    delta %= 2;

    result += delta;
    return result;
}