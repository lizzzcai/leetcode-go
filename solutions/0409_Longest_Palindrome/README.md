# 409. Longest Palindrome - Easy

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
## 1. Hash Table

### Intuition

    The question ask for the length of the longest Palindrome. We can count the number of each ch in given s.
    as the palindrome string is mirrored on the center point.
    assume we have the center point, and we have the count of each ch in the s.
    every time for each ch, we take 2 ch out from the count and pair it.
    we pair the ch on both side of the center point to make it palindrome.
    at the end, we can add one more to ch to the center points if there are any extra ch not being paired.


### Complexity Analysis
*   Time Complexity: O(N)
*   Space Complexity: O(N)