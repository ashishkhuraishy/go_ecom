package delivery

// Delivery is the common struct for info
// regarding delivery boys
type Delivery struct {
	ID         int
	Name       string
	PhoneNo    string
	IsActive   bool
	Location   string
	Incentives float32
}
