package models

// Campaigns struct
type Campaigns struct {
	Campaigns map[string]Campaign
}

// Campaign struct
type Campaign struct {
	ID         string
	Price      float64
	Content    Content
	Countries  []string
	Devices    []string
	Placements []string
}

// Content struct
type Content struct {
	Title       string
	Description string
	Landing     string
}
