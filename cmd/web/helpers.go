package main

import (
	"net/http"
	"runtime/debug"
  "fmt"
)

// the serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func(app *application) serverError(w http.ResponseWriter, err error) {
  trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
  app.errorLog.Output(2,trace)
  http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// the clientError helper sends a specific code and corresponding description
// to the use.
func(app *application) clientError(w http.ResponseWriter, status int) {
  http.Error(w, http.StatusText(status), status)
}

func(app *application) notFound(w http.ResponseWriter) {
  app.clientError(w, http.StatusNotFound)
}
