package roman_number

type Memo struct {
	storage map[int]int
}

var memo = Memo{storage: map[int]int{0: 0, 1: 1, 2: 1}}

func (m *Memo) Get(n int) int {
	v, ok := m.storage[n]
	if ok {
		return v
	}
	m.storage[n] = m.trib(n-3) + m.trib(n-2) + m.trib(n-1)
	return m.storage[n]
}

func (m *Memo) trib(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}
	return m.Get(n-3) + m.Get(n-2) + m.Get(n-1)
}
func tribonacci(n int) int {

	return memo.Get(n)
}
