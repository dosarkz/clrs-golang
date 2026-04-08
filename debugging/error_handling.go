package debugging

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
)

type transientError struct {
	err error
}

// Создается пользовательский
// тип transientError
func (t transientError) Error() string {
	return fmt.Sprintf("transient error: %v", t.err)
}
func getTransactionAmount(transactionID string) (float32, error) {
	if len(transactionID) != 5 {
		return 0, fmt.Errorf("id is invalid: %s",
			transactionID) // Возврат простой ошибки, если ID транзакции недействителен
	}
	amount, err := getTransactionAmountFromDB(transactionID)
	if err != nil {
		return 0, fmt.Errorf("failed to get transaction %s: %w",
			transactionID, err)
	}
	return amount, nil
	//Возврат ошибки типа transientError
	//	в случае сбоя при запросе в БД
}

func getTransactionAmountFromDB(transactionID string) (float32, error) {
	var err error
	err = fmt.Errorf("some error")
	if err != nil {
		return 0, transientError{err: err}
	}

	return 0, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	transactionID := r.URL.Query().Get("transaction")
	_, err := getTransactionAmount(transactionID)
	if err != nil {
		switch err := err.(type) {
		case transientError:
			http.Error(w, err.Error(), http.StatusServiceUnavailable)
		default:
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}
}

func getBalance(db *sql.DB, clientID string) (
	balance float32, err error) {
	rows, err := db.Query("select id from table", clientID)
	if err != nil {
		return 0, err
	}
	defer func() {
		closeErr := rows.Close()
		if err != nil { // Присвоение ошибки из rows.Close другой переменной
			//	Если ошибка уже не nil, определяем приоритет
			if closeErr != nil {
				log.Printf("failed to close rows: %v", err)
			}
			return
		}
		err = closeErr // В противном случае возвращаем closeErr
	}()
	if rows.Next() {
		err := rows.Scan(&balance)
		if err != nil {
			return 0, err
		}
		return balance, nil
	}
	return 0, sql.ErrNoRows
}
