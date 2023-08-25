package models

type Todo struct {
	ID        int    `json:"id"`
	Task_name string `json:"task_name"`
	Task      string `json:"task"`
	Status    bool   `json:"status"`
}
