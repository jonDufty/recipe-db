package session

import (
	"context"
	"errors"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/99designs/go/stringutil"
	"github.com/jonDufty/recipes/auth/models/user"
	"github.com/jonDufty/recipes/common/database"
	"github.com/russross/meddler"

	"github.com/oklog/ulid/v2"
)

const SESSION_AUTH_LIFETIME time.Duration = time.Minute * 15
const COOKIE_NAME string = "sunday_sesh"

type Session struct {
	ID        string    `json:"id" meddler:"id"`
	UserID    int       `json:"user_id" meddler:"user_id"`
	CreatedAt time.Time `json:"created_at" meddler:"created_at"`
	ExpiresAt time.Time `json:"expires_at" meddler:"expires_at"`
	IP        string    `json:"ip" meddler:"ip"`
}

func CreateSession(w http.ResponseWriter, r *http.Request, u *user.User) error {

	id := ulid.MustNew(1234, strings.NewReader(stringutil.Random(20)))

	cookie := &http.Cookie{
		Name:    COOKIE_NAME,
		Expires: time.Now().Add(SESSION_AUTH_LIFETIME),
		Path:    "/",
		Value:   id.String(),
	}

	session := &Session{
		ID:        id.String(),
		UserID:    u.ID,
		CreatedAt: time.Now(),
		ExpiresAt: cookie.Expires,
		IP:        r.RemoteAddr,
	}
	err := database.Insert(r.Context(), "session", session)
	if err != nil {
		return errors.New("Failed to add new sessions. " + err.Error())
	}

	http.SetCookie(w, cookie)

	return nil
}

type Uid string

const UidKey Uid = "uid"

func Middleware() func(handler http.Handler) http.Handler {

	return func(next http.Handler) http.Handler {

		f := func(w http.ResponseWriter, r *http.Request) {
			if err := CheckCookie(w, r); err != nil {
				http.Redirect(w, r, "/login", http.StatusUnauthorized)
			}
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(f)
	}
}

func CheckCookie(w http.ResponseWriter, r *http.Request) error {
	cookie, err := r.Cookie(COOKIE_NAME)
	if err != nil {
		return err
	}

	s, err := getUserFromSession(r.Context(), cookie.Value)
	if err != nil {
		return errors.New("cannot find user. " + err.Error())
	}

	if s.ExpiresAt.Before(time.Now()) {
		log.Print("Session has expired")
		clearCookie(w, r)
		return errors.New("session has expired")
	}
	return nil

}

func clearCookie(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:    COOKIE_NAME,
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	})
}

func getUserFromSession(ctx context.Context, sid string) (*Session, error) {
	tx, err := database.StartTx(ctx)
	if err != nil {
		return nil, err
	}

	s := &Session{}
	err = meddler.QueryRow(tx, s, "SELECT * FROM session WHERE id = ?", sid)
	if err != nil {
		tx.Rollback()
		return nil, errors.New("Cannot find session. " + err.Error())
	} else {
		tx.Commit()
	}

	return s, nil
}
