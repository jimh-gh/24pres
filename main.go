package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	// Define the URL to fetch data from
	url := "https://interactives.apelections.org/election-results/data-live/2024-11-05/results/national/P/progress.json"

	// Make the HTTP request to fetch the data
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Check if the HTTP request was successful
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Error: received non-200 response code: %d\n", resp.StatusCode)
		return
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return
	}

	// Parse the JSON data
	var results ElectionResults
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		return
	}

	//fmt.Printf("Parsed results: %+v\n", results)
	mesg := fmt.Sprintf("Electoral Votes:\n%s %d\n%s %d\nIn %d / %d\nPrecinctPct %.2f", results.Who0(), results.US0.Candidates[0].ElectWon,
		results.Who1(), results.US0.Candidates[1].ElectWon, results.US0.Candidates[0].ElectWon+results.US0.Candidates[1].ElectWon, results.US0.ElectTotal, results.US0.PrecinctsReportingPct)
	push("24 Race", mesg)
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
