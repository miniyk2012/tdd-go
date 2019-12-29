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
		wallet.Withdraw(Bitcoin(15))

		assertBalance(t, wallet, Bitcoin(5))
	})

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
