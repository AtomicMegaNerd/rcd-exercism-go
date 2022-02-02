package robotname

import (
	"errors"
)

// Sadly necessary for this exercise?
var nameDatabase map[string]bool

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var numbers = []rune("0123456789")

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		name, err := getNextAvailableName()
		if err == nil {
			r.name = name
			return name, nil
		}
		return "", err
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	if nameDatabase != nil {
		r.name = ""
		r.Name()
	}
}

func getNextAvailableName() (string, error) {
	if nameDatabase == nil {
		genNames()
	}

	for candidateName, taken := range nameDatabase {
		if !taken {
			nameDatabase[candidateName] = true
			return candidateName, nil
		}
	}
	return "", errors.New("no more names left")
}

func genNames() {
	nameDatabase = make(map[string]bool)
	newName := make([]rune, 5)
	for _, l1 := range letters {
		for _, l2 := range letters {
			for _, l3 := range numbers {
				for _, l4 := range numbers {
					for _, l5 := range numbers {
						newName = []rune{l1, l2, l3, l4, l5}
						nameDatabase[string(newName)] = false
					}
				}
			}
		}
	}
}
