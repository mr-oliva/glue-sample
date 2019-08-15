package repositories

import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/bookun/glue-sample/services/serviceA/controllers"
	"github.com/pkg/errors"
)

type File struct {
	path string
}

func NewFile(path string) (controllers.Repository, error) {
	return &File{path: path}, nil
}

func (f *File) GetAllUsers() ([]string, error) {
	fp, err := os.Open(f.path)
	if err != nil {
		return []string{}, err
	}
	defer fp.Close()
	users := []string{}
	sc := csv.NewReader(fp)
	for {
		line, err := sc.Read()
		if err != nil {
			break
		}
		users = append(users, line[1])
	}
	return users, nil
}

func (f *File) GetUserNameById(id int) (string, error) {
	fp, err := os.Open(f.path)
	if err != nil {
		return "", err
	}
	defer fp.Close()
	sc := csv.NewReader(fp)
	for {
		line, err := sc.Read()
		if err != nil {
			break
		}
		userID, err := strconv.Atoi(line[0])
		if err != nil {
			return "", err
		}
		if id == userID {
			return line[1], nil
		}
	}
	err = errors.Wrap(nil, "not found id")
	return "", err
}
