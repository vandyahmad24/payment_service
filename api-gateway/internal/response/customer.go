package response

type LoginCustomerResponse struct {
	Token string `json:"token"`
}

type RegisterCustomerResponse struct {
	CustomerId string `json:"customer_id"`
}
