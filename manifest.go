package main

import (
	"encoding/json"
	"io/ioutil"
	"sort"
)

func loadManifest() (map[string]Source, []Source) {
	rawJSON, err := ioutil.ReadFile("./resources/manifest.json")
	if err != nil {
		panic(err.Error())
	}

	var manifest map[string]Source
	err = json.Unmarshal(rawJSON, &manifest)
	if err != nil {
		panic(err.Error())
	}

	manifestSlice := make([]Source, len(manifest))
	offset, priorityCount := 0, 0
	for _, v := range manifest {
		manifestSlice[offset] = v
		if v.Priority != nil {
			priorityCount++
		}
		offset++
	}
	sort.Slice(manifestSlice, func(i int, j int) bool {
		if manifestSlice[i].Priority == nil {
			return false
		}
		if manifestSlice[j].Priority == nil {
			return true
		}
		return *manifestSlice[i].Priority < *manifestSlice[j].Priority
	})
	highPriorityManifest := manifestSlice[0:priorityCount]

	return manifest, highPriorityManifest
}
