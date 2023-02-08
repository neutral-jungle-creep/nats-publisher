package domain

type Order struct {
	OrderUID    string `json:"order_uid," fake:"{regex:[0-9a-z]{19}}"`
	TrackNumber string `json:"track_number" fake:"{regex:[A-Z]{14}}"`
	Entry       string `json:"entry" fake:"{regex:[A-Z]{4}}"`
	Delivery    struct {
		Name    string `json:"name" fake:"{name} {lastname}"`
		Phone   string `json:"phone" fake:"{phone}"`
		Zip     string `json:"zip" fake:"{zip}"`
		City    string `json:"city" fake:"{city}"`
		Address string `json:"address" fake:"{street}"`
		Region  string `json:"region" fake:"{city}"`
		Email   string `json:"email" fake:"{email}"`
	} `json:"delivery"`
	Payment struct {
		Transaction  string `json:"transaction" fake:"{regex:[0-9a-z]{19}}"`
		RequestId    string `json:"request_id" fake:"{regex:[0-9]{10}}"`
		Currency     string `json:"currency" fake:"rub"`
		Provider     string `json:"provider" fake:"provider"`
		Amount       int    `json:"amount" fake:"{number:1, 1000000}"`
		PaymentDt    int    `json:"payment_dt" fake:"{number:1, 1000000}"`
		Bank         string `json:"bank" fake:"{name}"`
		DeliveryCost int    `json:"delivery_cost" fake:"{number:1, 1000000}"`
		GoodsTotal   int    `json:"goods_total" fake:"{number:1, 1000000}"`
		CustomFee    int    `json:"custom_fee" fake:"{number:1, 1000000}"`
	} `json:"payment"`
	Items []struct {
		ChrtId      int    `json:"chrt_id" fake:"{number:1, 1000000}"`
		TrackNumber string `json:"track_number" fake:"{regex:[A-Z]{14}}"`
		Price       int    `json:"price" fake:"{number:1, 1000000}"`
		Rid         string `json:"rid" fake:"{regex:[0-9a-z]{21}}"`
		Name        string `json:"name" fake:"{name}"`
		Sale        int    `json:"sale" fake:"{number:1, 100}"`
		Size        string `json:"size" fake:"{regex:[0-9]{2}}"`
		TotalPrice  int    `json:"total_price" fake:"{number:1, 1000000}"`
		NmId        int    `json:"nm_id" fake:"{number:1, 2389212}"`
		Brand       string `json:"brand" fake:"{name}"`
		Status      int    `json:"status" fake:"{regex:[0-1]{3}}"`
	} `json:"items"`
	Locale            string `json:"locale" fake:"en"`
	InternalSignature string `json:"internal_signature" fake:"test"`
	CustomerId        string `json:"customer_id" fake:"text"`
	DeliveryService   string `json:"delivery_service" fake:"{name}"`
	Shardkey          string `json:"shardkey" fake:"{regex:[0-9]{1}}"`
	SmId              int    `json:"sm_id" fake:"{number:1, 99}"`
	DateCreated       string `json:"date_created" fake:"2023.01.02 22:15:50"`
	OofShard          string `json:"oof_shard" fake:"{regex:[0-9]{1}}"`
}
