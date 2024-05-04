//go:build !go1.16

package main

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)
}

 greeting  := "Hi"

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, greeting)
}
