package pointers

import "fmt"

// 从现有的类型创建新的类型
type Bitcoin int

// 实现 Bitcoin 的 fmt/print/Stringer 接口, '%s'时会使用该方法
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// 在 Go 中，当调用一个函数或方法时，参数会被复制, 因此w是一个副本, 用指针的话, 指向的就是同一个对象了
	w.balance += amount
	fmt.Println("address of balance in Deposit is", &w.balance)
}

func (w *Wallet) Balance() Bitcoin {
	fmt.Println("address of balance in Balance is", &w.balance)
	return w.balance
}
