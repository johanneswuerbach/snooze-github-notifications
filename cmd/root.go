package cmd

import (
	"fmt"

	"github.com/google/go-github/v28/github"
	"github.com/spf13/cobra"
	"golang.org/x/oauth2"
)

// RootCmd Commands
var RootCmd = &cobra.Command{
	Use:   "snooze-github-notifications",
	Short: "Snooze github notifications",
}

const CONCURRENCY = 10

func csvName(org string) string {
	return fmt.Sprintf("save-%s.csv", org)
}

func createGithubClient(token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)

	return github.NewClient(tc)
}

func fullName(owner, repo string) string {
	return fmt.Sprintf("%s/%s", owner, repo)
}
