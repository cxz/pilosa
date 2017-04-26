package pilosa

import "time"

const (
	// DefaultHost is the default hostname to use.
	DefaultHost = "localhost"

	// DefaultPort is the default port use with the hostname.
	DefaultPort = "10101"

	// DefaultClusterType sets the node intercommunication method.
	DefaultClusterType = "static"

	// DefaultInternalPort the port the nodes intercommunicate on.
	DefaultInternalPort = "14000"
)

// Config represents the configuration for the command.
type Config struct {
	DataDir string `toml:"data-dir"`
	Host    string `toml:"host"`

	Cluster struct {
		ReplicaN        int      `toml:"replicas"`
		Type            string   `toml:"type"`
		Hosts           []string `toml:"hosts"`
		InternalHosts   []string `toml:"internal-hosts"`
		PollingInterval Duration `toml:"polling-interval"`
		InternalPort    string   `toml:"internal-port"`
		GossipSeed      string   `toml:"gossip-seed"`
	} `toml:"cluster"`

	Plugins struct {
		Path string `toml:"path"`
	} `toml:"plugins"`

	AntiEntropy struct {
		Interval Duration `toml:"interval"`
	} `toml:"anti-entropy"`

	LogPath string `toml:"log-path"`
}

// NewConfig returns an instance of Config with default options.
func NewConfig() *Config {
	c := &Config{
		Host: DefaultHost + ":" + DefaultPort,
	}
	c.Cluster.ReplicaN = DefaultReplicaN
	c.Cluster.Type = DefaultClusterType
	c.Cluster.PollingInterval = Duration(DefaultPollingInterval)
	c.Cluster.Hosts = []string{}
	c.Cluster.InternalHosts = []string{}
	c.AntiEntropy.Interval = Duration(DefaultAntiEntropyInterval)
	return c
}

// Duration is a TOML wrapper type for time.Duration.
type Duration time.Duration

// String returns the string representation of the duration.
func (d Duration) String() string { return time.Duration(d).String() }

// UnmarshalText parses a TOML value into a duration value.
func (d *Duration) UnmarshalText(text []byte) error {
	v, err := time.ParseDuration(string(text))
	if err != nil {
		return err
	}

	*d = Duration(v)
	return nil
}

// MarshalText writes duration value in text format.
func (d Duration) MarshalText() (text []byte, err error) {
	return []byte(d.String()), nil
}
