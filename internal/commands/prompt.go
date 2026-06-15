package commands

import (

	"charm.land/huh/v2"
)

func RepoUrl() (string, error) {
	var url string

	err := huh.NewInput().
		Title("enter url to clone").
		Value(&url).
		Run()
	if err != nil {
		return "", err
	}

	return url, nil
}
