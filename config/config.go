package config

import (
	"fmt"
	env "github.com/caarlos0/env/v6"
)

// Config Available everywhere
var Config *MainConfig

// MainConfig Type exported for use as input type in functions
type MainConfig struct {
	Port       		string			`env:"PORT" envDefault:"8000"`
	FileName   		string			`env:"FILENAME" envDefault:"epicentre.log"`			
	MaxSize    		int				`env:"MAXSIZE" envDefault:"100"`
	MaxAge     		int				`env:"MAXAGE" envDefault:"10"`
	MaxBackUp  		int				`env:"MAZBACKUP" envDefault:"5"`
	Compress   		bool			`env:"COMPRESS" envDefault:"false"`
	Level      		string			`env:"LEVEL" envDefault:"info"`
	OutputType 		string			`env:"OUTPUTTYPE" envDefault:"json"`
	HerokuAPIToken	string			`env:"HEROKU_API_TOKEN" envDefault:"1111a111-a111-111a-111a-a1aa11aaa111"`
	GithubAPIToken	string			`env:"GITHUB_API_TOKEN" envDefault:"1111a111-a111-111a-111a-a1aa11aaa111"`
	DeployedAppURL	string			`env:"DEPLOYEDURL" envDefault:"https://epicentre.herokuapp.com"`
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