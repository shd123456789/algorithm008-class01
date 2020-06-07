var res []string
func findWords(board [][]byte, words []string) []string {
    res = make([]string, 0)
    trie := Constructor()
    for i := 0; i < len(words); i++ {
        trie.Insert(words[i])
    }
    for i := 0; i < len(board); i++ {
        for j := 0; j < len(board[0]); j++ {
            dfs(trie.root, i, j, board)
        }
    }
    return res
}

func dfs (node *TrieNode, i int, j int, board [][]byte) {
    if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) {
        return 
    }
    c := board[i][j]
    if c == 'V' || node == nil || node.children[c - 'a'] == nil {
        return
    }
    node = node.children[c - 'a']
    if node.data != "" {
        res = append(res, node.data)
        node.data = ""
    }
    board[i][j] = 'V'
    dfs(node, i - 1, j, board)
    dfs(node, i + 1, j, board)
    dfs(node, i, j - 1, board)
    dfs(node, i, j + 1, board)
    board[i][j] = c
}

type Trie struct {
    root *TrieNode
}

type TrieNode struct {
    data string
	children []*TrieNode
	isEnd bool
}

func Constructor() Trie {
    return Trie{
        root : &TrieNode{children : make([]*TrieNode, 26)},
    }
}

func (this *Trie) Insert(word string)  {
    p := this.root
    for i := 0; i < len(word); i++ {
        idx := word[i] - 'a'
        if p.children[idx] == nil {
            p.children[idx] = &TrieNode{children : make([]*TrieNode, 26)}
        }
        p = p.children[idx]
    }
    p.isEnd = true
    p.data = word
}