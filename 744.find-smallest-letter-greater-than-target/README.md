# Find Smallest Letter Greater Than Target
-----
## Description
Given a list of sorted characters letters containing ```only lowercase letters```, and given a target letter target, find the ```smallest element``` in the list that is l```arger than the given target```.

Letters also ```wrap around```. For example, if the target is target = 'z' and letters = ['a', 'b'], the answer is 'a'.

## Example
### Example 1
```
Input:
letters = ["c", "f", "j"]
target = "a"
Output: "c"
```

### Example 2
```
Input:
letters = ["c", "f", "j"]
target = "c"
Output: "f"
```

### Example 3
```
Input:
letters = ["c", "f", "j"]
target = "d"
Output: "f"
```

### Example 4
```
Input:
letters = ["c", "f", "j"]
target = "g"
Output: "j"
```

### Example 5
```
Input:
letters = ["c", "f", "j"]
target = "j"
Output: "c"
```

### Example 6
```
Input:
letters = ["c", "f", "j"]
target = "k"
Output: "c"
```

## Note
1. letters has a length in range ```[2, 10000]```.
2. letters consists of lowercase letters, and contains ```at least 2 unique``` letters.
3. target is a ```lowercase``` letter.

## Tags
**[Binary Search](https://leetcode.com/tag/BinarySearch)**
