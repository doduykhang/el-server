
package util

import (
	"strconv"
)

func ID(id int64) string {
	return strconv.FormatInt(id, 10)
}

func IDFromStr(id string) (int64, error) {
	return strconv.ParseInt(id, 10, 64)
}
