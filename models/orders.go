package models

// CartCheckout struct for describing cart checkout details.
type CartCheckout struct {
	Cart    []Cart    `json:"cart"`
	Address []Address `json:"address"`
	Payment float64   `json:"payment"`
}

// ProductCheckout struct for describing product checkout details.
type ProductCheckout struct {
	UserID   int       `json:"userID"`
	Product  Product   `json:"product"`
	Address  []Address `json:"address"`
	Quantity int       `json:"quantity"`
	Payment  float64   `json:"payment"`
}

// PlaceOrder model for describing a product or cart for placing
type PlaceOrder struct {
	UserID      int   `json:"userID"`
	AddressID   int   `json:"addressID"`
	ProductID   []int `json:"productID"`
	InventoryID []int `json:"inventoryID"`
	Quantity    []int `json:"quantity"`
	CartID      []int `json:"cartID"`
}

// COD model for describing COD payment
type COD struct {
	TotalPrice float64 `json:"totalPrice"`
	Quantity   int     `json:"quantity"`
	Status     bool    `json:"status"`
}

// RazorPay model for describing RazorPay payment.
type RazorPay struct {
	TotalPrice float64 `json:"totalPrice"`
	Quantity   int     `json:"quantity"`
	Status     bool    `json:"status"`
}

// Orders model for describing order details of a user.
type Orders struct {
	ID            int     `json:"orderID"`
	UserID        int     `json:"userID"`
	Product       Product `json:"product"`
	Quantity      int     `json:"quantity"`
	Status        string  `json:"status"`
	CartID        int     `json:"cartID"`
	PaymentStatus bool    `json:"paymentStatus"`
	PaymentType   string  `json:"paymentType"`
	SoldPrice     float64 `json:"soldPrice"`
}

// Payment model for payments
type Payment struct {
	COD      COD      `json:"cod"`
	RazorPay RazorPay `json:"razorPay"`
}

// Order struct for describing order details.
type Order struct {
	ID          int `json:"orderID"`
	UserID      int `json:"userID"`
	ProductID   int `json:"productID"`
	InventoryID int `json:"inventoryID"`
	Quantity    int `json:"quantity"`
}
