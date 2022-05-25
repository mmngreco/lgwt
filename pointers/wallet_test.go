package pointers

import (
	"testing"
)

func TestBitcoin(t *testing.T){
    got := Bitcoin(10).String()
    want := "10 BTC"

    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func TestWallet(t *testing.T) {

    t.Run("deposit", func(t *testing.T) {
        wallet := Wallet{}
        wallet.Deposit(Bitcoin(10))

        want := Bitcoin(10)
        assertBalance(t, wallet, want)

    })

    t.Run("withdraw", func(t *testing.T) {
        wallet := Wallet{balance: Bitcoin(20)}
        err := wallet.Withdraw(Bitcoin(10))

        want := Bitcoin(10)
        assertNoError(t, err)
        assertBalance(t, wallet, want)

    })

    t.Run("withdraw inssufficient funds", func(t *testing.T) {
        startingBalance := Bitcoin(20)
        wallet := Wallet{balance: startingBalance}
        err := wallet.Withdraw(Bitcoin(30))

        assertError(t, err, ErrInsufficientFunds)
        assertBalance(t, wallet, startingBalance)

    })
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
    got := wallet.Balance()
    if got != want {
        t.Errorf("got %s want %s", got, want)
    }
}

func assertNoError(t testing.TB, got error) {
    t.Helper()

    if got != nil {
        t.Fatal("got and error but didn't want one")
    }

}

func assertError(t testing.TB, got error, want error) {
    t.Helper()

    if got == nil {
        t.Fatal("didn't get an error but wanted one")
    }

    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

