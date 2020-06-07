func findCircleNum(M [][]int) int {
    uf := New(len(M))
    for i := 0; i < len(M); i++ {
        for j := 0; j < len(M[0]); j++ {
            if M[i][j] == 1 {
                uf.Union(i, j)
            }
        }
    }
    return uf.Count()
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