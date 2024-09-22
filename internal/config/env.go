package config

import (
	"fmt"
	"net"
	"time"

	"github.com/caarlos0/env/v10"
)

// Provider предоставляет интерфейс для получения конфигурации.
type Provider interface {
	Config() *Config
}

// Config общий конфиг.
type Config struct {
	DebugMode   bool   `env:"DEBUG_MODE" envDefault:"false"`
	Environment string `env:"ENV" envDefault:"local"`
	Bucket      Bucket
	Postgres    Postgres
	Redis       Redis
	GRPC        GRPC
	Log         Log
}

// Config возвращаем сам конфиг.
func (c Config) Config() *Config {
	return &c
}

// Postgres конфиг подключения к БД.
type Postgres struct {
	Host               string `env:"POSTGRES_HOST" envDefault:"localhost"`
	Port               string `env:"POSTGRES_PORT" envDefault:"5432"`
	User               string `env:"POSTGRES_USER" envDefault:"root"`
	Password           string `env:"POSTGRES_PASSWORD" envDefault:"password"`
	DB                 string `env:"POSTGRES_DB" envDefault:"postgres"`
	SslMode            string `env:"POSTGRES_SSL_MODE" envDefault:"disable"`
	DSN                string `env:"POSTGRES_DSN"`
	MaxOpenConnections int    `env:"POSTGRES_MAX_OPEN_CONNS" envDefault:"100"`
}

// GRPC конфиг подключения к grpc.
type GRPC struct {
	Host     string `env:"GRPC_HOST"`
	Port     string `env:"GRPC_PORT"`
	Protocol string `env:"GRPC_PROTOCOL"`
	Address  string
}

// Bucket конфиг.
type Bucket struct {
	LoginLimit    int `env:"BUCKET_LOGIN_LIMIT" envDefault:"10"`
	PasswordLimit int `env:"BUCKET_PASSWORD_LIMIT" envDefault:"100"`
	IPLimit       int `env:"BUCKET_IP_LIMIT" envDefault:"1000"`
	WindowSize    time.Duration
	TTL           time.Duration
	WindowSizeMin int `env:"BUCKET_WINDOW_SIZE" envDefault:"1"`
	TTLMin        int `env:"BUCKET_TTL" envDefault:"10"`
}

// Redis конфиг подключения к redis.
type Redis struct {
	Address  string
	Host     string `env:"REDIS_HOST" envDefault:"redis"`
	Port     string `env:"REDIS_PORT" envDefault:"6379"`
	Password string `env:"REDIS_PASSWORD" envDefault:""`
	DB       int    `env:"REDIS_DB" envDefault:"0"`
}

// Log конфиг для логов.
type Log struct {
	FileName   string `env:"LOG_FILENAME" envDefault:"logs/app.log"`
	Level      string `env:"LOG_LEVEL" envDefault:"info"`
	MaxSize    int    `env:"LOG_MAXSIZE" envDefault:"5"`
	MaxBackups int    `env:"LOG_MAXBACKUPS" envDefault:"3"`
	MaxAge     int    `env:"LOG_MAXAGE" envDefault:"10"`
	Compress   bool   `env:"LOG_COMPRESS" envDefault:"false"`
	StdOut     bool   `env:"LOG_STDOUT" envDefault:"false"`
}

// New создаем новый конфиг.
func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("loading config from env is failed: %w", err)
	}
	buildDSN(&cfg.Postgres)
	cfg.GRPC.Address = net.JoinHostPort(cfg.GRPC.Host, cfg.GRPC.Port)
	cfg.Redis.Address = net.JoinHostPort(cfg.Redis.Host, cfg.Redis.Port)
	cfg.Bucket.WindowSize = time.Duration(cfg.Bucket.WindowSizeMin) * time.Minute
	cfg.Bucket.TTL = time.Duration(cfg.Bucket.TTLMin) * time.Minute

	return cfg, nil
}

func buildDSN(p *Postgres) {
	p.DSN = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.User, p.Password, p.Host, p.Port, p.DB, p.SslMode)
}
