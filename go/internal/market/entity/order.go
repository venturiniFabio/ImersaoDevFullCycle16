package entity

type Order struct {
	Id            string
	Investor      *Investor
	Asset         *Asset
	Shares        int
	Price         float64
	PendingShares int
	OrderType     string
	Status        string
	Transaction   []*Transaction
}

func NewOrder(id string, investor *Investor, asset *Asset, shares int, price float64, orderType string) *Order {
	return &Order{
		Id:            id,
		Investor:      investor,
		Asset:         asset,
		Shares:        shares,
		Price:         price,
		PendingShares: shares,
		OrderType:     orderType,
		Status:        "OPEN",
		Transaction:   []*Transaction{},
	}
}
