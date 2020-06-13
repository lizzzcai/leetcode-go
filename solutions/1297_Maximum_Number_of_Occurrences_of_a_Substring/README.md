# 1297. Maximum Number of Occurrences of a Substring - Medium

```
Given a string s, return the maximum number of ocurrences of any substring under the following rules:

The number of unique characters in the substring must be less than or equal to maxLetters.
The substring size must be between minSize and maxSize inclusive.
 

Example 1:

Input: s = "aababcaab", maxLetters = 2, minSize = 3, maxSize = 4
Output: 2
Explanation: Substring "aab" has 2 ocurrences in the original string.
It satisfies the conditions, 2 unique letters and size 3 (between minSize and maxSize).
Example 2:

Input: s = "aaaa", maxLetters = 1, minSize = 3, maxSize = 3
Output: 2
Explanation: Substring "aaa" occur 2 times in the string. It can overlap.
Example 3:

Input: s = "aabcabcab", maxLetters = 2, minSize = 2, maxSize = 3
Output: 3
Example 4:

Input: s = "abcde", maxLetters = 2, minSize = 3, maxSize = 3
Output: 0
 

Constraints:

1 <= s.length <= 10^5
1 <= maxLetters <= 26
1 <= minSize <= maxSize <= min(26, s.length)
s only contains lowercase English letters.
```

# Relative Topics
* String
* Bit Manipulation


# Note
```
Time: 2020/06/13

```


# Solution
## 1. String Counter

### Intuition
If a string have ocurrences x times,
any if its substring must appear at least x times.

There must be a substring of length minSize, that has the most ocurrences.
So that we just need to count the ocurrences of all subtring with length minSize.

### Algorithm
Find the maximum ocurrences of all substrings with length k = minSize

### Complexity Analysis
*   Time Complexity: O(KN) where K = minSize
  
*   Space Complexity: O(KN)