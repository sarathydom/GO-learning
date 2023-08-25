package main

import (
	"fmt"
	"net/http"

	"github.com/justinas/nosurf"
)

func console(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("this is mille ware")
		next.ServeHTTP(w, r)
	})
}

func nosur(next http.Handler) http.Handler{
	
		csrfHandler:= nosurf.New(next);

		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: true,
			Secure: false,
			Path: "/",
			SameSite: http.SameSiteLaxMode,
		})

		return csrfHandler
		
	
}