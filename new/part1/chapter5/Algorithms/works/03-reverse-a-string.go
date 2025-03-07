package works

func Reverse(word string) string {
    letters := []byte(word)
    n := len(letters)
    var reversed []byte

    for i := n-1; i >= 0; i--{
        reversed = append(reversed, letters[i])
    }

    return string(reversed)
}

func Reverse2(word string) string {
    n := len(word)
    var reversed string
    for i := n-1; i >=0; i--{
        reversed += string(word[i])
    }

    return reversed
}

func Reverse3(word string) string {
    var reversed string
    for _, r := range word{
        reversed = string(r) + reversed
    }

    return reversed
}
