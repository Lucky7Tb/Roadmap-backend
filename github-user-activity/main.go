package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

type GithubActivityResponse struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	Repo struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`
	Payload struct {
		Action  string        `json:"action,omitempty"`
		Commits []interface{} `json:"commits,omitempty"`
		RefType string        `json:"ref_type,omitempty"`
		Forkee  struct {
			FullName string `json:"full_name,omitempty"`
		} `json:"forkee,omitempty"`
	} `json:"payload,omitempty"`
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide github username")
		os.Exit(1)
	} else {
		request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://api.github.com/users/%s/events", os.Args[1]), nil)
		if err != nil {
			fmt.Printf("Failed to create a request to github. Error: %s\n", err.Error())
			os.Exit(1)
		}

		response, err := http.DefaultClient.Do(request)
		if err != nil {
			fmt.Printf("Failed to make a request to github. Error: %s\n", err.Error())
			os.Exit(1)
		}
		if response.StatusCode == 403 {
			fmt.Println("Cannot see this username github activity")
			os.Exit(1)
		}

		if response.StatusCode == 503 {
			fmt.Println("Github api not available")
			os.Exit(1)
		}
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Printf("Failed to read response body. Error: %s\n", err.Error())
			os.Exit(1)
		}

		var githubActivities []GithubActivityResponse
		if err := json.Unmarshal(responseBody, &githubActivities); err != nil {
			fmt.Printf("Failed to convert json to struct. Error: %s\n", err.Error())
			os.Exit(1)
		}

		for _, activity := range githubActivities {
			switch activity.Type {
			case "CommitCommentEvent":
				fmt.Printf("- Comment to a commit in %s\n", activity.Repo.Name)
			case "CreateEvent":
				fmt.Printf("- Created %s in %s\n", activity.Payload.RefType, activity.Repo.Name)
			case "DeleteEvent":
				fmt.Printf("- Deleted %s in %s\n", activity.Payload.RefType, activity.Repo.Name)
			case "ForkEvent":
				fmt.Printf("- Fork %s to %s\n", activity.Repo.Name, activity.Payload.Forkee.FullName)
			case "IssueCommentEvent":
				fmt.Printf("- %s issue in %s\n", strings.ToLower(activity.Payload.Action), activity.Repo.Name)
			case "IssuesEvent":
				fmt.Printf("- %s issue in %s\n", strings.ToLower(activity.Payload.Action), activity.Repo.Name)
			case "PushEvent":
				fmt.Printf("- Pushed %d commits to %s\n", len(activity.Payload.Commits), activity.Repo.Name)
			case "WatchEvent":
				fmt.Printf("- Starred %s\n", activity.Repo.Name)
			case "PullRequestEvent":
				fmt.Printf("- PR %s in %s\n", strings.ToLower(activity.Payload.Action), activity.Repo.Name)
			case "PublicEvent":
				fmt.Printf("- Made %s public\n", activity.Repo.Name)
			default:
				fmt.Printf("- Do %s to %s\n", activity.Type, activity.Repo.Name)
			}
		}
	}
}
