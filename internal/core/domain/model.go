package domain

type Message struct {
	ID     string `json:"id"`
	Body   string `json:"body"`
	Status string `json:"status"`
}
