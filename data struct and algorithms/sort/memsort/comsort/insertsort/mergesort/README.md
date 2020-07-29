# 归并排序

归并排序是建立在归并操作上的一种有效的排序算法，是分治法的典型应用。

1. 分割：递归的把当前序列平均分割为两半
2. 整合：在保持元素顺序的同时将上一步得到的子序列整合在一起（归并）

## 算法
1. 分割待排序序列为两个子序列，递归进行子序列的归并排序，如果子序列为1则结束
2. 将两个已排序的子序列进行归并，得到排序序列

## 改进
如果不是划分为2个子序列，而是K个子序列，则为K路归并排序

## 过程
![](https://github.com/blockchain-develop/gopro/blob/master/img/sort/merge%20sort.png)

## 演示
![](https://github.com/blockchain-develop/gopro/blob/master/img/sort/Merge_sort_animation.gif)

## 总结
分类 排序算法
数据结构 数组
最坏时间复杂度 nlog(n)
最优时间复杂度 nlog(n)
平均时间复杂度 nlog(n)
最坏空间复杂度 n
