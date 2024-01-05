package types

type Account struct {
    ID          int    `json:"id"`
    AccountName string `json:"account_name"`
    Balance     float64 `json:"balance"`
}

type Category struct {
    ID           int    `json:"id"`
    CategoryName string `json:"category_name"`
}

type Transaction struct {
    ID          int     `json:"id"`
    Amount      float64 `json:"amount"`
    Date        string  `json:"date"`
    Description string  `json:"description"`
    AccountID   int     `json:"account_id"`
    CategoryID  int     `json:"category_id"`
}
