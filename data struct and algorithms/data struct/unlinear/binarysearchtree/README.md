# 二叉搜索树

BST(Binary Serach Tree)目的是为了提高查找的性能，其查找在平均和最坏的情况下都是log(n)的，接近二分查找。

BST:
每个节点的key大于或等于其左子树所有节点的key，小于或等于其右子树所有节点的key。

## 支持的操作

1. insert(k) : 插入一个key为k的节点到树的正确位置
2. find(k) : 查找key为k的节点
3. delete(k) : 删除key为k的节点
4. findmin(x) : 查找根为x的树的最小key节点
5. deletemin() : 查找树的最小key节点并删除
6. next-larger(x) : 查找节点x的下一个节点

## 平衡

满足BST性质且中序遍历相同的二叉查找树不是唯一的，这些二叉查找树是等价的，他们维护的是相同的一组数组。在这些二叉查找树上执行相同的操作，得到相同的结果。因此，我们可以维持BST性质的基础上，通过改变二叉查找树的形态，使得树上每个节点的左右子树大小得到平衡，从而使整个二叉查找树的深度维持在log(n)级别。
改变形态并保持BST性质的操作为二叉查找树的再平衡，是通过旋转来实现的，最基本的旋转操作为单旋转，又分为左旋转和右旋转。

## 参考
https://appliedgo.net/bintree/
https://visualgo.net/en/bst
https://blog.csdn.net/weixin_30648963/article/details/94788844
https://www.jianshu.com/p/ff4b93b088eb
https://baike.baidu.com/item/%E4%BA%8C%E5%8F%89%E6%8E%92%E5%BA%8F%E6%A0%91?fromtitle=%E4%BA%8C%E5%8F%89%E6%9F%A5%E6%89%BE%E6%A0%91&fromid=7077965