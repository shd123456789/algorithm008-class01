	递归:
		树的操作一般是用递归:
			1.节点的定义
			2.重复性
		本质
			循环:通过函数体来进行的循环
		模版（go）
	```go
		func recursion(level int, p1 int, p2 int, ...) {
			// recursion terminator 递归终极条件
			if level > MAX_LEVEL {
				process_result
				return
			}
			
			// process logic in current level 处理当前层逻辑
			process(level, data, ...)

			// drill down 下探到下一层
			recursion(level+1, p1)

			// reverse the current level status if needed 清扫当前层
		}
	```
	思维要点
		1：不要人肉进行递归（最大误区）
		2：找到最近最简方法，将其拆解成可重复解决的问题（重复子问题）
		3: 数学归纳法思维(当n成立时，可以推出n+1也能成立)
	分治:
		本质：是一种递归，找重复性，将一个大问题分解成子问题
		模版：递归drill down 下面加一步合并子问题
	回溯
		本质：是一种递归，类似枚举，一直沿着一条路径下探到最深一层，找不到结果就返回上一层选择另一条路径继续寻找
		模版：跟普通递归一样
		优化：通过预先知道哪一条路径不符合结果通过剪枝（跳过不符合结果的路径）进行优化
			
