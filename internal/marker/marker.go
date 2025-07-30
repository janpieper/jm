package marker

import (
	"os"
	"path/filepath"
)

const Marker = ".jm"

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

func Current(alternatives []string) (*string, error) {
	files, err := List(alternatives)
	if err != nil {
		return nil, err
	}

	if len(files) > 0 {
		file := files[0]
		return &file, nil
	}

	return nil, nil
}

func List(alternatives []string) ([]string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	var files []string
	markers := append([]string{Marker}, alternatives...)

	for {
		for _, marker := range markers {
			file := filepath.Join(dir, marker)
			if _, err := os.Stat(file); err == nil {
				files = append(files, filepath.Dir(file))
				break
			}
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			break
		}

		dir = parent
	}

	return files, nil
}
