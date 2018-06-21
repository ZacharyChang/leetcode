# Rotate Image
-----
## Description
You are given an n x n 2D matrix representing an image.
Rotate the image by 90 degrees (clockwise).

## Note
You have to rotate the image ```in-place```, which means you have to modify the input 2D matrix directly. ```DO NOT allocate``` another 2D matrix and do the rotation.

## Example
### Example 1
Input: ```[1,2,3,4,5,6,7]``` and ```k = 3```
Output: ```[5,6,7,1,2,3,4]```
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]

### Example 2
Input: ```[-1,-100,3,99]``` and ```k = 2```
Output: ```[3,99,-1,-100]```
Explanation: 
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]

## Tags
**[Array](https://leetcode.com/tag/array)**
