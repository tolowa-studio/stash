package config

import (
	"os"
	"time"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	// Store (PostgreSQL only)
	StoreDSN      string `env:"STASH_POSTGRES_DSN,required"`
	VectorDim     int    `env:"STASH_VECTOR_DIM,required"`
	MaxResultSize int    `env:"STASH_MAX_RESULT_SIZE,required"`

	// OpenAI (embeddings + reasoning)
	OpenAIAPIKey                string        `env:"STASH_OPENAI_API_KEY,required"`
	OpenAIBaseURL               string        `env:"STASH_OPENAI_BASE_URL,required"`
	EmbeddingModel              string        `env:"STASH_EMBEDDING_MODEL,required"`
	EmbedderMaxRetries          int           `env:"STASH_EMBEDDER_MAX_RETRIES" envDefault:"1"`
	EmbedderRateLimitCooldown   time.Duration `env:"STASH_EMBEDDER_RATE_LIMIT_COOLDOWN" envDefault:"5m"`
	EmbedderPaymentCooldown     time.Duration `env:"STASH_EMBEDDER_PAYMENT_COOLDOWN" envDefault:"1h"`
	EmbedderServerErrorCooldown time.Duration `env:"STASH_EMBEDDER_SERVER_ERROR_COOLDOWN" envDefault:"1m"`
	ReasonerModel               string        `env:"STASH_REASONER_MODEL,required"`
	ReasonerAPIKey              string        `env:"STASH_REASONER_API_KEY"`
	ReasonerBaseURL             string        `env:"STASH_REASONER_BASE_URL"`
	ReasonerMaxTokens           int64         `env:"STASH_REASONER_MAX_TOKENS" envDefault:"4096"`
	ReasonerMaxRetries          int           `env:"STASH_REASONER_MAX_RETRIES" envDefault:"2"`
	ReasonerRateLimitCooldown   time.Duration `env:"STASH_REASONER_RATE_LIMIT_COOLDOWN" envDefault:"5m"`
	ReasonerPaymentCooldown     time.Duration `env:"STASH_REASONER_PAYMENT_COOLDOWN" envDefault:"1h"`
	ReasonerServerErrorCooldown time.Duration `env:"STASH_REASONER_SERVER_ERROR_COOLDOWN" envDefault:"1m"`

	// Memory
	ContextTTL time.Duration `env:"STASH_CONTEXT_TTL,required"`

	// Server
	HTTPAddr  string `env:"STASH_HTTP_ADDR,required"`
	LogLevel  string `env:"STASH_LOG_LEVEL,required"`
	LogFormat string `env:"STASH_LOG_FORMAT,required"`

	// Consolidation
	ConsolidationBatchSize           int     `env:"STASH_CONSOLIDATION_BATCH_SIZE" envDefault:"100"`
	ConsolidationSimilarityThreshold float64 `env:"STASH_CONSOLIDATION_SIMILARITY_THRESHOLD" envDefault:"0.85"`
	ConsolidationDedupThreshold      float64 `env:"STASH_CONSOLIDATION_DEDUP_THRESHOLD" envDefault:"0.85"`
	ConsolidationWindow              string  `env:"STASH_CONSOLIDATION_WINDOW" envDefault:"168h"`
	DecayFactor                      float64 `env:"STASH_DECAY_FACTOR" envDefault:"0.95"`
	ExpiryThreshold                  float32 `env:"STASH_EXPIRY_THRESHOLD" envDefault:"0.1"`
	HypothesisAutoConfirmThreshold   float32 `env:"STASH_HYPOTHESIS_AUTO_CONFIRM_THRESHOLD" envDefault:"0.9"`
	HypothesisAutoRejectThreshold    float32 `env:"STASH_HYPOTHESIS_AUTO_REJECT_THRESHOLD" envDefault:"0.9"`

	// Retrieval learning is deliberately independent from fact confidence.
	RetrievalLearningEnabled bool          `env:"STASH_RETRIEVAL_LEARNING_ENABLED" envDefault:"false"`
	RetrievalOverfetchFactor int           `env:"STASH_RETRIEVAL_OVERFETCH_FACTOR" envDefault:"3"`
	RetrievalUtilityWeight   float64       `env:"STASH_RETRIEVAL_UTILITY_WEIGHT" envDefault:"0.08"`
	RetrievalMaxUtilityDelta float64       `env:"STASH_RETRIEVAL_MAX_UTILITY_DELTA" envDefault:"0.10"`
	RecallHistoryRetention   time.Duration `env:"STASH_RECALL_HISTORY_RETENTION" envDefault:"2160h"`
	EmbeddingBackfillBatch   int           `env:"STASH_EMBEDDING_BACKFILL_BATCH" envDefault:"25"`
}

func NewFromFile(filename string) (*Config, error) {
	if _, err := os.Stat(filename); err == nil {
		if err := godotenv.Load(filename); err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, err
	}

	cfg := &Config{}
	opts := env.Options{
		RequiredIfNoDef: true,
	}
	if err := env.ParseWithOptions(cfg, opts); err != nil {
		return nil, err
	}
	return cfg, nil
}
