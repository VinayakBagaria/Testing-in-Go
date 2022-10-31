package pointers

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s, wanted %s", got, want)
	}
}
