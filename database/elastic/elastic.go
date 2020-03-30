package elastic

import	"context"
import	"encoding/json"
import	"fmt"
import	"strconv"
import	"time"
import	elt "github.com/olivere/elastic"
import	"github.com/pkg/errors"
import	"github.com/rasyidkaromi/berita/models"


type ElasticDatabase interface {}

type elasticDatabase struct {}

func NewElasticDatabase(URL, index string, timeout int) (ElasticDatabase, error) {}
