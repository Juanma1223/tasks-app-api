package task

type Task struct {
	Id     int    `json:"id"`
	Text   string `json:"text"`
	ListId int    `json:"listId"`
}
