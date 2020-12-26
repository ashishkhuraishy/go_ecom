package user

var (
	// Users is a variable which stores the data
	// of all the users
	Users []User
)

// Create is used to create a user and add to the
// database
func (u *User) Create() error {
	return nil
}

// Get will populate the user with values
// based on the user with the id if any present
func (u *User) Get() error {
	return nil
}

// Update will updated the user with
// the given id on the db
func (u *User) Update() error {
	return nil
}

// Delete will delete the user with the given
// id
func (u *User) Delete() error {
	return nil
}

// GetAll will return a list of users
// found on the db
func (u *User) GetAll() error {
	return nil
}

// GetAllWith will return the list of users
func (u *User) GetAllWith(category string) error {
	return nil
}
