package main

type Source struct {
	CSS       *string `json:"css"`
	CSSMap    *string `json:"cssMap"`
	Priority  *int    `json:"priority"`
	Source    string  `json:"source"`
	SourceMap *string `json:"sourceMap"`
}
