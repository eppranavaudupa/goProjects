package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sort"
	"time"
)

type Expenses struct {
	Amount   float64   `json:amount`
	Category string    `json:category`
	Note     string    `json:note`
	Date     time.Time `json:time`
}

var expense []Expenses

func main() {
	http.HandleFunc("/expenses", handleExpenses)
	http.HandleFunc("/expenses/sorted", handleSortedExpenses)
	http.HandleFunc("/expenses/save", handleSave)
	fmt.Println("Server is running at port:8010")
	http.ListenAndServe(":8010", nil)
}
func handleExpenses(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Header().Set("content-type", "application/json")
		json.NewEncoder(w).Encode(expense)
	} else if r.Method == http.MethodPost {
		var exp Expenses
		err := json.NewDecoder(r.Body).Decode(&exp)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}
		exp.Date = time.Now()
		expense = append(expense, exp)
		w.WriteHeader(http.StatusCreated)
		fmt.Fprint(w, "Expense Added")
	}
}
func handleSortedExpenses(w http.ResponseWriter, r *http.Request) {
	if len(expense) == 0 {
		http.Error(w, "No expense to sort", http.StatusNotFound)
	}
	sortedExpense := make([]Expenses, len(expense))
	copy(sortedExpense, expense)
	sort.Slice(sortedExpense, func(i, j int) bool {
		return (sortedExpense[i].Amount > sortedExpense[j].Amount)
	})
	w.Header().Set("content-Type", "application/json")
	json.NewEncoder(w).Encode(sortedExpense)
}
func handleSave(w http.ResponseWriter, r *http.Request) {
	file, err := os.Create("expenseList.pdf")
	if err != nil {
		http.Error(w, "Error While Creating the File", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	for _, e := range expense {
		line := fmt.Sprintf("₹%.2f | %s | %s | %s\n", e.Amount, e.Category, e.Note, e.Date.Format("2006-01-02 15:04"))
		file.WriteString(line)

	}
	fmt.Println("file saaved Successfully")

}
