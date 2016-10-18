# snooze-github-notifications

Snooze github notification for an entire org and re-watch them once you are ready for them. Finally uninterrupted time for some OSS in your vacation!

## Setup

```bash
git clone
glide install
```

## Usage

Generate a github token with `repo` scope https://github.com/settings/tokens/new and export it as `GITHUB_TOKEN` environment variable.

### Stop watching

```bash
snooze-github-notifications stop-org my-org
```

All currently watched repositories will be stored inside a `save-my-org.csv` file inside the current directory.

### Resume watching

```bash
snooze-github-notifications resume-org my-org
```
