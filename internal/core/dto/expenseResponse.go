package dto

type ExpenseResponse struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Date string `json:"billingDate"`
	Cost string `json:"cost"`
}
