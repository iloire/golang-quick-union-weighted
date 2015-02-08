package quickunionweighted

type QuickUnion struct {
  arr []uint
  w []uint
}

func (q *QuickUnion) Initialise(n uint) {
  q.w = make([]uint, n)
  q.arr = make([]uint, n)

  for i := 0; uint(i) < n; i++ {
    q.w[i] = uint(1)
    q.arr[i] = uint(i)
  }
}

func (this *QuickUnion) Union(p, q uint){
  rootp := this.root(p)
  rootq := this.root(q)

  if this.w[rootp] < this.w[rootq]{
    this.arr[rootp] = rootq // point to rootq
    this.w[rootq] += this.w[rootp]
  } else {
    this.arr[rootq] = rootp
    this.w[rootp] += this.w[rootq]
  }  
}

func (this *QuickUnion) IsConnected(p, q uint) bool{
  return this.root(p) == this.root(q)
}

func (q *QuickUnion) root(p uint) uint {
  for p != q.arr[p]{
    p = q.arr[p]
  }
  return p
}