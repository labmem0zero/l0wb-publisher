package main

import (
	"encoding/json"
	"fmt"
	stan "github.com/nats-io/stan.go"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type Order struct {
	OrderUid    string `json:"order_uid" db:"order_uid"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Entry       string `json:"entry" db:"entry"`
	Delivery    Delivery `json:"delivery"`
	Payment 	Payment `json:"payment"`
	Items 		[]Item `json:"items"`
	Locale            string    `json:"locale" db:"locale"`
	InternalSignature string    `json:"internal_signature" db:"internal_signature"`
	CustomerId        string    `json:"customer_id" db:"customer_id"`
	DeliveryService   string    `json:"delivery_service" db:"delivery_service"`
	Shardkey          string    `json:"shardkey" db:"shardkey"`
	SmId              int       `json:"sm_id" db:"sm_id"`
	DateCreated       string `json:"date_created" db:"date_created"`
	OofShard          string    `json:"oof_shard" db:"oof_shard"`
}

type Delivery    struct {
	Name    string `json:"name" db:"name"`
	Phone   string `json:"phone" db:"phone"`
	Zip     string `json:"zip" db:"zip"`
	City    string `json:"city" db:"city"`
	Address string `json:"address" db:"address"`
	Region  string `json:"region" db:"region"`
	Email   string `json:"email" db:"email"`
}

type Payment struct {
	Transaction  string `json:"transaction" db:"transaction"`
	RequestId    string `json:"request_id" db:"request_id"`
	Currency     string `json:"currency" db:"currency"`
	Provider     string `json:"provider" db:"provider"`
	Amount       int    `json:"amount" db:"amount"`
	PaymentDt    int    `json:"payment_dt" db:"payment_dt"`
	Bank         string `json:"bank" db:"bank"`
	DeliveryCost int    `json:"delivery_cost" db:"delivery_cost"`
	GoodsTotal   int    `json:"goods_total" db:"goods_total"`
	CustomFee    int    `json:"custom_fee" db:"custom_fee"`
}

type Item struct {
	ChrtId      int    `json:"chrt_id" db:"chrt_id"`
	TrackNumber string `json:"track_number" db:"track_number"`
	Price       int    `json:"price" db:"price"`
	Rid         string `json:"rid" db:"rid"`
	Name        string `json:"name" db:"name"`
	Sale        int    `json:"sale" db:"sale"`
	Size        string `json:"size" db:"size"`
	TotalPrice  int    `json:"total_price" db:"total_price"`
	NmId        int    `json:"nm_id" db:"nm_id"`
	Brand       string `json:"brand" db:"brand"`
	Status      int    `json:"status" db:"status"`
}

func Sender(sc stan.Conn){
	for t:=range time.NewTicker(1*time.Second).C{
		jsonstruct:=Order{
			"OrderUid_"+strconv.Itoa(int(t.Unix())),
			"TrackNumber_"+strconv.Itoa(int(t.Unix())),
			"Entry_"+strconv.Itoa(int(t.Unix())),
			Delivery{
				"Name_"+strconv.Itoa(int(t.Unix())),
				"Phone"+strconv.Itoa(int(t.Unix())),
				"Zip"+strconv.Itoa(int(t.Unix())),
				"City"+strconv.Itoa(int(t.Unix())),
				"Address_"+strconv.Itoa(int(t.Unix())),
				"Region_"+strconv.Itoa(int(t.Unix())),
				"Email_"+strconv.Itoa(int(t.Unix())),
			},
			Payment{
				"Transaction_"+strconv.Itoa(int(t.Unix())),
				"RequestId_"+strconv.Itoa(int(t.Unix())),
				"Currency_"+strconv.Itoa(int(t.Unix())),
				"Provider_"+strconv.Itoa(int(t.Unix())),
				rand.Intn(99999999),
				rand.Intn(99999999),
				"Bank_"+strconv.Itoa(int(t.Unix())),
				rand.Intn(99999999),
				rand.Intn(99999999),
				rand.Intn(99999999),
			},
			[]Item{},
			"Locale_"+strconv.Itoa(int(t.Unix())),
			"IS_"+strconv.Itoa(int(t.Unix())),
			"Cid_"+strconv.Itoa(int(t.Unix())),
			"DS_"+strconv.Itoa(int(t.Unix())),
			"SK_"+strconv.Itoa(int(t.Unix())),
			rand.Intn(99999999),
			time.Now().String(),
			"OS_"+strconv.Itoa(int(t.Unix())),
		}
		for i:=1;i<4;i++{
			item:=Item{
				rand.Intn(99999999),
				"TrackNumber_"+strconv.Itoa(int(t.Unix())),
				rand.Intn(99999999),
				"Rid_"+strconv.Itoa(int(t.Unix())),
				"Name_"+strconv.Itoa(int(t.Unix())),
				rand.Intn(99999999),
				"Size_"+strconv.Itoa(int(t.Unix())),
				rand.Intn(99999999),
				rand.Intn(99999999),
				"Brand_"+strconv.Itoa(int(t.Unix())),
				rand.Intn(99999999),
			}
			jsonstruct.Items=append(jsonstruct.Items,item)
		}
		data,_:=json.Marshal(jsonstruct)
		fmt.Println(jsonstruct.OrderUid)
		err:=sc.Publish("orders", data)
		if err!=nil{
			log.Println("Publish error:",err)
		}
	}
}

func main() {
	sc, _ := stan.Connect("test-cluster", "go-nats-streaming-json-publisher")
	go Sender(sc)
	time.Sleep(110*time.Minute)
}
