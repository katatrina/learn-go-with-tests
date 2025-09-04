package pointers

import (
	"errors"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		
		assertBalance(t, wallet, Bitcoin(10))
	})
	
	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})
	
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	
	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got an error but didn't want one")
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()
	if got == nil {
		t.Fatal("didn't get an error but wanted one")
	}
	
	if !errors.Is(got, expected) {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
