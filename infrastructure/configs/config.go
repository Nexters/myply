package configs

import (
	"os"

	"github.com/joho/godotenv"
)

// Phase is myply server phase type
type Phase int64

// Config is myply' configuration instance, singleton
type Config struct {
	Phase       Phase
	MongoURI    string
	MongoDBName string
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

func NewConfig() *Config {
	phase := parsePhase(os.Getenv("PHASE"))

	if phase == Test {
		godotenv.Load(".env.test")
	} else if phase == Local {
		godotenv.Load(".env.local")
	}

	return &Config{
		Phase:       phase,
		MongoURI:    os.Getenv("MONGO_URI"),
		MongoDBName: os.Getenv("MONGO_DB_NAME"),
	}
}
