//Filename: cmd/api/entries.go

package main

import (
	"fmt"
	"net/http"
)

// createEntryHandler for the "POST" /v1/entries" endpoint
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new entry.....")
}

// showEntryHandlerfor the "GET" /v1/entries/:id" endpoint
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	// Displays the Entries ID
	fmt.Fprintf(w, "show the details for entry %d\n", id)
}
