# 472. Concatenated Words - Hard

```
Given a list of words (without duplicates), please write a program that returns all concatenated words in the given list of words.
A concatenated word is defined as a string that is comprised entirely of at least two shorter words in the given array.

Example:
Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]

Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats"; 
 "dogcatsdog" can be concatenated by "dog", "cats" and "dog"; 
"ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".
Note:
The number of elements of the given array will not exceed 10,000
The length sum of elements in the given array will not exceed 600,000.
All the input string will only include lower case letters.
The returned elements order does not matter.
```

# Relative Topics
* Dynamic Programming
* DFS
* Trie

# Note
```
Time: 2020/09/14

```


# Solution
## 1. DFS

### Intuition

iterate each word, and create the prefix and suffix over the word from begin till end.
1. check if prefix and suffix in the word list, if yes, return true for this word
2. check if prefix in the word list and dfs over the suffix, if both true, then return true
3. check if suffix in the word list and dfs over the prefix, if both true, then return true


### Complexity Analysis
*   Time Complexity: O(NW^2)
  
*   Space Complexity: O(NW)