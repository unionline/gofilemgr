package initializers

import (
	env "gofilemgr/internal/env/redis"
	"gofilemgr/internal/initializers/config"
	"gofilemgr/internal/initializers/db"
	"gofilemgr/internal/initializers/model"
)

// Initialize ...
func Initialize() {
	config.Init()
	db.Init()
	model.Init()
	env.RedisInit()
}
