
package util
import (
	"net/http"
)

func SadResp(err error, status int, w http.ResponseWriter) {
	http.Error(w, err.Error(), status)
}
