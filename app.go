package main

import (
	"log"
	"net/http"
	"os"
)

func handle(w http.ResponseWriter, r *http.Request) {
  path := r.URL.Path;
  if len(path) < 2 {
    w.Write([]byte("hello world"));
    return;
  }

  http.Redirect(w, r, "https://cdn.jsdelivr.net/gh/BlueHtml/img" + r.URL.Path, http.StatusFound);//重定向
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/", handle);

	log.Println("listening on", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
