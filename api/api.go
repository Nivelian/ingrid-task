package api

import (
	"errors"
	"fmt"
	"github.com/Nivelian/ingrid-task/helpers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

func routesHandler(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	err := errors.New("incorrect params")

	srcs := q["src"]
	if srcs == nil || len(srcs) == 0 || len(srcs) > 1 {
		BadRequestErr(helpers.LogErr(err, "Provide 1 'src' param"), w)
		return
	}
	src := srcs[0]

	dsts := q["dst"]
	if dsts == nil || len(dsts) == 0 {
		BadRequestErr(helpers.LogErr(err, "Provide at least 1 'dst' param"), w)
		return
	}

	routes, err := GetRoutes(src, dsts)
	if err != nil {
		InternalServerErr(err, w)
		return
	}

	if err := SendStruct(routes, w); err != nil {
		InternalServerErr(err, w)
	}
}

func StartServer(port string) {
	router := mux.NewRouter()

	router.HandleFunc("/routes", routesHandler).Queries("src", "{src}", "dst", "{dst}").Methods(http.MethodGet)

	logrus.Info(fmt.Sprintf("Starting server on port %v", port))
	logrus.Fatal(http.ListenAndServe(port, router))
}
