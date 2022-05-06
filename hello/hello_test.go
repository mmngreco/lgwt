package main

import "testing"

func TestHello(t *testing.T) {

    assert := func(t testing.TB, got, want string) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

    t.Run("Say hello to Chris", func(t *testing.T) {
        got := Hello("Chris", "")
        want := "Hello, Chris"
        assert(t, got, want)
    })

    t.Run("Say hello to the world using empty name", func(t *testing.T) {
        got := Hello("", "")
        want := "Hello, World"
        assert(t, got, want)
    })


    t.Run("Say hello in spanish to Paco", func(t *testing.T) {
        got := Hello("Paco", "Spanish")
        want := "Hola, Paco"
        assert(t, got, want)
    })

    t.Run("Say hello in french to Camille", func(t *testing.T) {
        got := Hello("Camille", "French")
        want := "Bonjour, Camille"
        assert(t, got, want)
    })


}
