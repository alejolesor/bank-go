package handler

import (
	"VERITRAN/internal/domain/account/entities"
	"VERITRAN/internal/platform/dto"
	"VERITRAN/internal/usecases/account"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/httplog"
)

type Operations struct {
	depositUseCase    account.Account
	createUseCase     account.AccountCreate
	withDrawalUseCase account.AccountWithDraw
	transferUseCase   account.AccountTransfer
}

func NewOperations(depositUseCase account.Account, createUseCase account.AccountCreate, withDrawaluseCase account.AccountWithDraw, transferUseCase account.AccountTransfer) *Operations {
	return &Operations{
		depositUseCase:    depositUseCase,
		createUseCase:     createUseCase,
		withDrawalUseCase: withDrawaluseCase,
		transferUseCase:   transferUseCase,
	}
}

func (b *Operations) Deposite(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())
	account := dto.Deposit{}

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Err(err).Msg("failed decoder body")
		http.Error(w, "Request invalid,failed decoded body", http.StatusBadRequest)
		return
	}
	newBalance, err := b.depositUseCase.Deposit(account.NickName, account.Amount)
	if err != nil {
		log.Err(err).Msg("failed deposit")
		http.Error(w, "Error doing deposit : "+err.Error(), http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("Success of Operations, new Available Balance %v", newBalance)))
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (b *Operations) Create(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())
	account := dto.Account{}
	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Err(err).Msg("failed decoder body")
		http.Error(w, "Request invalid,failed decoded body", http.StatusBadRequest)
		return
	}

	errCreate := b.createUseCase.Create(entities.NewAccount(account.NickName, account.Name, account.LastName, account.Tipe, account.Amount))
	if errCreate != nil {
		log.Err(err).Msg("failed create account")
		http.Error(w, "Error doing deposit : "+err.Error(), http.StatusBadRequest)
		return
	}

}

func (b *Operations) WithDrawal(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())
	account := dto.WithDrawal{}

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Err(err).Msg("failed decoder body")
		http.Error(w, "Request invalid,failed decoded body", http.StatusBadRequest)
		return
	}
	newBalance, err := b.withDrawalUseCase.WithDrawall(account.NickName, account.WithDrawal)
	if err != nil {
		log.Err(err).Msg("failed withdrawal")
		http.Error(w, "Error doing withdrawal : "+err.Error(), http.StatusNotFound)
		return
	}
	_, err = w.Write([]byte(fmt.Sprintf("Success of Operations, new Available Balance %v", newBalance)))
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

}

func (b *Operations) Transfer(w http.ResponseWriter, r *http.Request) {
	log := httplog.LogEntry(r.Context())
	account := dto.Transfer{}

	err := json.NewDecoder(r.Body).Decode(&account)
	if err != nil {
		log.Err(err).Msg("failed decoder body")
		http.Error(w, "Request invalid,failed decoded body", http.StatusBadRequest)
		return
	}

	transfer := entities.NewTransfer(account.Depositor.NickName, account.Depositor.Amount, account.Receiver.NickName)

	result, err := b.transferUseCase.Transfer(*transfer)
	if err != nil {
		log.Err(err).Msg("failed NewTransfer")
		http.Error(w, "Error doing Transfer : "+err.Error(), http.StatusNotFound)
		return
	}

	_, err = w.Write([]byte(fmt.Sprintf("Confirmed, %s", *result)))
	if err != nil {
		log.Err(err).Msg("failed write body")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

}
