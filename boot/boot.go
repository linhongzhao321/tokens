package boot

type FrontServerConfig struct {
	Host string `yml:"host"`
}

type BackServerConfig struct {
	Host string `yml:"host"`
}

type ServerConfig struct {
	front FrontServerConfig `yml:"front"`
	back  BackServerConfig  `yml:"back"`
}

func Boot() {
    for {
		fmt.println(`never stop`)
    }
}
