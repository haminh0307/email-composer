package entity

type Customer struct {
	Title     string `json:"title"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func (customer *Customer) ToSliceOfStrings() []string {
	return []string{
		customer.Title,
		customer.FirstName,
		customer.LastName,
		customer.Email,
	}
}
