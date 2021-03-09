# GO代码规范

## make slice
最好预先知道slice的大小，这样在make时指定slice的容量，在后续的append过程中，slice不会再触发扩容以及内容拷贝，效率更高。

## fallthrough
其他语言中，switch-case 结构中一般都需要在每个 case 分支结束处显式的调用 break 语句以防止 前一个 case 分支被贯穿后调用下一个 case 分支的逻辑，go 编译器从语法层面上消除了这种重复的工作，让开发者更轻松；但有时候我们的场景就是需要贯穿多个 case，但是编译器默认是不贯穿的，这个时候 fallthrough 就起作用了，让某个 case 分支再次贯穿到下一个 case 分支。