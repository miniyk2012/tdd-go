package pointers

import (
	"fmt"
	"testing"
)

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	fmt.Println("address of balance in test is", &wallet.balance)
	if got != want {
		t.Errorf("want '%s' got '%s'", want, got)
	}
}

func assertError(t *testing.T, err error, want error) {
	if err == nil {
		// Fatal会停止函数后续的执行, 否则(如t.Error)测试将继续进行下一步，并且因为一个空指针而引起 panic
		t.Fatal("didn't get an error but wanted one")
	}
	if err != want {
		t.Errorf("got '%s', want '%s'", err, want)
	}
}

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, InsufficientFundsError)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		error := wallet.Withdraw(Bitcoin(15))

		assertBalance(t, wallet, Bitcoin(5))
		assertNoError(t, error)
	})

}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}

func TestReturn(t *testing.T) {
	t.Run("指针", func(t *testing.T) {
		origin := Bitcoin(10)
		wallet := Wallet{origin}
		ReturnWallet := func(w *Wallet) *Wallet {
			return w // w是指针
		}

		newWallet := ReturnWallet(&wallet)
		want := Bitcoin(15)
		newWallet.balance = want

		if newWallet.Balance() != want {
			t.Errorf("newWallet want '%s', got '%s'", want, newWallet.Balance())
		}
		if wallet.Balance() != want {
			t.Errorf("wallet want '%s', got '%s'", want, wallet.Balance())
		}
	})

	t.Run("值复制", func(t *testing.T) {
		origin := Bitcoin(10)
		wallet := Wallet{origin}
		ReturnWallet := func(w Wallet) Wallet {
			return w // w是一个新的Wallet
		}

		newWallet := ReturnWallet(wallet)
		want := Bitcoin(15)
		newWallet.balance = want

		if newWallet.Balance() != want {
			t.Errorf("newWallet want '%s', got '%s'", want, newWallet.Balance())
		}
		if wallet.Balance() != origin {
			t.Errorf("wallet want '%s', got '%s'", origin, wallet.Balance())
		}
	})

}

func TestArray(t *testing.T) {
	var arr1 [3]int
	arr2 := arr1 // copy一个新的array
	fmt.Println(arr1[0])

	var sl1 []int
	fmt.Println(sl1 == nil) // true
	fmt.Println(sl1)
	fmt.Printf("len=%d, cap=%d\n", len(sl1), cap(sl1))

	fmt.Println(arr2)
	arr1[0] = 5
	fmt.Println(arr2)
	fmt.Println(arr1)

	sl2 := make([]int, 0)
	fmt.Println(sl2)
	fmt.Println(sl2 == nil) // false
	fmt.Printf("len=%d, cap=%d\n", len(sl2), cap(sl2))

	fmt.Println(nil)

	fmt.Println(make([]int, 5, 6))
}

func TestSlice(t *testing.T) {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b)
	c := b[1:4] // c与slice共享内存
	c[2] = 66
	fmt.Println(b)
	fmt.Println(c)

	x := [3]string{"Лайка", "Белка", "Стрелка"}
	s := x[:] // s与数组共享内存
	s[1] = "yangkai"
	fmt.Println(x)
	fmt.Println(s)
}

func TestLen(t *testing.T) {
	s := make([]byte, 5)
	s = s[2:4]
	for _, v := range s {
		fmt.Println(v)  // for是根据len来的, len=2就迭代2次
	}
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s))
	s = s[:cap(s)]                                 // We can grow s to its capacity by slicing it againd
	for _, v := range s {
		fmt.Println(v)  // len增加为3, 就迭代3次
	}
	fmt.Printf("len=%d, cap=%d\n", len(s), cap(s)) // A slice cannot be grown beyond its capacity
}

func TestAppend(t *testing.T) {
	a := make([]int, 1, 5)
	fmt.Printf("len=%d, cap=%d\n", len(a), cap(a))
	// a == []int{0}
	a = append(a, 1, 2, 3)
	// a == []int{0, 1, 2, 3}
	fmt.Printf("len=%d, cap=%d\n", len(a), cap(a))
}

func Filter(s []int, fn func(int) bool) []int {
	//r := make([]int, 0)
	var r []int  // == nil
	for i := range s {
		if fn(s[i]) {
			r = append(r, s[i])
		}
	}
	return r
}
func notZero(v int) bool {
	return v != 0
}

func TestFilter(t *testing.T) {
	s := []int {10, 0, 12, 0, -2}
	f := Filter(s, notZero)
	fmt.Println(f)
	fmt.Println(s)
}