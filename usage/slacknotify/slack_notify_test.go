package slacknotify

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"testing"
)

type softFork struct {
	Chain string
	Height int64
	ExpectedParentHash string
	ActualParentHash string
}

func (sf *softFork) string() string {
	return fmt.Sprintf("*soft fork event*\n>*chain name* : %s\n>*height* : %d\n>*expected parent hash* : %s\n>*actual parent hash* : %s\n",
		sf.Chain, sf.Height, sf.ExpectedParentHash, sf.ActualParentHash)
}

type sectionMsg struct {
	Text string `json:"text,omitempty"`
	Type string `json:"type"`
}

type blockMsg struct {
	Msg *sectionMsg `json:"text,omitempty"`
	Type string `json:"type"`
}

type slackMsg struct {
	Blocks []*blockMsg `json:"blocks"`
}

func TestSlackNotify(t *testing.T) {
	incomingWebHooKUrl := "https://hooks.slack.com/services/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"
	fork := softFork{
		Chain:              "ethereum",
		Height:             10000,
		ExpectedParentHash: "0x0000",
		ActualParentHash:   "0x0001",
	}
	msg := slackMsg{
		Blocks: []*blockMsg{
			{
				Type: "section",
				Msg: &sectionMsg{
					Type: "mrkdwn",
					Text: fork.string(),
				},
			},
		},
	}
	msgJson, _ := json.Marshal(msg)
	req, err := http.NewRequest("POST", incomingWebHooKUrl, strings.NewReader(string(msgJson)))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Accepts", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		fmt.Printf("response status code: %d", resp.StatusCode)
		panic(err)
	}
}
