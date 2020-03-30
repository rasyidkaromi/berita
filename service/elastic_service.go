package service

import "log"
import	"github.com/pkg/errors"
import  "gopkg.in/dealancer/validate.v2"
import	"github.com/rasyidkaromi/berita/models"
import  "github.com/rasyidkaromi/berita/constans"
import  "github.com/rasyidkaromi/berita/database/elastic"

type ElasticService interface {}

type elasticService struct {}

func NewElasticService() ElasticService {}
