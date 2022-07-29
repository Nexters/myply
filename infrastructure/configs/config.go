package configs

import (
	"os"
	"strconv"
	"time"

	"github.com/google/wire"
	"github.com/joho/godotenv"
)

var Set = wire.NewSet(NewConfig)

// Phase is myply server phase type
type Phase int64

// Config is myply' configuration instance, singleton
type Config struct {
	Phase        Phase
	MongoURI     string
	MongoDBName  string
	MongoTTL     time.Duration
	RedisURI     string
	RedisTimeout time.Duration
}

const (
	// Test phase
	Test Phase = iota + 1
	// Local phase
	Local
	// Production phase
	Production
)

func parsePhase(p string) Phase {
	switch p {
	case "test":
		return Test
	case "local":
		return Local
	case "prod":
		return Production
	}
	return Local
}

// String converts phase to string
func (p Phase) String() string {
	switch p {
	case Test:
		return "test"
	case Local:
		return "local"
	case Production:
		return "production"
	}
	return "local"
}

func NewConfig() (*Config, error) {
	phase := parsePhase(os.Getenv("PHASE"))

	if phase == Test {
		godotenv.Load(".env.test")
	} else if phase == Local {
		godotenv.Load(".env.local")
	}

	mongoTTL, err := strconv.Atoi(os.Getenv("MONGO_TTL"))
	if err != nil {
		return nil, err
	}
	redisTimeout, err := strconv.Atoi(os.Getenv("REDIS_TIMEOUT"))
	if err != nil {
		return nil, err
	}
	return &Config{
		Phase:        phase,
		MongoURI:     os.Getenv("MONGO_URI"),
		MongoDBName:  os.Getenv("MONGO_DB_NAME"),
		MongoTTL:     time.Duration(mongoTTL) * time.Second,
		RedisURI:     os.Getenv("REDIS_URI"),
		RedisTimeout: time.Duration(redisTimeout) * time.Second,
	}, nil
}
