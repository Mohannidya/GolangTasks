package models

type commonModel struct {
	Name          string `json:"item_name`
	order_id      string `json:"order_id"`
	customer_name string `json:"customer_name"`
	address       string `json:"address"`
	item          string `json:"item"`
	size          string `json:"size"`
	status        string `json:"status"`
	created_at    string `json:"created_at"`
	updated_at    string `json:"updated_at"`
}
