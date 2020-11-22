# 最长公共子串

这个问题有多个解法，wiki给出了广义后缀树和动态规划两个解法，同时我们这里还有一个朴素解法。

如果X和Y为两个字符串,LCS(i,j)为X[0...i]和Y[0...j]子串的公共子序列，那么有
LCS(i,j) = LCS(i-1,j-1)+1 if X[i] = Y[j] or 0 else

## 参考
最长公共子串的wiki https://en.wikipedia.org/wiki/Longest_common_substring_problem

