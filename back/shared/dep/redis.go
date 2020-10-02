package dep

var RedisHost = GetEnv("REDIS_HOST", "redis")
var RedisPort = GetEnv("REDIS_PORT", "6379")
var RedisPassword = GetEnv("REDIS_PASSWORD", "")
