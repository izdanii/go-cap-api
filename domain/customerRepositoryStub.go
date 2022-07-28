package domain

type CustomerRepositoryStub struct {
	Customer []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.Customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "User1", "Jakarta", "12345", "29-08-2022", "1"},
		{"2", "User2", "Bandung", "67890", "18-09-2020", "1"},
		{"3", "User3", "Surabaya", "91827", "22-07-2021", "1"},
		{"4", "User4", "Yogyakarta", "09816", "13-03-2000", "1"},
	}
	return CustomerRepositoryStub{Customer: customers}
}
