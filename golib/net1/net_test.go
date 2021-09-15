package net1

import (
	"fmt"
	"net/url"
	"testing"
)

func TestURL(t *testing.T) {
	slackURL := ""
	u, err := url.ParseRequestURI(slackURL)
	if err != nil {
		fmt.Println("slack url is invalid, slack notify will not work")
	}
	if u == nil {
		fmt.Println("slack url is invalid, slack notify will not work")
	} else {
		fmt.Println("slack url is valid")
	}
}
