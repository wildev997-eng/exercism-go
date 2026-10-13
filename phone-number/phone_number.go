package phonenumber

import (
	"errors"
	"fmt"
	"strings"
)

func Number(phoneNumber string) (string, error) {
	var removed string
	if len(phoneNumber) < 10 {
		return "", errors.New("the number must be 10 digit long")
	}
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			removed += string(char)
		}
	}
	if removed[0] == '1' {
		removed = removed[1:]
	}
	if len(removed) != 10 || !strings.Contains("23456789", string(removed[0])) || !strings.Contains("23456789", string(removed[3])) {
		return "", errors.New("Not valid phone number")
	}
	return removed, nil
}

func AreaCode(phoneNumber string) (string, error) {
	var removed string
	if len(phoneNumber) < 10 {
		return "", errors.New("the number must be 10 digit long")
	}
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			removed += string(char)
		}
	}
	if removed[0] == '1' {
		removed = removed[1:]
	}
	if len(removed) != 10 || !strings.Contains("23456789", string(removed[0])) || !strings.Contains("23456789", string(removed[3])) {
		return "", errors.New("Not valid phone number")
	}
	return removed[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	var removed string
	if len(phoneNumber) < 10 {
		return "", errors.New("the number must be 10 digit long")
	}
	for _, char := range phoneNumber {
		if char >= '0' && char <= '9' {
			removed += string(char)
		}
	}
	if removed[0] == '1' {
		removed = removed[1:]
	}
	if len(removed) != 10 || !strings.Contains("23456789", string(removed[0])) || !strings.Contains("23456789", string(removed[3])) {
		return "", errors.New("Not valid phone number")
	}
	return fmt.Sprintf("(%s) %s-%s", removed[0:3], removed[3:6], removed[6:]), nil
}
