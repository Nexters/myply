package member

import "strings"

type Keywords []string

type Member struct {
	DeviceToken string
	Name        string
	Keywords    Keywords
}

func (ks Keywords) ToString() string {
	return strings.Join(ks, ",")
}
