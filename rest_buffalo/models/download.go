package models

// Helm chart download model just needs name and url
type Download struct {
	Name string `json:"name"`
	SourceLocation string `json:"source_location"`
}
