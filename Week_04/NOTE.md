
   搜索：
	
	概念：绝大部分处理的都是暴力或者说比较简单朴树的搜索，通常情况下做的一件事情就是把全部节点遍历一遍（有且仅有一遍）
	类型：对与节点的访问顺序不同分为：深度优先:depth first search 广度优先 breadthf irst search
	模版：

		dfs：// 递归写法
		// 二叉树
		visited = set() 
		def dfs(node, visited):
		    if node in visited: # terminator
		    	# already visited 
		    	return 
			visited.add(node) 
			# process current node here. 
			...
			for next_node in node.children(): 
				if next_node not in visited: 
			dfs(next_node, visited)
		// 多叉树
		func dfs(root *node) {
			if visited[root] {  // already visited
				return
			}
			visited[root] = true  // process current node
			for _,nextNode := range root.Children {
				dfs(nextNode)	
			}
		}
		// 非递归写法
		def DFS(self, tree): 
			if tree.root is None: 
				return [] 
			visited, stack = [], [tree.root]
			while stack: 
				node = stack.pop() 
				visited.add(node)
				process (node) 
				nodes = generate_related_nodes(node) 
				stack.push(nodes) 
			# other processing work 
		bfs:
			def BFS(graph, start, end):
			    visited = set()
				queue = [] 
				queue.append([start]) 
				while queue: 
					node = queue.pop() 
					visited.add(node)
					process(node) 
					nodes = generate_related_nodes(node) 
					queue.push(nodes)
				# other processing work 
				...
			}
贪心算法：
	
	概念：贪心算法是选择单下最优，近而希望推出全局最优的算法，与动态规划的区别：保存之前的运算结果，并根据以前的结果对当前进行选择，有回退功能
	适用: 问题嫩够分解成子问题解决，子问题的最优解能递推到问题的最优解，这种子问题称为最优子结构（没有回退）
二分查找
	
	前提：目标函数单调性，存在上下界，可以通过索引访问
	模版:
		
		left, right = 0, len(array) - 1 
		while left <= right: 
		  mid = (left + right) / 2 
		  if array[mid] == target: 
			    # find the target!! 
			    break or return result 
		  elif array[mid] < target: 
			    left = mid + 1 
		  else: 
			    right = mid - 1
