package utils

import "github.com/google/uuid"

func GetStrId(idLen int) string {
	uuidString := uuid.New().String()
	if idLen > len(uuidString) {
		idLen = len(uuidString)
	}
	return uuidString[:idLen]
}
