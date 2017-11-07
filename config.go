package main

import (
	"encoding/json"
)

const configFile = ".deltarc"

// Config represents user-provided options (via the file specified by configFile)
type Config struct {
	Context           *int     `json:"context"`
	ShowEmpty         *bool    `json:"showEmpty"`
	ShouldCollapse    *bool    `json:"shouldCollapse"`
	Highlight         *bool    `json:"highlight"`
	UnmodifiedOpacity *float32 `json:"unmodifiedOpacity"`
	DiffFontSize      *int32   `json:"diffFontSize"`
}

func loadConfig() (config Config, err error) {
	cfg := `{
  "context": 9,
  "showEmpty": true,
  "shouldCollapse": false,
  "highlight": false,
  "unmodifiedOpacity": 0.8,
  "diffFontSize": 12
}
`
	err = json.Unmarshal([]byte(cfg), &config)
	return
}
