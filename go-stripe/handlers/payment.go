package handlers

import (
	"encoding/json"
	"fmt"
	"go-stripe/models"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
	"github.com/stripe/stripe-go/v78/webhook"
	"gorm.io/gorm"
)

type PaymentHandler struct {
	DB *gorm.DB
}

var StripeSecretKey string 

func InitStripe() {
	StripeSecretKey = os.Getenv("STRIPE_SECRET_KEY")
	stripe.Key = StripeSecretKey
}

func (h *PaymentHandler) ShowSendMoney(c *gin.Context) {
	c.HTML(http.StatusOK, "send.html", gin.H{"StripePublishableKey": os.Getenv("STRIPE_PUBLISHABLE_KEY")})
}

func (h *PaymentHandler) CreateCheckoutSession(c *gin.Context) {
	userID := c.MustGet("user_id").(uint)
	recipientEmail := c.PostForm("recipient_email")
	amount := c.PostForm("amount")

	if recipientEmail == "" || amount == "" {
		c.HTML(http.StatusBadRequest, "send.html", gin.H{"error":"Recipient email and amount are requiured"})
		return
	}

	transaction := models.Transaction {
		UserID: userID,
		RecipientEmail: recipientEmail,
		Amount: convertToCents(amount),
		Currency: "usd",
		Status: "pending",
	}

	if err := h.DB.Create(&transaction).Error; err != nil {
		c.HTML(http.StatusInternalServerError, "send.html", gin.H{"error": "Failed to create transaction"})
		return
	}

	// Create Stripe Checkout Session
    params := &stripe.CheckoutSessionParams{
        PaymentMethodTypes: stripe.StringSlice([]string{
            "card",
        }),
        LineItems: []*stripe.CheckoutSessionLineItemParams{
            {
                PriceData: &stripe.CheckoutSessionLineItemPriceDataParams{
                    Currency: stripe.String("usd"),
                    ProductData: &stripe.CheckoutSessionLineItemPriceDataProductDataParams{
                        Name: stripe.String(fmt.Sprintf("Payment to %s", recipientEmail)),
                    },
                    UnitAmount: stripe.Int64(transaction.Amount),
                },
                Quantity: stripe.Int64(1),
            },
        },
        Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
        SuccessURL: stripe.String(fmt.Sprintf("http://localhost:8080/success?session_id={CHECKOUT_SESSION_ID}&transaction_id=%d", transaction.ID)),
        CancelURL:  stripe.String("http://localhost:8080/dashboard"),
        Metadata: map[string]string{
            "transaction_id": fmt.Sprintf("%d", transaction.ID),
        },
    }

    s, err := session.New(params)
    if err != nil {
        c.HTML(http.StatusInternalServerError, "send.html", gin.H{
            "error": err.Error(),
        })
        return
    }

    // Update transaction with Stripe session ID
    h.DB.Model(&transaction).Update("stripe_payment_id", s.ID)

    c.HTML(http.StatusOK,"send.html", gin.H{
        "sessionId": s.ID,
    })
}

func (h *PaymentHandler) HandleSuccess(c *gin.Context) {
    sessionID := c.Query("session_id")
    transactionID := c.Query("transaction_id")
    
    if sessionID == "" {
        c.Redirect(http.StatusFound, "/dashboard")
        return
    }

    // Retrieve the session from Stripe
    s, err := session.Get(sessionID, nil)
    if err != nil {
        c.HTML(http.StatusOK, "success.html", gin.H{
            "Success": false,
            "Message": "Payment verification failed",
        })
        return
    }

    // Update transaction status
    var transaction models.Transaction
    if result := h.DB.First(&transaction, transactionID); result.Error == nil {
        if s.PaymentStatus == stripe.CheckoutSessionPaymentStatusPaid {
            h.DB.Model(&transaction).Update("status", "completed")
        }
    }

    c.HTML(http.StatusOK, "success.html", gin.H{
        "Success": true,
        "Message": "Payment completed successfully!",
    })
}

func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
    const MaxBodyBytes = int64(65536)
    payload, err := io.ReadAll(io.LimitReader(c.Request.Body, MaxBodyBytes))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
        c.Status(http.StatusServiceUnavailable)
        return
    }

    event, err := webhook.ConstructEvent(payload, c.GetHeader("Stripe-Signature"), 
        os.Getenv("STRIPE_WEBHOOK_SECRET"))
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
        c.Status(http.StatusBadRequest)
        return
    }

    switch event.Type {
    case "checkout.session.completed":
        var session stripe.CheckoutSession
        err := json.Unmarshal(event.Data.Raw, &session)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing session: %v\n", err)
            c.Status(http.StatusBadRequest)
            return
        }
        
        // Update transaction status
        var transaction models.Transaction
        if result := h.DB.Where("stripe_payment_id = ?", session.ID).First(&transaction); result.Error == nil {
            h.DB.Model(&transaction).Update("status", "completed")
            fmt.Printf("Transaction %d completed successfully\n", transaction.ID)
        }
        
    case "checkout.session.expired":
        var session stripe.CheckoutSession
        err := json.Unmarshal(event.Data.Raw, &session)
        if err != nil {
            fmt.Fprintf(os.Stderr, "Error parsing session: %v\n", err)
            c.Status(http.StatusBadRequest)
            return
        }
        
        // Update transaction status
        var transaction models.Transaction
        if result := h.DB.Where("stripe_payment_id = ?", session.ID).First(&transaction); result.Error == nil {
            h.DB.Model(&transaction).Update("status", "expired")
            fmt.Printf("Transaction %d expired\n", transaction.ID)
        }
    }

    c.Status(http.StatusOK)
}

func convertToCents(amount string) int64 {
    // Simple conversion - in production, use proper decimal handling
    var dollars, cents int64
    fmt.Sscanf(amount, "%d.%d", &dollars, &cents)
    return dollars*100 + cents
}