package main

import (
    "fmt"
    "strings"
    "unicode"
)

func main() {
    f := func(c rune) bool {
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
    }
    fmt.Printf("Correct fields are: %q", strings.FieldsFunc("  [User;ID''Transaction'',...Time;;...", f))
}
