// https://leetcode.com/problems/length-of-longest-fibonacci-subsequence/
// Runtime: 396 ms, faster than 21.74% of JavaScript online submissions for Length of Longest Fibonacci Subsequence.
// Memory Usage: 65.3 MB, less than 50.00% of JavaScript online submissions for Length of Longest Fibonacci Subsequence.
const lenLongestFibSubseq = function(list) {
    let result = 0;
    const map = {};
    for (let i = 0; i < list.length; i++) {
        map[list[i]] = i;
    }
  
    const memoization = [];
    for (let i = 0; i < list.length; i++) {
        memoization.push(Array(list.length).fill(0));
    }
  
    for (let i = 1; i < list.length; i++) {
        for (let j = i + 1; j < list.length; j++) {
            const key = list[j] - list[i];
            if (map[key] > -1 && key < list[i]) {
                memoization[j][i] = memoization[i][map[key]] + 1;
                result = Math.max(memoization[j][i], result);
            }
        }
    }
  
    return result ? result + 2 : result;
};