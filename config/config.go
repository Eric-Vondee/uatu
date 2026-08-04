package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/Eric-Vondee/metron"
	"github.com/spf13/viper"
)

type PostgresConfig struct {
	DSN string `mapstructure:"POSTGRES_DSN"`
}
type Config struct {
	AllowedOrigins []string `mapstructure:"ALLOWED_ORIGINS" validate:"required"`
	PORT           string   `mapstructure:"PORT" validate:"required"`
	Otel           struct {
		IsEnabled bool   `mapstructure:"OTEL_ENABLED"`
		Endpoint  string `mapstructure:"OTEL_ENDPOINT" validate:"required"`
		UseTLS    bool   `mapstructure:"OTEL_USE_TLS" validate:"required"`
		Headers   string `mapstructure:"OTEL_HEADERS" validate:"required"`
	} `mapstructure:",squash"`
	Database struct {
		Postgres PostgresConfig `mapstructure:",squash" validate:"required"`
	} `mapstructure:",squash"`
	RPC struct {
		ARB       string `mapstructure:"ARBITRUM_RPC_URL" validate:"required"`
		AVAX      string `mapstructure:"AVAX_RPC_URL" validate:"required"`
		BASE      string `mapstructure:"BASE_RPC_URL" validate:"required"`
		BSC       string `mapstructure:"BSC_RPC_URL" validate:"required"`
		CELO      string `mapstructure:"CELO_RPC_URL" validate:"required"`
		ETH       string `mapstructure:"ETHEREUM_RPC_URL" validate:"required"`
		GNOSIS    string `mapstructure:"GNOSIS_RPC_URL" validate:"required"`
		INK       string `mapstructure:"INK_RPC_URL" validate:"required"`
		LINEA     string `mapstructure:"LINEA_RPC_URL" validate:"required"`
		LISK      string `mapstructure:"LISK_RPC_URL" validate:"required"`
		MANTLE    string `mapstructure:"MANTLE_RPC_URL" validate:"required"`
		MONAD     string `mapstructure:"MONAD_RPC_URL" validate:"required"`
		OP        string `mapstructure:"OPTIMISM_RPC_URL" validate:"required"`
		PLASMA    string `mapstructure:"PLASMA_RPC_URL" validate:"required"`
		POLYGON   string `mapstructure:"POLYGON_RPC_URL" validate:"required"`
		ROBINHOOD string `mapstructure:"ROBINHOOD_RPC_URL" validate:"required"`
		WORLD     string `mapstructure:"WORLD_RPC_URL" validate:"required"`
		UNICHAIN  string `mapstructure:"UNICHAIN_RPC_URL" validate:"required"`
	} `mapstructure:",squash"`
}

func InitializeConfig() (*Config, error) {
	var cfg Config

	viper.SetConfigType("env")

	for _, path := range []string{".env", "../.env"} {
		if _, err := os.Stat(path); err != nil {
			continue
		}
		viper.SetConfigFile(path)
		if err := viper.ReadInConfig(); err != nil {
			return nil, fmt.Errorf("error reading config file %q: %w", path, err)
		}
		break
	}

	viper.AutomaticEnv()
	bindEnvs(reflect.TypeOf(cfg))

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}
	return &cfg, nil
}

func bindEnvs(t reflect.Type) {
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Type.Kind() == reflect.Struct {
			bindEnvs(field.Type)
			continue
		}
		name := strings.Split(field.Tag.Get("mapstructure"), ",")[0]
		if name == "" || name == "-" {
			continue
		}
		_ = viper.BindEnv(name)
	}
}

func (c *Config) Validate() error {
	return metron.ValidateStruct(c)
}

func (c *Config) GetRPC(slug string) string {
	switch slug {
	case "arbitrum":
		return c.RPC.ARB
	case "avalanche":
		return c.RPC.AVAX
	case "base":
		return c.RPC.BASE
	case "bsc":
		return c.RPC.BSC
	case "celo":
		return c.RPC.CELO
	case "ethereum":
		return c.RPC.ETH
	case "gnosis":
		return c.RPC.GNOSIS
	case "ink":
		return c.RPC.INK
	case "linea":
		return c.RPC.LINEA
	case "lisk":
		return c.RPC.LISK
	case "mantle":
		return c.RPC.MANTLE
	case "monad":
		return c.RPC.MONAD
	case "optimism":
		return c.RPC.OP
	case "plasma":
		return c.RPC.PLASMA
	case "polygon":
		return c.RPC.POLYGON
	case "robinhood":
		return c.RPC.ROBINHOOD
	case "worldchain":
		return c.RPC.WORLD
	case "unichain":
		return c.RPC.UNICHAIN
	default:
		return ""
	}
}
