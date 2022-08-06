package middlewares

import (
	"database/sql"
	"errors"
	"net/http"

	"el.com/m/models"
	"el.com/m/util"
	"github.com/golang-jwt/jwt/v4"
)

var (
	jwtKey = []byte("my_secret_key")
)

type Claims struct {
	ID uint
	jwt.StandardClaims
}

func EnsureAuthenticatedJwtMw(db *sql.DB, role uint) func(next http.Handler) http.Handler {
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
			claims := &Claims{}

			tkn, err := jwt.ParseWithClaims(sessionToken, claims, func(token *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})

			if err != nil {
				if err == jwt.ErrSignatureInvalid {
					util.SadResp(err, http.StatusUnauthorized, w)
				}
				util.SadResp(err, http.StatusUnauthorized, w)
				return
			}
			if !tkn.Valid {
				util.SadResp(err, http.StatusUnauthorized, w)
				return
			}

			//skip if not spicify role
			if role != util.AllRole {

				account, err := models.Accounts(
					models.AccountWhere.ID.EQ(claims.ID),
				).One(r.Context(), db)

				if err != nil {
					util.SadResp(err, http.StatusUnauthorized, w)
					return
				}

				if account.RoleID != role {
					util.SadResp(errors.New("Unauthorized"), http.StatusUnauthorized, w)
					return
				}
			}
			var id uint
			if role == util.AdminRole {
				manager, err := models.Managers(
					models.ManagerWhere.AccountID.EQ(claims.ID),
				).One(r.Context(), db)

				if err != nil {
					util.SadResp(err, http.StatusUnauthorized, w)
					return
				}

				id = manager.ID
			}

			if role == util.UserRole {
				user, err := models.Users(
					models.UserWhere.AccountID.EQ(claims.ID),
				).One(r.Context(), db)

				if err != nil {
					util.SadResp(err, http.StatusUnauthorized, w)
					return
				}

				id = user.ID
			}

			r = util.RequestWithUserID(r, id)
			next.ServeHTTP(w, r)
		})
	}
}