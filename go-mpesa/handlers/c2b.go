package handlers

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"go-mpesa/db"
	"go-mpesa/models"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// expected request payload from frontend
type C2BRequest struct {
	Phone  string  `json:"phone"`
	Amount float64 `json:"amount"`
}

// mpesa stk push
type stkPushRequest struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}

// mpesa response
type stkPushResponse struct {
	MerchantRequestID string `json:"MerchantRequestID"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
	ResponseCode      string `json:"ResponseCode"`
	ResponseDesc      string `json:"ResponseDescription"`
	CustomerMsg       string `json:"CustomerMessage"`
}

// init stk push
func C2BHandler(c *gin.Context) {
	var req C2BRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Inavlid input"})
		return
	}

	// Generate timetamp
	timestamp := time.Now().Format("20060102150405")

	// password = Base64(business short code + passkey + timestamp)
	shortCode := os.Getenv("MPESA_SHORTCODE")
	passKey := os.Getenv("MPESA_PASSKEY")
	password := encodePassword(shortCode, passKey, timestamp)

	// Get OAuth token
	token, err := GetAccessToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get token"})
		return
	}

	// build stk push payload
	payload := stkPushRequest{
		BusinessShortCode: shortCode,
		Password:          password,
		Timestamp:         timestamp,
		TransactionType:   "CustomerPayBillOnline",
		Amount:            fmt.Sprintf("%.0f", req.Amount),
		PartyA:            req.Phone,
		PartyB:            shortCode,
		PhoneNumber:       req.Phone,
		CallBackURL:       os.Getenv("MPESA_CALLBACK_URL"),
		AccountReference:  "TestAccount",
		TransactionDesc:   "Payment test",
	}

	body, _ := json.Marshal(payload)

	// cal mpesa API
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	httpReq, _ := http.NewRequest("POST", url, bytes.NewBuffer(body))
	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(httpReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "stk push request failed"})
		return
	}
	defer resp.Body.Close()

	var stkResp stkPushResponse
	if err := json.NewDecoder(resp.Body).Decode(&stkResp); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Decode failed"})
		return
	}

	// persist to db
	tx := models.MpesaTransaction{
		TransactionType: "C2B",
		Amount:          req.Amount,
		PhoneNumber:     req.Phone,
		Status:          "pending",
		CheckoutID:      stkResp.CheckoutRequestID,
		Description:     stkResp.CustomerMsg,
	}
	db.DB.Create(&tx)

	c.JSON(http.StatusOK, stkResp)
}

// --- Common type for metadata items ---
type CallbackItem struct {
	Name  string      `json:"Name"`
	Value interface{} `json:"Value"`
}

// --- Callback (Safaricom will POST here) ---
type stkCallback struct {
	Body struct {
		StkCallback struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			ResultDesc        string `json:"ResultDesc"`
			CallbackMetadata  struct {
				Item []CallbackItem `json:"Item"`
			} `json:"CallbackMetadata"`
		} `json:"stkCallback"`
	} `json:"Body"`
}

func extractReceipt(items []CallbackItem) string {
	for _, it := range items {
		if it.Name == "MpesaReceiptNumber" {
			return fmt.Sprintf("%v", it.Value)
		}
	}
	return ""
}

func C2BCallback(c *gin.Context) {
	var callback stkCallback
	if err := c.BindJSON(&callback); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid callback"})
		return
	}

	cb := callback.Body.StkCallback

	// Update transaction
	var tx models.MpesaTransaction
	if err := db.DB.Where("checkout_id = ?", cb.CheckoutRequestID).First(&tx).Error; err != nil {
		log.Println("Transaction not found")
		return
	}

	if cb.ResultCode == 0 {
		tx.Status = "Success"
		tx.ReceiptNumber = extractReceipt(cb.CallbackMetadata.Item)
	} else {
		tx.Status = "Failed"
	}
	tx.Description = cb.ResultDesc
	db.DB.Save(&tx)

	c.JSON(http.StatusOK, gin.H{"message": "callback received"})
}

func encodePassword(shortCode, passkey, timestamp string) string {
	raw := shortCode + passkey + timestamp
	return base64.StdEncoding.EncodeToString([]byte(raw))
}
