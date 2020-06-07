func solve(board [][]byte)  {
    if len(board) == 0 {
        return
    }
    row,col := len(board),len(board[0])
    uf := New((row * col) + 1)
    rootNode := row * col
    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if board[i][j] == 'O' { 
                if i == 0 || i == row - 1 || j == 0 || j == col - 1 {
                    uf.Union(col * i + j, rootNode)
                } else {
                    if i > 0 &&  board[i - 1][j] == 'O' {
                        uf.Union(col * i + j , col * (i - 1) + j)
                    }
                    if  i < row - 1 && board[i + 1][j] == 'O' {
                        uf.Union(col * i + j , col * (i + 1) + j)
                    }
                    if  j > 0 && board[i][j - 1] == 'O' {
                        uf.Union(col * i + j , col * i + j - 1)
                    }
                    if  j < col - 1 && board[i][j + 1] == 'O' {
                        uf.Union(col * i + j , col * i + j + 1)
                    }
                }
            }
        }
    }

    for i := 0; i < row; i++ {
        for j := 0; j < col; j++ {
            if uf.Connected(col * i + j, rootNode) {
                board[i][j] = 'O'
            } else {
                board[i][j] = 'X'
            }
        }
    }
}


type Uf struct {
	count int
	parent []int
	size []int
}

func New(n int) Uf {
	uf := Uf {
		parent : make([]int, n),
		size : make([]int, n),
        count : n,
	}
	for i := 0; i < n; i++ {
		uf.parent[i] = i
		uf.size[i] = 1
	}
	return uf
}

func (this *Uf) Union(p int, q int) {
	rootP := this.Find(p)
	rootQ := this.Find(q)
	if rootP == rootQ {
		return
	}
	if this.size[rootQ] > this.size[rootP] {
		this.parent[rootP] = rootQ
		this.size[rootQ] += this.size[rootP]
	} else {
		this.parent[rootQ] = rootP
		this.size[rootP] += this.size[rootQ]
	}
	this.count--
}

func (this *Uf) Connected(p int, q int) bool {
	rootP := this.Find(p)
	rootQ := this.Find(q)
	return rootP == rootQ
}

func (this *Uf) Find(p int) int {
	for {
		if p == this.parent[p] {
			break
		}
		this.parent[p] = this.parent[this.parent[p]]
		p = this.parent[p]
	}
	return p
}

func (this *Uf) Count() int {
	return this.count
}