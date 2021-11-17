package service

import "myapp/entitty"

type Customer_service interface {
	Save(entitty.Customer) entitty.Customer
	View() []entitty.Customer
}

type customer_service struct {
	customer []entitty.Customer
}

func New() Customer_service {
	return &customer_service{
		customer: []entitty.Customer{},
	}
}

func (s *customer_service) Save(customer entitty.Customer) entitty.Customer {
	s.customer = append(s.customer, customer)
	return customer
}
func (s *customer_service) View() []entitty.Customer {
	return s.customer
}

/*type Saving struct{
   detail entitty.Customer
}
type Viewing struct{
	details entitty.Customer
}
func (s *Saving) Save(entitty.Customer) entitty.Customer{
	return entitty.Customer
}
func ( v *Viewing) View() []entitty.Customer{
	return []entitty.Customer
}
func New() Customer_service{
return &Saving{}
return & Viewing{}
}*/
