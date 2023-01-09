package middleware

import (
	"fmt"
	"github.com/gorilla/mux"
)

func InjectMiddlewareToNamedRoute(router *mux.Router, routeNames ...string) func(middlewareFuncs ...mux.MiddlewareFunc) error {
	return func(middlewareFuncs ...mux.MiddlewareFunc) error {

		for _, routeName := range routeNames {
			for _, mw := range middlewareFuncs {
				route := router.GetRoute(routeName)
				if route == nil {
					return fmt.Errorf("router is not exist")
				}
				route.Handler(
					mw(route.GetHandler()),
				)
			}
		}

		return nil
	}
}
