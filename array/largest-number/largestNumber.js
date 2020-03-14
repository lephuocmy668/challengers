// https://leetcode.com/problems/largest-number
// Runtime: 64 ms, faster than 79.00% of JavaScript online submissions for Largest Number.
// Memory Usage: 36.1 MB, less than 100.00% of JavaScript online submissions for Largest Number.
function largestNumber(nums) {
    const str = nums
        .sort((a, b) => parseInt(`${b}${a}`) - parseInt(`${a}${b}`))
        .join('');
    return parseInt(str, 10) === 0 ? '0' : str;
}