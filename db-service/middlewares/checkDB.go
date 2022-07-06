package middlewares

import (
	"net/http"

	"github.com/davidpolme/mutant-detector/db-service/db"
)

//TODO: implement checkDB middleware
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check if table is created
		c := make(chan bool)
		go isTableCreated(c)

		if condition := !<-c; condition {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"message":"Por favor espere unos segundos mientras se crea la base de datos"}`))
			next.ServeHTTP(w, r)
		} else {
			next.ServeHTTP(w, r)
		}
	}

}

func isTableCreated(c chan bool) chan bool {
	err := db.CreateTable()
	if err != nil {
		c <- true
		return c
	}
	c <- false
	return c
}
