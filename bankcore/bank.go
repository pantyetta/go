package bank

// 顧客情報
type Customer struct {
	Name    string
	Address string
	Phone   string
}

// 顧客アカウント
type Account struct {
	Customer
	Number  int32
	Balance float64
}
