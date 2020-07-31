# 140. Word Break II - Hard

```
Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]
Example 2:

Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.
Example 3:

Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]
```

# Relative Topics
* Dynamic Programming
* Backtracking


# Note
```
Time: 2020/07/31
```


# Solution
## 1. Backtracking

### Intuition
using dfs to find all the possible result, and use memo to store the processed result to speed up.
memo[s] : for a given s, what are the construct sentences

### Complexity Analysis
*   Time Complexity: O(2^N)
  
*   Space Complexity: O(2^N)
