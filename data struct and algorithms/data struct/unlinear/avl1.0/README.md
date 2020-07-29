# AVL树

AVL树是自平衡二叉查找树，AVL树中任何节点的两个子树的高度差最大为1，所以它被称为高度平衡树。查找、插入和删除在平均和最坏的情况下的时间复杂度都为log(n)。增加和删除可能需要通过一次或者多次旋转来重新平衡树。

节点的平衡因子是它的左子树的高度减去右子树的高度，带有平衡因子1、0或-1的节点被认为是平衡的。带有-2、2平衡因子的节点被认为是不平衡的，需要重新平衡树。

## 参考
AVL树的wiki https://zh.wikipedia.org/wiki/AVL%E6%A0%91
Balancing a binary search tree https://appliedgo.net/balancedtree/

https://www.sohu.com/a/270452030_478315
https://www.jianshu.com/p/65c90aa1236d
https://baike.baidu.com/item/AVL%E6%A0%91/10986648?fr=aladdin
go编程指导中的AVL树的实现 https://www.golangprograms.com/golang-program-for-implementation-of-avl-trees.html