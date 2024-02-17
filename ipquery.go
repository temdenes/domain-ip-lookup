package main

import (
	"fmt"
	"net"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			domain := r.FormValue("domain")
			ips, err := net.LookupIP(domain)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}

			fmt.Fprint(w, `<html><body>`)
			for _, ip := range ips {
				fmt.Fprintf(w, "<p>IP address: %v</p>", ip)
			}
			fmt.Fprint(w, `</body></html>`)
		} else {
			fmt.Fprint(w, `<html><body><form method="POST">
				<input type="text" name="domain" placeholder="Enter a domain name">
				<input type="submit" value="Get IPs">
			</form></body></html>`)
		}
	})

	http.ListenAndServe(":8080", nil)
}