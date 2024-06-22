package handler

import (
	"html/template"
	"log"
	"net/http"
	"thr/model"
	"thr/node"
)

// PeminjamanReadAllHandler handles the request to get all peminjaman records.
func PeminjamanReadAllHandler(w http.ResponseWriter, r *http.Request) {
	// Get all peminjaman records from the model
	peminjamanList := model.GetAllPeminjaman()

	// Debug: Log the number of peminjaman records fetched
	log.Printf("Fetched %d peminjaman records\n", len(peminjamanList))

	// Parse the HTML template
	tmpl, err := template.ParseFiles("view/ViewPeminjaman.html")
	if err != nil {
		log.Printf("Error parsing template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Create the data to pass to the template
	data := struct {
		PeminjamanList []node.PeminjamanBuku
	}{
		PeminjamanList: peminjamanList,
	}

	// Debug: Log data being sent to template
	log.Printf("Data sent to template: %+v\n", data)

	// Execute the template with the data
	if err := tmpl.Execute(w, data); err != nil {
		log.Printf("Error executing template: %v\n", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
func PeminjamanUpdateStatusHandler() {

}
