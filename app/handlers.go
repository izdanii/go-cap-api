package app

import (
	"capi/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// type Customer struct {
// 	ID      int    `json:"ID" xml:"Id"`
// 	Name    string `json:"Name" xml:"Name"`
// 	City    string `json:"City" xml:"City"`
// 	ZipCode string `json:"ZipCode" xml:"zipcode"`
// }

// var customers []Customer = []Customer{
// 	{1, "User1", "Jakarta", "12345"},
// 	{2, "User2", "Bandung", "67890"},
// }

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello Celerates!")
// }
type CustomerHandler struct {
	service service.CustomerService
}

func (ch *CustomerHandler) getAllCustomers(w http.ResponseWriter, r *http.Request) {
	customerStatus := r.URL.Query().Get("status")

	customers, err := ch.service.GetAllCustomers(customerStatus)

	if err != nil {
		writeResponse(w, err.Code, err.AsMessage())
		return
	} else {
		writeResponse(w, http.StatusOK, customers)
	}
	// fmt.Fprint(w, "get customers endpoint!")
	// if r.Header.Get("Content-type") == "application/xml" {
	// 	w.Header().Add("Content-Type", "application/xml")
	// 	xml.NewEncoder(w).Encode(customers)

	// } else {
	// 	w.Header().Add("Content-Type", "application/json")
	// 	json.NewEncoder(w).Encode(customers)
	// }
}

func (ch *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {

	//*get route variable
	vars := mux.Vars(r)
	customerID := vars["customer_id"]
	customer, err := ch.service.GetCustomerByID(customerID)
	if err != nil {
		// w.WriteHeader(err.Code)
		// fmt.Fprint(w, err.Message)
		writeResponse(w, err.Code, err.AsMessage())
		return
	}

	writeResponse(w, http.StatusOK, customer)
	w.Header().Add("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(customer)
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		panic(err)
	}
}

// 	w.WriteHeader(http.StatusBadRequest)
// 	fmt.Fprintf(w, "invalid customer id")
// }

// 	//*searching customer data
// 	var cust Customer

// 	for _, data := range customers {
// 		if data.ID == id {
// 			cust = data
// 		}
// 	}

// 	if cust.ID == 0 {
// 		w.WriteHeader(http.StatusNotFound)
// 		fmt.Fprintf(w, "customer data not found")
// 		return
// 	}
// 	//*return customer data

// func addCustomer(w http.ResponseWriter, r *http.Request) {
// 	var cust Customer
// 	json.NewDecoder(r.Body).Decode(&cust)

// 	nextID := getNextID()
// 	cust.ID = nextID

// 	customers = append(customers, cust)

// 	w.WriteHeader(http.StatusCreated)
// 	fmt.Fprintln(w, "customer sucessfully created")
// }

// func getNextID() int {
// 	lastIndex := len(customers) - 1
// 	lastCustomer := customers[lastIndex]

// 	return lastCustomer.ID + 1
// }

// func getDeleteCust(w http.ResponseWriter, r *http.Request) {

// 	//*get route variable
// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]

// 	id, _ := strconv.Atoi(customerId)

// 	var cust Customer

// 	for index, data := range customers {
// 		if data.ID == id {
// 			customers = append(customers[:index], customers[index+1:]...)
// 			return
// 		}
// 	}
// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(cust)
// }

// func getPutCust(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	customerId := vars["customer_id"]

// 	id, _ := strconv.Atoi(customerId)

// 	var cust Customer

// 	for index, data := range customers {
// 		if data.ID == id {
// 			customers = append(customers[:index], customers[index+1:]...)
// 			var customerData Customer
// 			_ = json.NewDecoder(r.Body).Decode(&customerData)
// 			customerData.ID = id
// 			customers = append(customers, customerData)
// 			json.NewEncoder(w).Encode(customerData)

// 		}
// 	}

// 	w.Header().Add("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(cust)
// }
