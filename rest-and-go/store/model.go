package store

type Product struct {
	ID     int 	     	`bson:"_id"`
	Title  string        `json:"title"`
	Image  string        `json:"image"`
	Price   uint64       `json:"price"`
	Rating  uint8        `json:"rating"`
}

type products []Product

