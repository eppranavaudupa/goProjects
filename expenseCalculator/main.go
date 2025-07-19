package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type Expenses struct {
	Amount   float64
	Category string
	Note     string
	Date     time.Time
}

var expense []Expenses

func main() {
	for {
		fmt.Printf("Select The Options\n")
		fmt.Println("1.View Expense")
		fmt.Println("2.Add Expense")
		fmt.Println("3.Sort Items")
		fmt.Println("4.Save")
		fmt.Println("5.close")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			viewExp(expense)
		case 2:
			expense = addExp(expense)
		case 3:
			sortExp(expense)
		case 4:
			saveExp(expense)
		case 5:
			closeExp()
			return
		default:
			fmt.Println("Invalid Option")
		}
	}

}
func viewExp(expense []Expenses) {
	if len(expense) == 0 {
		fmt.Printf("No Expense \n")
		return
	}
	for i, exp := range expense {
		fmt.Printf("%d. ₹%.2f | %s | %s | %s\n", i+1, exp.Amount, exp.Category, exp.Note, exp.Date.Format("2006-01-02 15:04"))
	}
}
func addExp(expense []Expenses) []Expenses {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter The Amount: ")
	var amount float64
	fmt.Scan(&amount)
	fmt.Println("Enter The Category: ")
	category, _ := reader.ReadString('\n')

	fmt.Println("NOTE : ")
	note, _ := reader.ReadString('\n')

	expen := Expenses{
		Amount:   amount,
		Category: strings.TrimSpace(category),
		Note:     strings.TrimSpace(note),
		Date:     time.Now(),
	}
	expense = append(expense, expen)
	return expense
}
func sortExp(expense []Expenses) {
	if len(expense) == 0 {
		fmt.Println("No Expense To sort")
		return
	}
	sort.Slice(expense, func(i, j int) bool {
		return expense[i].Amount > expense[j].Amount
	})
	viewExp(expense)
}
func saveExp(expense []Expenses){
	pdf,err:=os.Create("Expense.pdf")
	if err!=nil{
		fmt.Printf("Error while creating file",err)
		return
	}
	defer pdf.Close()
	for _,i:= range expense{
		line:=fmt.Sprintf("₹%.2f | %s | %s | %s\n",i.Amount, i.Category, i.Note, i.Date.Format("2006-01-02 15:04"))
		pdf.WriteString(line)
	}
}
func closeExp(){
	fmt.Println("CLOSE.")
	return
}