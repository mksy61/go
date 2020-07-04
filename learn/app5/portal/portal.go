//Package portal for learning
package portal

import (
	"fmt"
	"net/http"
)

//RunWebPortal function
func RunWebPortal(addr string) error {
	http.HandleFunc("/", rootHandler)
	err := http.ListenAndServe(addr, nil)
	return err
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to portal\n\n")
	fmt.Fprintf(w, "Server\t:%s\n", r.Host)
	fmt.Fprintf(w, "Client\t:%s\n", r.RemoteAddr)
	fmt.Fprintf(w, "%s\n", r.UserAgent())
}
