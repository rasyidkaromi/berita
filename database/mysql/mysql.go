package mysql

import "context"
import "fmt"
import "time"
import "strings"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "github.com/jmoiron/sqlx"
import "github.com/pkg/errors"
import "github.com/rasyidkaromi/berita/models"
import "github.com/rasyidkaromi/berita/constans"

type MySQLDatabase interface {
	Get(filter map[string]interface{}) (*models.News, error)
	Store(data *models.News) error
	Update(data map[string]interface{}, id int) (*models.News, error)
	Delete(id int) error
}

type mySQLDatabase struct {
	url     string
	timeout time.Duration
}

func NewMysqlDatabase(URL string, timeout int) (MySQLDatabase, error) {}

