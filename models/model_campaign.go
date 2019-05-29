package models

type campaigns struct {
	Campaigns map[string]campaign
}

type campaign struct {
	ID         string
	Price      float64
	Content    content
	Countries  []string
	Devices    []string
	Placements []string
}

type content struct {
	Title       string
	Description string
	Landing     string
}
