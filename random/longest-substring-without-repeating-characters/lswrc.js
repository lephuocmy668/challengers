// https://leetcode.com/problems/longest-substring-without-repeating-characters
// Runtime: 96 ms, faster than 97.04% of JavaScript online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 43.9 MB, less than 64.00% of JavaScript online submissions for Longest Substring Without Repeating Characters.
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  let maxLength = 0;
  let subString = '';
  for (let i = 0; i < s.length; i++) {
    const char = s[i];
    const indexOfChar = subString.indexOf(char);
    if (indexOfChar > -1) {
      subString = subString.slice(indexOfChar + 1);
    }
    subString = subString + char;
    if (maxLength < subString.length) {
      maxLength = subString.length;
    }
  }
  return maxLength;
};

// Runtime: 364 ms, faster than 17.13% of JavaScript online submissions for Longest Substring Without Repeating Characters.
// Memory Usage: 46.1 MB, less than 18.51% of JavaScript online submissions for Longest Substring Without Repeating Characters.
/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
  let maxLength = 0;
  let charMap = {};
  let left = 0;
  let right = 0;
  while (right < s.length) {
    const char = s[right];
    if (charMap[char]) {
      delete charMap[s[left]];
      left++;
    } else {
      charMap[char] = true;
      right++;
    }
    maxLength = Math.max(maxLength, Object.keys(charMap).length);
  }
  return maxLength;
};
