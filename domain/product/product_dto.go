package product

// Product defines the items in a restaurent
// that the users can buy
type Product struct {
	ID           int64   `json:"id"`
	RestaurentID int64   `json:"restaurent_id"`
	Name         string  `json:"name"`
	Veg          bool    `json:"veg"`
	Available    bool    `json:"available"`
	Price        float32 `json:"price"`
}
