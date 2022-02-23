package main

import (
	"fmt"
	"github.com/spoto/go-tutorial/ch4/github"
	"log"
	"os"
	"time"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Printf("\nLess than a month old:\n")
	now := time.Now()
	oneMonthAgo := now.AddDate(0, -1, 0)
	oneYearAgo := now.AddDate(-1, 0, 0)
	for _, item := range result.Items {
		if item.CreatedAt.After(oneMonthAgo) {
			print(item)
		}
	}
	fmt.Printf("\nLess than a year old:\n")
	for _, item := range result.Items {
		if !item.CreatedAt.After(oneMonthAgo) && item.CreatedAt.After(oneYearAgo) {
			print(item)
		}
	}
	fmt.Printf("\nMore than a year old:\n")
	for _, item := range result.Items {
		if !item.CreatedAt.After(oneYearAgo) {
			print(item)
		}
	}
}

func print(item *github.Issue) (int, error) {
	return fmt.Printf("#%-5d %s %9.9s %.55s\n", item.Number, item.CreatedAt.Format("[02/01/2006]"), item.User.Login, item.Title)
}
