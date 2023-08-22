package models

type Todo struct {
	ID        int    `json:"id, min=1, max=1000"`
	Task_name string `json:"task_name"`
	Task      string `json:"task"`
	Status    bool   `json:"status"`
}
