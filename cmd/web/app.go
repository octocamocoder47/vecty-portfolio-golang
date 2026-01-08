package main

import (
	"encoding/json"
	"net/http"
	// "vecty-portfolio/cmd/web/components"
	"vecty-portfolio/cmd/web/data"
)

func LoadResume() (*data.Resume, error) {
	resp, err := http.Get("resume.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var resume data.Resume
	err = json.NewDecoder(resp.Body).Decode(&resume)
	if err != nil {
		return nil, err
	}
	return &resume, nil
}
