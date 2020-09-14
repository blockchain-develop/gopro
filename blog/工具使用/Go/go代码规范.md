# GO代码规范

## make slice
最好预先知道slice的大小，这样在make时指定slice的容量，在后续的append过程中，slice不会再触发扩容以及内容拷贝，效率更高。