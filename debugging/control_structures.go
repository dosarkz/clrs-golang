package debugging

type Store struct {
	m map[string]*Customer
}

type Foo struct{}

type Customer struct {
	ID string
}

func (s *Store) storeCustomers(customers []Customer) {
	for _, customer := range customers {
		current := customer // copy
		s.m[current.ID] = &current
	}
}
