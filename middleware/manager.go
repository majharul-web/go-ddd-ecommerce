package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	globalMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globalMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middleware ...Middleware) *Manager {
	mngr.globalMiddlewares = append(mngr.globalMiddlewares, middleware...)
	return mngr
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	for _, middleware := range mngr.globalMiddlewares {
		n = middleware(n)
	}
	return n
	
	
}