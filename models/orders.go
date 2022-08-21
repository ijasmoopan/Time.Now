package models

// CartCheckout struct for describing cart checkout details.
type CartCheckout struct {
	Cart    []Cart    `json:"cart"`
	Address []Address `json:"address"`
	// Payment float64   `json:"payment"`
}

// ProductCheckout struct for describing product checkout details.
type ProductCheckout struct {
	UserID   int       `json:"userID"`
	Product  Product   `json:"product"`
	Address  []Address `json:"address"`
	Quantity int       `json:"quantity"`
	// Payment  float64   `json:"payment"`
}

// PlaceOrder model for describing a product or cart for placing
type PlaceOrder struct {
	UserID      int     `json:"userID"`
	AddressID   int     `json:"addressID"`
	ProductID   []int   `json:"productID"`
	InventoryID []int   `json:"inventoryID"`
	Quantity    int     `json:"quantity"`
	CartID      int     `json:"cartID"`
	TotalPrice  float64 `json:"totalPrice"`
	PaymentID   int     `json:"paymentID"`
}

// COD model for describing COD payment
type COD struct {
	PaymentStatus bool `json:"status"`
}

// RazorPay model for describing RazorPay payment.
type RazorPay struct {
	PaymentStatus bool `json:"status"`
}

// Orders model for describing order details of a user.
type Orders struct {
	ID          int     `json:"orderID"`
	UserID      int     `json:"userID"`
	Product     Product `json:"product"`
	Quantity    int     `json:"quantity"`
	Payment     Payment `json:"payment"`
	TotalPrice  float64 `json:"totalPrice"`
	Status      string  `json:"status"`
	CartID      int     `json:"cartID"`
	OrderedAt   string  `json:"orderedAt"`
	DeliveredAt string  `json:"deliveredAt"`
}

// Payment model for payments
type Payment struct {
	ID            int      `json:"paymentID"`
	UserID        int      `json:"userID"`
	Amount        *float64 `json:"amount"`
	PaymentType   *string  `json:"paymentType"`
	PaymentStatus bool     `json:"paymentStatus"`
}

// PaymentResponse model for payments
type PaymentResponse struct {
	UserID      int      `json:"userID"`
	TotalPrice  float64  `json:"totalPrice"`
	OfferPrice  float64  `json:"offerPrice"`
	Savings     float64  `json:"savings"`
	PaymentType []string `json:"paymentType"`
}

// Order struct for describing order details.
type Order struct {
	ID          int `json:"orderID"`
	UserID      int `json:"userID"`
	ProductID   int `json:"productID"`
	InventoryID int `json:"inventoryID"`
	Quantity    int `json:"quantity"`
}

// PaymentRequest model for payments
type PaymentRequest struct {
	UserID    *int `json:"userID"`
	ProductID *int `json:"productID"`
}
