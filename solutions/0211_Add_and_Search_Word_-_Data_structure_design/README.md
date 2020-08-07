# 211. Add and Search Word - Data structure design - Medium

```
Design a data structure that supports the following two operations:

void addWord(word)
bool search(word)
search(word) can search a literal word or a regular expression string containing only letters a-z or .. A . means it can represent any one letter.

Example:

addWord("bad")
addWord("dad")
addWord("mad")
search("pad") -> false
search("bad") -> true
search(".ad") -> true
search("b..") -> true
Note:
You may assume that all words are consist of lowercase letters a-z.
```

# Relative Topics
* Tire
* Design
* Backtracking


# Note
```
Time: 2020/08/07
```


# Solution
## 1. Tire

### Intuition
Use Tire Tree to store all the words.
When searching, use dfs to process the first letter of the word, and then check it it's in the Tire tree.
if yes, dfs for the rest. if the first letter is ".", then iterate over the existing key in the Tire for dfs.

when the word reach the end, check if isWord is True.

### Complexity Analysis
*   Time Complexity: ADD O(L), Search O(L^N) worse case, 
  
*   Space Complexity: O(N*26)
