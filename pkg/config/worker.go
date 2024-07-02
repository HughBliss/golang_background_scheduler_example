package config

type WorkerConfig struct {
	Redis *RedisConfig `yaml:"redis"`

	CriticalQueuePriority int `yaml:"critical_queue_priority" env:"WORKER_SERVER_CRITICAL_QUEUE_PRIORITY" env-default:"6"`
	DefaultQueuePriority  int `yaml:"default_queue_priority" env:"WORKER_SERVER_DEFAULT_QUEUE_PRIORITY" env-default:"3"`
	LowQueuePriority      int `yaml:"low_queue_priority" env:"WORKER_SERVER_LOW_QUEUE_PRIORITY" env-default:"1"`
}
