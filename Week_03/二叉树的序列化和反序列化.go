
// 3遍 时间复杂度为O(n),空间复杂度为O(n) res存放的数组长度
type Codec struct {
    res []string
}

func Constructor() Codec {
    return Codec{}
}

// 通过前序遍历来序列号
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return "none" // none是用来标识左右子数的边界
    }
    return strconv.Itoa(root.Val) + "," + this.serialize(root.Left) + 
    "," + this.serialize(root.Right)
}


// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {    
    this.res = strings.Split(data, ",")
    return this.dfsDeserialize()
}

// 反序列化的原理是前序遍历的特性每递归一层数组的值往后移一位
func (this *Codec) dfsDeserialize() *TreeNode {
    node := this.res[0]
    this.res = this.res[1:]
    if node == "none" {
        return nil
    }
    v,_ := strconv.Atoi(node)
    return &TreeNode{
        Val: v,
        Left: this.dfsDeserialize(),
        Right: this.dfsDeserialize(),
    }
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */