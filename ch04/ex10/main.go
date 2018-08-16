// ex10 repo:golang/go is:open json decoder
package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	format := "#%-5d %9.9s %.55s\n"
	now := time.Now()

	withInMonth := make([]*github.Issue, 0)
	withInYear := make([]*github.Issue, 0)
	pastYear := make([]*github.Issue, 0)

	month := now.AddDate(0, -1, 0)
	year := now.AddDate(-1, 0, 0)

	for _, item := range result.Items {
		switch {
		case item.CreatedAt.After(month):
			withInMonth = append(withInMonth, item)
		case item.CreatedAt.After(year) && item.CreatedAt.Before(month):
			withInYear = append(withInYear, item)
		case item.CreatedAt.Before(year):
			pastYear = append(pastYear, item)
		}
	}

	if len(withInMonth) > 0 {
		fmt.Printf("\n1ヶ月未満:\n")
		for _, item := range withInMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(withInMonth) > 0 {
		fmt.Printf("\n１年未満:\n")
		for _, item := range withInMonth {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
	if len(pastYear) > 0 {
		fmt.Printf("\n１年以上:\n")
		for _, item := range pastYear {
			fmt.Printf(format, item.Number, item.User.Login, item.Title)
		}
	}
}
