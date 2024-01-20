package Customer

type CustomerModel struct {
	CustomerName string `json:"customer_name"`
	Age          uint   `json:"age"`
	Address      string `json:"address"`
}
