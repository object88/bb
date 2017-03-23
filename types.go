package main

type Source struct {
	BundleName string  `json:"source"`
	Map        *string `json:"map"`
	Priority   *int    `json:"priority"`
}
