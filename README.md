# Golang Boilerplate for OpenAI Service
## Description
* This repository is for the golang-based server boilerplate that wraps OpenAI Service with PostgreSQL and go-chi.
* The codebase uses the [Reverse Engineered ChatGPT API](https://github.com/acheong08/ChatGPT) by [Antonio Cheong](https://github.com/acheong08) and it is also largely influenced by [go-gpt3](https://github.com/PullRequestInc/go-gpt3) from [PullRequest Inc](https://github.com/PullRequestInc).
* It currently lacks the support for GPT-4 yet whereby please feel free to send the pull request on the repository.

## Environment Variables
* Please make sure that the environment variable has been defined in the `./config.yaml` like the following:
```yaml
Environment: DEV
OpenAIEnv:
  API_KEY: "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // if you are going to use the official API
  ACCESS_TOKEN: "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" // if you are going to use the reverse-engineered API
```