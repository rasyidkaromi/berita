package redis

import	"encoding/json"
import	"time"
import	"github.com/go-redis/redis"
import	"github.com/pkg/errors"
import	"fmt"
import	"strconv"
import "github.com/rasyidkaromi/berita/models"
import "github.com/rasyidkaromi/berita/constans"

type CacheDatabase interface {
	GetBy(param models.GetPayload) ([]models.News, error)
	Store(data []models.News) error
	Update(data models.News) error
	Delete(data models.News) error
}

type newsRedisRepository struct {
	client     *redis.Client
	expiration time.Duration
}

func client(redisURL string) (*redis.Client, error) {}

func NewRedisDatabase(redisURL string, expiration int) (CacheDatabase, error) {}
