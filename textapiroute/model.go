package textapiroute

type InputRequest struct {
	Text string `json:"text" validate:"required"`
}

// WordCount holds word and count pair
type WordCounts struct {
	Word  string
	Count int
}
