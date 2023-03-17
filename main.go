package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	//_ "github.com/denisenkom/go-mssqldb"
	"io"
	"log"
	"net/http"
	"os"
)

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
		log.Println("log2")

		w.Write([]byte("Call Failed"))
		log.Println("log3")

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
	}

	log.Println("Call Successful")

	w.Write([]byte("Successful"))
}

func VerifyWebhookSignature(requestBody string, webhookSignature string, webhookSecret string) bool {
	body := []byte(requestBody)

	isValid := VerifySignature(body, webhookSignature, webhookSecret)

	return isValid
}

func VerifySignature(body []byte, signature string, key string) bool {
	h := hmac.New(sha256.New, []byte(key))
	h.Write(body)
	expectedSignature := hex.EncodeToString(h.Sum(nil))
	if expectedSignature != signature {
		return false
	}
	return true
}

func main() {
	webhookbody := `{"entity":"event","account_id":"acc_HVFD0PFnHPAzKj","event":"payment.authorized","contains":["payment"],"payload":{"payment":{"entity":{"id":"pay_JUEM4c0pSLpFEW","entity":"payment","amount":12300,"currency":"INR","status":"authorized","order_id":"order_JUELuT6cFaHkvd","invoice_id":null,"international":false,"method":"netbanking","amount_refunded":0,"refund_status":null,"captured":false,"description":"#JUELZ1z1EC0pwi","card_id":null,"bank":"SBIN","wallet":null,"vpa":null,"email":"prathmesh.bijjargi@razorpay.com","contact":"+917411714931","notes":[],"fee":null,"tax":null,"error_code":null,"error_description":null,"error_source":null,"error_step":null,"error_reason":null,"acquirer_data":{"bank_transaction_id":"6416615"},"created_at":1652339804}}},"created_at":1652339806}`
	signature := "ea962f3a2a8090aa14ef10fc178306edcb46ee7ccadae8afe2275cda26f3d2a6"
	secret := "123456"
	b := VerifyWebhookSignature(webhookbody, signature, secret)
	fmt.Println(b)

}
