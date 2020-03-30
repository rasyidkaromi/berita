package service

import "log"
import "gopkg.in/dealancer/validate.v2"
import "github.com/pkg/errors"
import "github.com/rasyidkaromi/berita/models"
import "github.com/rasyidkaromi/berita/constans"
import "github.com/rasyidkaromi/berita/database/mysql"

type MysqlService interface {}

type mysqlservice struct {}

func NewMysqlService() MysqlService {}
