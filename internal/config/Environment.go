package config

type Environment string

const (
	DevEnvironment     Environment = "dev"
	StagingEnvironment Environment = "staging"
	ProdEnvironment    Environment = "prod"
)
