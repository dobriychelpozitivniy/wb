package model

type Message struct {
	OrderUID          string   `json:"order_uid" validate:"nonzero, nonnil"`
	TrackNumber       string   `json:"track_number" validate:"nonzero, nonnil"`
	Entry             string   `json:"entry" validate:"nonzero, nonnil"`
	Delivery          Delivery `json:"delivery" validate:"nonzero, nonnil"`
	Payment           Payment  `json:"payment" validate:"nonzero, nonnil"`
	Items             []Item   `json:"items" validate:"nonzero, nonnil"`
	Locale            string   `json:"locale" validate:"nonzero, nonnil"`
	InternalSignature string   `json:"internal_signature" validate:"nonzero, nonnil"`
	CustomerID        string   `json:"customer_id" validate:"nonzero, nonnil"`
	DeliveryService   string   `json:"delivery_service" validate:"nonzero, nonnil"`
	ShardKey          string   `json:"shardkey" validate:"nonzero, nonnil"`
	SmID              int      `json:"sm_id" validate:"nonzero, nonnil"`
	DateCreated       string   `json:"date_created" validate:"nonzero, nonnil"`
	OofShard          string   `json:"oof_shard" validate:"nonzero, nonnil"`
}

type Delivery struct {
	Name    string `json:"name" validate:"nonzero, nonnil"`
	Phone   string `json:"phone" validate:"nonzero, nonnil"`
	Zip     string `json:"zip" validate:"nonzero, nonnil"`
	City    string `json:"city" validate:"nonzero, nonnil"`
	Address string `json:"address" validate:"nonzero, nonnil"`
	Region  string `json:"region" validate:"nonzero, nonnil"`
	Email   string `json:"email" validate:"nonzero, nonnil"`
}

type Payment struct {
	Transaction  string `json:"transaction" validate:"nonzero, nonnil"`
	RequestID    string `json:"request_id" validate:"nonzero, nonnil"`
	Currency     string `json:"currency" validate:"nonzero, nonnil"`
	Provider     string `json:"provider" validate:"nonzero, nonnil"`
	Amount       int    `json:"amount" validate:"nonzero, nonnil"`
	PaymentDT    int    `json:"payment_dt" validate:"nonzero, nonnil"`
	Bank         string `json:"bank" validate:"nonzero, nonnil"`
	DeliveryCost int    `json:"delivery_cost" validate:"nonzero, nonnil"`
	GoodsTotal   int    `json:"goods_total" validate:"nonzero, nonnil"`
	CustomFee    int    `json:"custom_fee" validate:"nonzero, nonnil"`
}

type Item struct {
	ChrtID      int    `json:"chrt_id" validate:"nonzero, nonnil"`
	TrackNumber string `json:"track_number" validate:"nonzero, nonnil"`
	Price       int    `json:"price" validate:"nonzero, nonnil"`
	RID         string `json:"rid" validate:"nonzero, nonnil"`
	Name        string `json:"name" validate:"nonzero, nonnil"`
	Sale        int    `json:"sale" validate:"nonzero, nonnil"`
	Size        string `json:"size" validate:"nonzero, nonnil"`
	TotalPrice  int    `json:"total_price" validate:"nonzero, nonnil"`
	NMID        int    `json:"nm_id" validate:"nonzero, nonnil"`
	Brand       string `json:"brand" validate:"nonzero, nonnil"`
	Status      int    `json:"status" validate:"nonzero, nonnil"`
}
