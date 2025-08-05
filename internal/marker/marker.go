package marker

import (
	"os"
	"path/filepath"
)

const Marker = ".jm"

type MarkerMatch struct {
	Path   string
	Marker string
}

func Exists() bool {
	stat, _ := os.Stat(Marker)
	return stat != nil
}

func Create() error {
	return os.WriteFile(Marker, []byte{}, 0666)
}

func Remove() error {
	return os.Remove(Marker)
}

func Current(alternatives []string) (*MarkerMatch, error) {
	matches, err := List(alternatives)
	if err != nil {
		return nil, err
	}

	if len(matches) > 0 {
		match := matches[0]
		return &match, nil
	}

	return nil, nil
}

func List(alternatives []string) ([]MarkerMatch, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var matches []MarkerMatch
	markers := append([]string{Marker}, alternatives...)

	for {
		for _, marker := range markers {
			file := filepath.Join(dir, marker)
			if _, err := os.Stat(file); err == nil {
				matches = append(matches, MarkerMatch{filepath.Dir(file), marker})
				break
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}

		dir = parent
	}

	return matches, nil
}
