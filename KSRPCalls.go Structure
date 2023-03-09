package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type acquirer_datas struct {
	Bank_transaction_id string `json:"bank_transaction_id"`
}
type entities struct {
	Id                 string         `json:"id"`
	Entity             string         `json:"entity"`
	Amount             int            `json:"amount"`
	Currency           string         `json:"currency"`
	Base_amount        int            `json:"base_amount"`
	Status             string         `json:"status"`
	Order_id           string         `json:"order_id"`
	Invoice_id         string         `json:"invoice_id"`
	International      bool           `json:"international"`
	Method             string         `json:"method"`
	Amount_refunded    int            `json:"amount_refunded"`
	Amount_transferred int            `json:"amount_transferred"`
	Refund_status      string         `json:"refund_status"`
	Captured           bool           `json:"captured"`
	Description        string         `json:"description"`
	Card_id            string         `json:"card_id"`
	Bank               string         `json:"bank"`
	Wallet             string         `json:"wallet"`
	Vpa                string         `json:"vpa"`
	Email              string         `json:"email"`
	Contact            string         `json:"contact"`
	Notes              string         `json:"notes"`
	Fee                int            `json:"fee"`
	Tax                int            `json:"tax"`
	Error_code         string         `json:"error_code"`
	Error_description  string         `json:"error_description"`
	Error_source       string         `json:"error_source"`
	Error_step         string         `json:"error_step"`
	Error_reason       string         `json:"error_reason"`
	Acquirer_data      acquirer_datas `json:"acquirer_data"`
}
type payments struct {
	Entity     entities `json:"entity"`
	Created_at string   `json:"created_at"`
}
type payloads struct {
	Payment payments `json:"payment"`
}
type input1 struct {
	Entity     string   `json:"entity"`
	Account_id string   `json:"account_id"`
	Event      string   `json:"event"`
	Contains   []string `json:"contains"`
	Payload    payloads `json:"payload"`
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's Homepage"))
}

func Ksrp(w http.ResponseWriter, req *http.Request) {

	if req.Method != http.MethodPost {
		w.Header().Set("Allow", "POST")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Method Not allowed"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var t input1
	err := decoder.Decode(&t)
	if err != nil {
		panic(err)
	}

	log.Println("1", t.Entity)

	log.Println("2", t.Account_id)
	log.Println("3", t.Event)

	log.Println("4", t.Contains)

	/*
		body, err := io.ReadAll(req.Body)
		if err != nil {
			w.Write([]byte("Call Failed"))
			return
		}

		var file *os.File
		options := os.O_CREATE | os.O_APPEND
		log.Println("log0")
		file, err = os.OpenFile("read1.txt", options, os.FileMode(0644))
		log.Println("log1")
		if err != nil {
			w.Write([]byte("Call Failed"))
			return
		}

		_, err = file.Write([]byte(string(body)))
		if err != nil {
			w.Write([]byte("Call Failed"))
			return
		}

		_, err = file.Write([]byte("\n"))
		if err != nil {
			w.Write([]byte("Call Failed"))
			return
		}

		err = file.Close()
		if err != nil {
			w.Write([]byte("Call Failed"))
			return
		}*/

	log.Println("Call Successful")

	w.Write([]byte("Successful"))
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/ksrp", Ksrp)

	log.Println("Starting server on :8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		var w http.ResponseWriter
		w.Write([]byte("Call Failed"))
		return

	}

}
