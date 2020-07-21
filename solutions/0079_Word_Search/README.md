# 79. Word Search - Medium

```
Given a 2D board and a word, find if the word exists in the grid.

The word can be constructed from letters of sequentially adjacent cell, where "adjacent" cells are those horizontally or vertically neighboring. The same letter cell may not be used more than once.

Example:

board =
[
  ['A','B','C','E'],
  ['S','F','C','S'],
  ['A','D','E','E']
]

Given word = "ABCCED", return true.
Given word = "SEE", return true.
Given word = "ABCB", return false.
 

Constraints:

board and word consists only of lowercase and uppercase English letters.
1 <= board.length <= 200
1 <= board[i].length <= 200
1 <= word.length <= 10^3
```

# Relative Topics
* Backtracking


# Note
```
Time: 2020/07/21
```


# Solution
## 1. Backtracking

### Intuition
dfs starting from each cell, and explore the neighbor. save the visit of each cell. for each cell, if board[i][j] == word[idx], when we keep explore to the 4 neigbors and find if any cell meet word[idx+1]. if idx == len(word), return true. everytime we explore the neighbor from a cell, we need to add this cell into visited to avoid it being used again.

### Complexity Analysis
*   Time Complexity: O(N^2)
  
*   Space Complexity: O(L)
