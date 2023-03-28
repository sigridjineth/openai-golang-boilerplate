package config

const GlobalConfigKey = "globalconfig"
const DefaultConfigPath = "./config.yaml"
const TestConfigPath = "../config.yaml"

type GlobalConfig struct {
	Environment Environment `koanf:"Environment" envDefault:"dev"`
	OpenAIEnv   OpenAIENV
}

func (g *GlobalConfig) IsDev() bool {
	return g.Environment == DevEnvironment
}

func (g *GlobalConfig) IsStaging() bool {
	return g.Environment == StagingEnvironment
}

func (g *GlobalConfig) IsProd() bool {
	return g.Environment == ProdEnvironment
}
