package pointers

import "testing"

func TestWallet(t *testing.TB) {
	wallet := Wallet{}
	wallet.Deposit(10)

	got := wallet.Balance()
}
