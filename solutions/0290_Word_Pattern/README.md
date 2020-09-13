# 290. Word Pattern - Easy

```
Given a pattern and a string str, find if str follows the same pattern.

Here follow means a full match, such that there is a bijection between a letter in pattern and a non-empty word in str.

Example 1:

Input: pattern = "abba", str = "dog cat cat dog"
Output: true
Example 2:

Input:pattern = "abba", str = "dog cat cat fish"
Output: false
Example 3:

Input: pattern = "aaaa", str = "dog cat cat dog"
Output: false
Example 4:

Input: pattern = "abba", str = "dog dog dog dog"
Output: false
Notes:
You may assume pattern contains only lowercase letters, and str contains lowercase letters that may be separated by a single space.
```

# Relative Topics
* Hash Table



# Note
```
Time: 2020/09/07
```


# Solution

## 1. Hash Table

### Intuition
The most naive way to start thinking about this problem is to have a single hash map, tracking which character (in pattern) maps to what word (in str). As you scan each character-word pair, update this hash map for characters which are not in the mapping. If you see a character which already is one of the keys in mapping, check whether the current word matches with the word the character maps to. If they do not match, you can immediately return False, otherwise, just keep on scanning until the end.

This type of check will work well for cases such as:

"abba" and "dog cat cat dog" -> Returns True.
"abba" and "dog cat cat fish" -> Returns False.
But it will fail for:

"abba" and "dog dog dog dog" -> Returns True (Expected False).
A fix for this is to have two hash maps, one for mapping characters to words and the other for mapping words to characters. While scanning each character-word pair,

If the character is NOT in the character to word mapping, you additionally check whether that word is also in the word to character mapping.
If that word is already in the word to character mapping, then you can return False immediately since it has been mapped with some other character before.
Else, update both mappings.
If the character IS IN in the character to word mapping, you just need to check whether the current word matches with the word which the character maps to in the character to word mapping. If not, you can return False immediately.

### Complexity Analysis
*   Time Complexity: O(N)
  
*   Space Complexity: O(M) 
