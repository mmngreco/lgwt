package main


import "fmt"


const default_name = "World"


const (
    spa_prefix = "Hola"
    eng_prefix = "Hello"
    fre_prefix = "Bonjour"
)

func Hello(name, language string) string {
    if name == "" {
        name = default_name
    }
    var prefix string

    switch language {
    case "Spanish":
        prefix = spa_prefix
    case "English":
        prefix = eng_prefix
    case "French":
        prefix = fre_prefix
    default:
        prefix = eng_prefix
    }

    s := fmt.Sprintf("%s, %s", prefix, name)
    return s
}

func main() {
    fmt.Println(Hello("world", "English"))
}
