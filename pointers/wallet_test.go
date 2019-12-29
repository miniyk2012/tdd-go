package pointers

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		fmt.Println("address of balance in test is", &wallet.balance)
		if wallet.Balance() != want {
			t.Errorf("want '%s' got '%s'", want, got)
		}
	}
	assertError := func(t *testing.T, err error, want string) {
		if err == nil {
			// Fatal会停止执行, 否则(如Error)测试将继续进行下一步，并且因为一个空指针而引起 panic
			t.Fatal("didn't get an error but wanted one")
		}
		if err.Error() != want {
			t.Errorf("got '%s', want '%s'", err, want)
		}
	}

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
		assertError(t, err, "cannot withdraw, insufficient funds")
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(15))

		assertBalance(t, wallet, Bitcoin(5))
	})

}
