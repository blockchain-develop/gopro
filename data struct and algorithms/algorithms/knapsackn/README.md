# 最长公共子序列

如果X和Y为两个字符串,LCS(i,j)为X[0...i]和Y[0...j]子串的公共子序列，那么有
LCS(i,j) = LCS(i-1,j-1)+1 if X[i] = Y[j] or max(LCS(i,j-1),LCS(i-1,j)) else

如果i为第i件物品，W(i)为物品重量，V(i)为物品价值，MV(i)为第i件物品取舍后背包中的总价值，那么有
MV(i) = MV(i - 1) + V(i) if 选择i or MV(i -1) else


## 参考
mit 6.046 2005 公开课
geeksforgeeks https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/
最长公共子序列的wiki https://en.wikipedia.org/wiki/Longest_common_subsequence_problem
最长公共子序列算法博客
https://blog.csdn.net/so_geili/article/details/53737001
https://blog.csdn.net/zhaoluwei/article/details/52193985

