package internal

import (
    "runtime"
    "strings"
)

func Compare(a, b string) int {
    return strings.Compare(normalize(a), normalize(b))
}

func normalize(s string) string {
    return strings.TrimRight(s, lineBreak())
}

func lineBreak() string {
    if runtime.GOOS == "windows" {
        return "\r\n"
    } else {
        return "\n"
    }
}
