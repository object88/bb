package main

// Source describes the shape of entries in the loaded manifest.json file
type Source struct {
	CSS       *string `json:"css"`
	CSSMap    *string `json:"cssMap"`
	Priority  *int    `json:"priority"`
	Source    string  `json:"source"`
	SourceMap *string `json:"sourceMap"`
}
