package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type Book struct {
	Title     string `json:"title"`
	Pages     int    `json:"pages"`
	PagesRead int    `json:"pagesRead"`
}

var library []Book

func main() {
	loadLibrary()
	fmt.Println("Welcome to BookTracker CLI tool!")
	for {
		fmt.Println("\nSelect an option:")
		fmt.Println("1. Add a book")
		fmt.Println("2. Update reading progress")
		fmt.Println("3. Check reading progress")
		fmt.Println("4. Exit")
		fmt.Print("Enter your choice: ")

		choice := getInput()

		switch choice {
		case 1:
			addBook()
		case 2:
			updateProgress()
		case 3:
			checkProgress()
		case 4:
			saveLibrary()
			fmt.Println("Exiting BookTracker CLI tool.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func getInput() int {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input, _ := strconv.Atoi(scanner.Text())
	return input
}

func addBook() {
	fmt.Println("\nAdding a new book:")
	fmt.Print("Enter the title: ")
	title := getInputString()
	fmt.Print("Enter the number of pages: ")
	pages := getInput()
	library = append(library, Book{Title: title, Pages: pages, PagesRead: 0})
	fmt.Println("Book added successfully!")
}

func updateProgress() {
	fmt.Println("\nUpdating reading progress:")
	fmt.Println("Select a book to update:")
	for i, book := range library {
		fmt.Printf("%d. %s\n", i+1, book.Title)
	}
	index := getInput() - 1
	if index < 0 || index >= len(library) {
		fmt.Println("Invalid choice.")
		return
	}
	fmt.Print("Enter the number of pages you've read: ")
	pagesRead := getInput()
	library[index].PagesRead = pagesRead
	fmt.Println("Reading progress updated successfully!")
}

func checkProgress() {
	fmt.Println("\nChecking reading progress:")
	fmt.Println("Select a book to check:")
	for i, book := range library {
		fmt.Printf("%d. %s\n", i+1, book.Title)
	}
	index := getInput() - 1
	if index < 0 || index >= len(library) {
		fmt.Println("Invalid choice.")
		return
	}
	book := library[index]
	percent := float64(book.PagesRead) / float64(book.Pages) * 100
	fmt.Printf("You've read %.2f%% of %s still in page %d\n", percent, book.Title, book.PagesRead)
}

func getInputString() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func loadLibrary() {
	file, err := os.Open("library.json")
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&library)
	if err != nil {
		fmt.Println("Error decoding library:", err)
	}
}

func saveLibrary() {
	file, err := os.Create("library.json")
	if err != nil {
		fmt.Println("Error creating library file:", err)
		return
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	err = encoder.Encode(library)
	if err != nil {
		fmt.Println("Error encoding library:", err)
	}
}

