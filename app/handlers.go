package app

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Customer struct {
	ID      int    `json:"ID" xml:"Id"`
	Name    string `json:"Name" xml:"Name"`
	City    string `json:"City" xml:"City"`
	ZipCode string `json:"ZipCode" xml:"zipcode"`
}

var customers []Customer = []Customer{
	{1, "User1", "Jakarta", "12345"},
	{2, "User2", "Bandung", "67890"},
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello Celerates!")
}

func getAllCustomers(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprint(w, "get customers endpoint!")
	if r.Header.Get("Content-type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		// json.NewEncoder(w).Encode(customers)

	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}
}

func getCustomer(w http.ResponseWriter, r *http.Request) {

	//*get route variable
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	//*convert string to int
	id, err := strconv.Atoi(customerId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "invalid customer id")
	}
	//*searching customer data
	var cust Customer

	for _, data := range customers {
		if data.ID == id {
			cust = data
		}
	}

	if cust.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "customer data not found")
		return
	}
	//*return customer data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)
}

func addCustomer(w http.ResponseWriter, r *http.Request) {
	var cust Customer
	json.NewDecoder(r.Body).Decode(&cust)

	//*get last id
	nextID := getNextID()
	cust.ID = nextID

	//*save data to array
	customers = append(customers, cust)

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "customer sucessfully created")
}

func getNextID() int {
	lastIndex := len(customers) - 1
	lastCustomer := customers[lastIndex]

	return lastCustomer.ID + 1
}

func getDeleteCust(w http.ResponseWriter, r *http.Request) {

	//*get route variable
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	//*convert string to int
	id, _ := strconv.Atoi(customerId)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprintf(w, "invalid customer id")
	// }
	//*searching customer data
	var cust Customer

	for index, data := range customers {
		if data.ID == id {
			customers = append(customers[:index], customers[index+1:]...)
			return
		}
	}
	// if cust.ID == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprintf(w, "customer data not found")
	// 	return
	// }
	//*return customer data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)
}

func getPutCust(w http.ResponseWriter, r *http.Request) {

	//*get route variable
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	//*convert string to int
	id, _ := strconv.Atoi(customerId)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprintf(w, "invalid customer id")
	// }
	//*searching customer data
	var cust Customer

	for index, data := range customers {
		if data.ID == id {
			customers = append(customers[:index], customers[index+1:]...)
			var customerData Customer
			_ = json.NewDecoder(r.Body).Decode(&customerData)
			customerData.ID = id
			customers = append(customers, customerData)
			json.NewEncoder(w).Encode(customerData)

		}
	}

	// if cust.ID == 0 {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	fmt.Fprintf(w, "customer data not found")
	// 	return
	// }
	//*return customer data
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cust)
}
