package controllers

import (
	"net/http"

	"github.com/Xebec19/simple-bank/internal/db"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type AccountController struct {
	queries *db.Queries
}

func NewAccountController(queries *db.Queries) *AccountController {
	return &AccountController{
		queries: queries,
	}
}

type CreateAccountRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Balance   int32  `json:"balance" binding:"required,min=0"`
	Currency  string `json:"currency" binding:"required,oneof=rupee dollar"`
}

type AccountResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Balance   int32  `json:"balance"`
	Currency  string `json:"currency"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_at"`
}

// CreateAccount godoc
// @Summary Create a new account
// @Description Create a new bank account with the provided details
// @Tags accounts
// @Accept json
// @Produce json
// @Param account body CreateAccountRequest true "Account details"
// @Success 201 {object} AccountResponse
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /accounts [post]
func (ac *AccountController) CreateAccount(c *gin.Context) {
	var req CreateAccountRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	params := db.CreateAccountParams{
		FirstName: pgtype.Text{String: req.FirstName, Valid: true},
		LastName:  pgtype.Text{String: req.LastName, Valid: true},
		Balance:   pgtype.Int4{Int32: req.Balance, Valid: true},
		Currency:  db.NullCurrency{Currency: db.Currency(req.Currency), Valid: true},
	}

	account, err := ac.queries.CreateAccount(c.Request.Context(), params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create account"})
		return
	}

	accountUUID, _ := uuid.FromBytes(account.ID.Bytes[:])

	response := AccountResponse{
		ID:        accountUUID.String(),
		FirstName: account.FirstName.String,
		LastName:  account.LastName.String,
		Balance:   account.Balance.Int32,
		Currency:  string(account.Currency.Currency),
		Status:    string(account.Status.Status),
		CreatedAt: account.CreatedAt.Time.String(),
	}

	c.JSON(http.StatusCreated, response)
}

// GetAccounts godoc
// @Summary Get all accounts
// @Description Retrieve a list of all bank accounts
// @Tags accounts
// @Produce json
// @Success 200 {array} AccountResponse
// @Failure 500 {object} map[string]interface{}
// @Router /accounts [get]
func (ac *AccountController) GetAccounts(c *gin.Context) {
	accounts, err := ac.queries.GetAccounts(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
		return
	}

	var response []AccountResponse
	for _, account := range accounts {
		accountUUID, _ := uuid.FromBytes(account.ID.Bytes[:])
		response = append(response, AccountResponse{
			ID:        accountUUID.String(),
			FirstName: account.FirstName.String,
			LastName:  account.LastName.String,
			Balance:   account.Balance.Int32,
			Currency:  string(account.Currency.Currency),
			Status:    string(account.Status.Status),
			CreatedAt: account.CreatedAt.Time.String(),
		})
	}

	c.JSON(http.StatusOK, response)
}
