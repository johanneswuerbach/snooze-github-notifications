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
		Use:   "stop-org",
		Short: "Stop notifications for a github organization",
		Run: func(cmd *cobra.Command, args []string) {
			if token == "" {
				token = os.Getenv("GITHUB_TOKEN")
			}

			if token == "" {
				log.Fatalf("No GitHub token found.")
			}

			log.Printf("Args %v", args)

			var org = args[0]
			log.Printf("Stop watching repos in org %s", org)

			client := createGithubClient(token)

			opt := &github.ListOptions{
				PerPage: 10,
			}
			// get all pages of results
			var filteredRepos []*github.Repository
			for {
				repos, resp, err := client.Activity.ListWatched("", opt)
				if err != nil {
					log.Fatal(err)
				}

				for _, repo := range repos {
					if *repo.Owner.Login == org {
						filteredRepos = append(filteredRepos, repo)
					}
				}

				if resp.NextPage == 0 {
					break
				}
				opt.Page = resp.NextPage
			}

			repoCount := len(filteredRepos)
			log.Printf("Snoozing %d repositories", repoCount)

			var data = [][]string{}

			for _, repo := range filteredRepos {
				data = append(data, []string{"repo", *repo.Owner.Login, *repo.Name})
			}

			file, err := os.Create(csvName(org))
			checkError("Cannot create file", err)
			defer file.Close()

			writer := csv.NewWriter(file)

			err = writer.WriteAll(data)
			checkError("Cannot write to file", err)
			file.Close()

			for i, repo := range filteredRepos {
				owner := *repo.Owner.Login
				name := *repo.Name
				_, err := client.Activity.DeleteRepositorySubscription(owner, name)

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

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
