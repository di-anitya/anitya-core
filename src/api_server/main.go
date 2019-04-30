package apiserver

import (
	"github.com/urfave/negroni"
)

func main() {
	router := NewRouter()
	n := negroni.Classic()
	n.UseHandler(router)
	n.Run(":8080")
}
