package restful

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

// ISession
type ISession interface {
	Set(key, value interface{}) error
	Get(key interface{}) interface{}
	Del(key interface{}) error
	SID() string
}

// ISessionProvider
type ISessionProvider interface {
	SessionInit(sid string) (ISession, error)
	SessionRead(sid string) (ISession, error)
	SessionDestroy(sid string) error
	SessionGC(maxAge int64)
}

// SessionManager
type SessionManager struct {
	cookieName string
	maxAge     int64
	provider   ISessionProvider
	lock       sync.Mutex
}

// SessionManager.GenerateSID
func (s *SessionManager) GenerateSID() string {
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(b)
}

// SessionManager.SessionStart
func (s *SessionManager) SessionStart(w http.ResponseWriter, r *http.Request) (session ISession) {
	s.lock.Lock()
	defer s.lock.Unlock()
	cookie, err := r.Cookie(s.cookieName)
	if err != nil || cookie.Value == "" {
		sid := s.GenerateSID()
		session, _ = s.provider.SessionInit(sid)
		newCookie := http.Cookie{
			Name:     s.cookieName,
			Value:    url.QueryEscape(sid),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(s.maxAge),
		}
		http.SetCookie(w, &newCookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		if session, _ = s.provider.SessionRead(sid); session == nil {
			session, _ = s.provider.SessionInit(sid)
			newCookie := http.Cookie{
				Name:     s.cookieName,
				Value:    url.QueryEscape(sid),
				Path:     "/",
				HttpOnly: true,
				MaxAge:   int(s.maxAge),
			}
			http.SetCookie(w, &newCookie)
		}
	}
	return
}

// SessionManager.SessionDestory
func (s *SessionManager) SessionDestory(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(s.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	s.provider.SessionDestroy(cookie.Value)
	expiredTime := time.Now()
	newCookie := http.Cookie{
		Name:     s.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  expiredTime,
		MaxAge:   -1,
	}
	http.SetCookie(w, &newCookie)
}

// SessionManager.SessionGC
func (s *SessionManager) SessionGC() {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.provider.SessionGC(s.maxAge)
	time.AfterFunc(time.Duration(s.maxAge), func() {
		s.SessionGC()
	})
}

// NewSessionManager
func NewSessionManager(provider ISessionProvider, cookieName string, maxAge int64) *SessionManager {
	return &SessionManager{
		cookieName: cookieName,
		maxAge:     maxAge,
		provider:   provider,
	}
}
