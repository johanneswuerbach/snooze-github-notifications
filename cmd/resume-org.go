package cmd

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/spf13/cobra"
)

func init() {
	var token string

	var cmd = &cobra.Command{
		Use:   "resume-org",
		Short: "Resume notifications for a github organization",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = os.Getenv("GITHUB_TOKEN")
			}

			if token == "" {
				log.Fatalf("No GitHub token found.")
			}

			if token == "" {
				token = os.Getenv("GITHUB_TOKEN")
			}

			if token == "" {
				log.Fatalf("No GitHub token found.")
			}

			log.Printf("Args %v", args)

			var org = args[0]
			log.Printf("Re-watch repos in org %s", org)

			client := createGithubClient(token)

			file, err := os.Open(csvName(org))
			if err != nil {
				log.Fatal(err)
			}

			defer file.Close()

			reader := csv.NewReader(file)

			repos, err := reader.ReadAll()

			repoCount := len(repos)
			log.Printf("Re-watching repositories %v", repoCount)

			subscribed := true

			for i, line := range repos {
				owner := line[1]
				name := line[2]
				_, _, err := client.Activity.SetRepositorySubscription(owner, name, &github.Subscription{
					Subscribed: &subscribed,
				})

				if err != nil {
					log.Printf("(%d/%d) Error: %s, %v", i+1, repoCount, fullName(owner, name), err)
				} else {
					log.Printf("(%d/%d) %s", i+1, repoCount, fullName(owner, name))
				}
			}
		},
	}

	RootCmd.AddCommand(cmd)
	cmd.Flags().StringVar(&token, "token", "", "GitHub token")
}
