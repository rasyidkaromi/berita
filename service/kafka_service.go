package service

import	"encoding/json"
import	"fmt"
import	"log"
import	"time"
import	"github.com/rasyidkaromi/berita/database/kafka"
import	"github.com/rasyidkaromi/berita/models"
import	"github.com/pkg/errors"

type MySQLDatabase interface {}
type KafkaService interface {}

type kafkaService struct {}

func NewKafkaService() KafkaService {}
