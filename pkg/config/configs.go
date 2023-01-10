package configs

type (
	App struct {
		Name    string `yaml:"name"`
		Version string `yaml:"version"`
	}

	HTTP struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}

	MariaDB struct {
		URL      string `yaml:"url"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}

	RabbitMQ struct {
		URL      string `yaml:"url"`
		Consumer struct {
			Exchange struct {
				Name       string `yaml:"name"`
				Type       string `yaml:"type"`
				Durable    bool   `yaml:"durable"`
				AutoDelete bool   `yaml:"auto_delete"`
				Internal   bool   `yaml:"internal"`
				NoWait     bool   `yaml:"no_wait"`
			} `yaml:"exchange"`
			Queue struct {
				Name       string `yaml:"name"`
				Durable    bool   `yaml:"durable"`
				AutoDelete bool   `yaml:"auto_delete"`
				Exclusive  bool   `yaml:"exclusive"`
				NoWait     bool   `yaml:"no_wait"`
			} `yaml:"queue"`
			BindingOptions struct {
				RoutingKey string `yaml:"routing_key"`
				NoWait     bool   `yaml:"no_wait"`
			} `yaml:"binding_options"`
		}
	}
)
