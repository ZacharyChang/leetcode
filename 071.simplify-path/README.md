# Simplify Path
-----
## Description
Given an absolute path for a file (Unix-style), simplify it.

## Example
```
path = "/home/", => "/home"
path = "/a/./b/../../c/", => "/c"
```

## Corner Cases
1. Did you consider the case where path = "/../"?
In this case, you should return "/".
2. Another corner case is the path might ```contain multiple slashes '/' together```, such as "/home//foo/".
In this case, you should ```ignore redundant slashes``` and return "/home/foo".

## Tags
**[String](https://leetcode.com/tag/string/)**

**[Stack](https://leetcode.com/tag/stack/)**
