### BFS
常配合队列数据结构来解题
### DFS
可以使用递归或者配合队列的方法来解决

图的遍历可以增加一个visited的set


二叉树是一种非常重要的数据结构，很多其它数据结构都是基于二叉树的基础演变而来的。
对于二叉树，有深度遍历和广度遍历，深度遍历有前序、中序以及后序三种遍历方法，
广度遍历即我们平常所说的层次遍历。因为树的定义本身就是递归定义，
因此采用递归的方法去实现树的三种遍历不仅容易理解而且代码很简洁，
而对于广度遍历来说，需要其他数据结构的支撑，比如堆了。所以，对于一段代码来说，
可读性有时候要比代码本身的效率要重要的多。

四种主要的遍历思想为：

前序遍历：根结点 ---> 左子树 ---> 右子树

中序遍历：左子树---> 根结点 ---> 右子树

后序遍历：左子树 ---> 右子树 ---> 根结点

``` 
由于从给定的某个节点出发，有多个可以前往的下一个节点（树不是线性数据结构），
所以在顺序计算（即非并行计算）的情况下，只能推迟对某些节点的访问——即以某种方式保存起来以便稍后再访问。
常见的做法是采用栈（LIFO）或队列（FIFO）。由于树本身是一种自我引用（即递归定义）的数据结构，
因此很自然也可以用递归方式，或者更准确地说，用corecursion，来实现延迟节点的保存。
这时（采用递归的情况）这些节点被保存在call stack中。
``` 