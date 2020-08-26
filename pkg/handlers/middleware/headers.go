package middleware

import "net/http"

//AddCommonHeaders to request
func AddCommonHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		res.Header().Add("content-type", "application/json")
		next.ServeHTTP(res, req)
	})
}
