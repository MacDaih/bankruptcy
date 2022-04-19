package http

import "net/http"

type HttpServer struct {
	Get func(http.ResponseWriter, *http.Request)
	Add func(http.ResponseWriter, *http.Request)
}

func fourOfour(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}

func ServeHTTP(
	port string,
	get func(http.ResponseWriter, *http.Request),
	add func(http.ResponseWriter, *http.Request),
	e chan error,
) {
	http.HandleFunc("/get", get)
	http.HandleFunc("/add", add)

	http.HandleFunc("/", fourOfour)

	e <- http.ListenAndServe(port, nil)
}
