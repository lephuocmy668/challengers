// Runtime: 100 ms, faster than 61.90% of JavaScript online submissions for Coin Change.
// Memory Usage: 41.5 MB, less than 9.09% of JavaScript online submissions for Coin Change.
// https://leetcode.com/problems/coin-change
/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function (coins, amount) {
    const memoizations = new Array(amount + 1);
    memoizations.fill(-1);
    memoizations[0] = 0;

    for (let index = 1; index < memoizations.length; index++) {
        let num = Number.MAX_SAFE_INTEGER;
        for (const coin of coins) {
            if (index - coin >= 0 && memoizations[index - coin] >= 0) {
                num = Math.min(num, memoizations[index - coin] + 1);
            }
        }
        if (num != Number.MAX_SAFE_INTEGER) memoizations[index] = num;
    }

    return memoizations[amount];
};