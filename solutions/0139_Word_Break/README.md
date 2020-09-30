# 139. Word Break - Hard

```
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
```

# Relative Topics
* Dynamic Programming


# Note
```
Time: 2020/09/29
```


# Solution
## 1. Dynamic Programming

### Intuition
have a dp array, for each i, iterativly check the wordDict and check if dp[i-len(word)] valid and is true, if yes, then go for next.

### Complexity Analysis
*   Time Complexity: O(N*LM)
  
*   Space Complexity: O(N+LM)
