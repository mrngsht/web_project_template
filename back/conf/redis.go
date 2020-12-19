package conf

var RedisHost = getEnv("REDIS_HOST", "redis")
var RedisPort = getEnv("REDIS_PORT", "6379")
var RedisPassword = getEnv("REDIS_PASSWORD", "")
