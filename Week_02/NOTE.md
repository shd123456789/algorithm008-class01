二叉树
	linked list 是特殊化的tree
	tree是特殊化的grapth
二叉搜索树
	左子树上所有结点的值均小于它的根节点的值
	右子树上所有结点的值均大于它的根节点的值
	中序遍历：升序排序
	删除如果是父节点，要拿右子树的最小节点替换
树为什么适合递归
	树本身没有所谓的后继结构，或者说便于循环的结构，更多的是左节点右节点，递归的状态展开容易实现树形结构
堆
	迅速找到一堆数中的最大值或者最小值的数据结构
	实现有二叉堆、斐波那契堆
	二叉堆性质
		是一颗完全树
		树中任意节点的值总是 >= 其子节点的值
		一般通过数组来实现
		索引为i的左孩子的索引是(2*i+1)
		索引为i的右孩子的索引是(2*i+2)
		索引为i的父节点的索引是floor((i-1)/2);
	插入删除操作:
		删除：1:将尾部堆元素替换到顶部 2:依次从根部向下调整整个堆的结构
		插入：1:新元素一律先插入到堆的尾部 2:依次向上调整整个堆的结构
	图的属性


