package config

import (
	"fmt"
	env "github.com/caarlos0/env/v6"
)

// Config Available everywhere
var Config *MainConfig

// MainConfig Type exported for use as input type in functions
type MainConfig struct {
	Port			string		`env:"PORT" envDefault:"8000"`
	FileName		string		`env:"FILENAME" envDefault:"epicentre.log"`			
	MaxSize			int		`env:"MAXSIZE" envDefault:"100"`
	MaxAge			int		`env:"MAXAGE" envDefault:"10"`
	MaxBackUp		int		`env:"MAZBACKUP" envDefault:"5"`
	Compress		bool		`env:"COMPRESS" envDefault:"false"`
	Level			string		`env:"LEVEL" envDefault:"info"`
	OutputType		string		`env:"OUTPUTTYPE" envDefault:"json"`
	MongoURI		string		`env:"MONGO_URI" envDefault:"localhost:27017"`
	HerokuAPIToken		string		`env:"HEROKU_API_TOKEN" envDefault:"1111a111-a111-111a-111a-a1aa11aaa111"`
	GithubAPIToken		string		`env:"GITHUB_API_TOKEN" envDefault:"1111a111-a111-111a-111a-a1aa11aaa111"`
	DeployedAppURL		string		`env:"DEPLOYED_URL" envDefault:"https://epicentre.herokuapp.com"`
	KafkaBrokerList		string		`env:"KAFKA_BROKERS" envDefault:"host1:9094,host2:9094,host3:9094"`
	KafkaUsername		string		`env:"KAFKA_USERNAME" envDefault:"username"`
	KafkaPassword		string		`env:"KAFKA_PASSWORD" envDefault:"password"`
	KafkaGroupID		string		`env:"KAFKA_GROUPID" envDefault:"123456"`
	KafkaTopicPrefix	string		`env:"KAFKA_TOPIC_PREFIX" envDefault:"same_as_username"`
}

// LoadConfig Loads the config
func LoadConfig() error  {
	fmt.Println("Initializing configuration variables...")

	cfg := MainConfig{}
	if err := env.Parse(&cfg); err != nil {
		return err
	}

	// Address of parsed configs passing to Global Config variable
	Config = &cfg

	fmt.Println("Initialization Finished")
	
	return nil
}
