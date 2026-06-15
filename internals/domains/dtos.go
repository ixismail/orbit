package domains

// OrbitConfig represents the entire config.yaml file
type OrbitConfig struct {
	Global GlobalSettings      `mapstructure:"global"`
	Routes map[string]RouteDef `mapstructure:"routes"`
}

type GlobalSettings struct {
	CorsBypass bool `mapstructure:"cors_bypass"`
	AutoTLS    bool `mapstructure:"auto_tls"`
}

type RouteDef struct {
	Target     string       `mapstructure:"target"`

	// Pointer so we know if it was explicitly set because
	// Go sets the default value for bool to false if not explicitly set
	CorsBypass *bool        `mapstructure:"cors_bypass"`
	Chaos      *ChaosConfig `mapstructure:"chaos"`
}

type ChaosConfig struct {
	LatencyMS        int `mapstructure:"latency_ms"`
	ErrorRatePercent int `mapstructure:"error_rate_percent"`
}