package types

type Bucket struct {
	ID       string `json:"id,omitempty"`
	Name     string `json:"name"`
	Language string `json:"language"`
}

type WordUsage struct {
	Meaning  string   `json:"meaning"`
	Examples []string `json:"examples"`
}

type Card struct {
	ID          string    `json:"id,omitempty"`
	Word        string    `json:"word"`
	HiddenChars []int     `json:"hiddenChars"`
	Usage       WordUsage `json:"usage"`
	Level       int       `json:"level"`
}
