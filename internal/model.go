package internal


type Card struct {
    Question, Answer string
}

func NewCard(q, a string) *Card {
    return &Card { q, a }
}
