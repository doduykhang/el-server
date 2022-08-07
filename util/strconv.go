package util

import (
	"strconv"
)

func ID(id int64) string {
	return strconv.FormatInt(id, 10)
}

func IDFromStr(id string) (uint, error) {
	u64, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		return 0, err
	}
	ID := uint(u64)
	return ID, nil
}
