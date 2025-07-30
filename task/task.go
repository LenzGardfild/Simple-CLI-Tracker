package task

type Task struct {
	ID int `json:"id"`
	Text string `json:"text"`
	Status string `json:"status"`
}