package handlers

import "net/http"

func Root(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Write([]byte("Root / handler"))
}

func Upload(w http.ResponseWriter, r *http.Request) {
	// TODO
	w.Write([]byte("Upload /upload handler"))
}
