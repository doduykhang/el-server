package middlewares

import (
	"net/http"
	"strconv"

	"el.com/m/util"
)

func EnsureAuthenticatedJwtMw() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			c, err := r.Cookie("session_token")
			if err != nil {
				if err == http.ErrNoCookie {
					// If the cookie is not set, return an unauthorized status
					util.SadResp(err, http.StatusUnauthorized, w)
					return
				}
				// For any other type of error, return a bad request status
				util.SadResp(err, http.StatusUnauthorized, w)
				return
			}
			sessionToken := c.Value
			n, err := strconv.ParseInt(sessionToken, 10, 64)
			if err == nil {
				util.SadResp(err, http.StatusUnauthorized, w)
				return
			}
			r = util.RequestWithUserID(r, n)
			next.ServeHTTP(w, r)
		})
	}
}
