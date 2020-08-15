# 435. Non-overlapping Intervals - Medium

* Link:
https://leetcode.com/problems/non-overlapping-intervals/
* Difficulty: Medium

# Description

```
Given a string which consists of lowercase or uppercase letters, find the length of the longest palindromes that can be built with those letters.

This is case sensitive, for example "Aa" is not considered a palindrome here.

Note:
Assume the length of given string will not exceed 1,010.

Example:

Input:
"abccccdd"

Output:
7

Explanation:
One longest palindrome that can be built is "dccaccd", whose length is 7.
```

# Relative Topics
* Hash Table


# Note
```
Time: 2020/08/15

```


# Solution
## 1. Greedy

### Intuition
    we sort by the end time, and select the first one as the current end time.
    iterate from the second one, if the start time < current end, which mean there are overlap between them.
    we will remove this interval. Otherwise, update the the current end to the end time of this interval.

    why sort by end:
    [ [1,4], [2,3], [3,4] ], the interval with early start might be very long and incompatible with many intervals. But if we choose the interval that ends early, we'll have more space left to accommodate more intervals. Hope it helps.

### Complexity Analysis
*   Time Complexity: O(NlogN)
*   Space Complexity: O(logN)