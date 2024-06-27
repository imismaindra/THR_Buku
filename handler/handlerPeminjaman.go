package handler

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
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
func PeminjamanUpdateStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		// Parse form values
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		// Get form values
		peminjamanIDStr := r.URL.Path[len("/peminjaman/update/"):]
		peminjamanID, err := strconv.Atoi(peminjamanIDStr)
		if err != nil {
			http.Error(w, "Invalid peminjaman ID", http.StatusBadRequest)
			return
		}

		newStatusStr := r.FormValue("newStatus")
		newStatus, err := strconv.Atoi(newStatusStr)
		if err != nil {
			http.Error(w, "Invalid new status", http.StatusBadRequest)
			return
		}

		// Call the model function to update peminjaman status
		success := model.UpdatePeminjamanStatus(peminjamanID, newStatus)
		if !success {
			http.Error(w, "Failed to update peminjaman status", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Status berhasil diperbarui"))
	} else {
		// If not POST method, return method not allowed
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// PeminjamanDetailHandler handles the request to get detail peminjaman.
func PeminjamanDetailHandler(w http.ResponseWriter, r *http.Request) {
	peminjamanIDStr := r.URL.Path[len("/peminjaman/detail/"):]
	peminjamanID, err := strconv.Atoi(peminjamanIDStr)
	if err != nil {
		http.Error(w, "Invalid peminjaman ID", http.StatusBadRequest)
		return
	}

	// Call the model function to get peminjaman detail
	detailPeminjaman, err := model.GetPeminjamanDetail(peminjamanID)
	if err != nil {
		http.Error(w, "Failed to get peminjaman detail", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(detailPeminjaman)
}
func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode JSON request body
	var requestData struct {
		UserId   string   `json:"userId"`
		UserName string   `json:"userName"`
		Cart     []string `json:"cart"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	// Parse userId to integer
	memberId, err := strconv.Atoi(requestData.UserId)
	if err != nil {
		http.Error(w, "Invalid member ID", http.StatusBadRequest)
		return
	}

	// Convert cart items from []string to []int
	var cartItems []int
	for _, itemIdStr := range requestData.Cart {
		itemId, err := strconv.Atoi(itemIdStr)
		if err != nil {
			http.Error(w, "Invalid cart item ID", http.StatusBadRequest)
			return
		}
		cartItems = append(cartItems, itemId)
	}

	member := node.Member{
		User:   node.MemberNode{Id: memberId, Nama: requestData.UserName},
		Alamat: "Alamat Member",
		NoTelp: "No Telp Member",
	}

	// Insert the borrowing details into the database
	peminjaman, ok := model.InsertPeminjaman(member, cartItems)
	if !ok {
		http.Error(w, "Failed to checkout", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Checkout berhasil. ID Peminjaman: %d", peminjaman.IdPeminjaman)
}
