package service

import 	"log"
import	"github.com/pkg/errors"
import	"github.com/rasyidkaromi/berita/models"
import  "github.com/rasyidkaromi/berita/database/redis"

type CacheDatabase interface {}

type RedisService interface {}

type redisService struct {}

func NewRedisService() RedisService {}
