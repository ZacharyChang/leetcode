# Prefix and Suffix Search
-----
## Description
Given many words, `words[i]` has `weight i`.

Design a class WordFilter that supports one function, `WordFilter.f(String prefix, String suffix)`. It will return the word with given prefix and suffix with `maximum weight`. If no word exists, return `-1`.

## Example
### Example 1
```
Input:
WordFilter(["apple"])
WordFilter.f("a", "e") // returns 0
WordFilter.f("b", "") // returns -1
```

## Note
1. words has length in range `[1, 15000]`.
2. For each test case, up to `words.length` queries `WordFilter.f` may be made.
3. `words[i]` has length in range `[1, 10]`.
4. prefix, suffix have lengths in range `[0, 10]`.
5. `words[i]` and `prefix`, `suffix` queries consist of `lowercase letters only`.

## Tags
**[Trie](https://leetcode.com/tag/trie)**
