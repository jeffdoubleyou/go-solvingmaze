package models

type Status struct {
	Success  bool     `json:"success"`
	Messages []string `json:"messages,omitempty"`
}
