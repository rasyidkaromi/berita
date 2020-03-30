package controller

import	"github.com/rasyidkaromi/berita/models"
import	"encoding/json"
import	"github.com/pkg/errors"

type JsonController interface {
	Decode(input []byte) (*models.BeritaStruct, error)
	Encode(input *models.BeritaStruct) ([]byte, error)
	DecodeMap(input []byte) (map[string]interface{}, error)
	EncodeMap(input map[string]interface{}) ([]byte, error)
	EncodeGetData(input []models.BeritaStruct) ([]byte, error)
}

type jsonController struct{}

func NewJSON() JsonController {
	return &jsonController{}
}

func (u *jsonController) Decode(input []byte) (*models.BeritaStruct, error) {
	news := new(models.BeritaStruct)
	if e := json.Unmarshal(input, news); e != nil {
		return nil, errors.Wrap(e, "serializer.Logic.Decode")
	}
	return news, nil
}

func (u *jsonController) Encode(input *models.BeritaStruct) ([]byte, error) {
	rawMsg, e := json.Marshal(input)
	if e != nil {
		return nil, errors.Wrap(e, "serializer.Logic.Encode")
	}
	return rawMsg, nil
}

func (u *jsonController) DecodeMap(input []byte) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	if e := json.Unmarshal(input, &res); e != nil {
		return res, errors.Wrap(e, "serializer.Logic.DecodeMap")
	}
	return res, nil
}

func (u *jsonController) EncodeMap(input map[string]interface{}) ([]byte, error) {
	rawMsg, e := json.Marshal(input)
	if e != nil {
		return nil, errors.Wrap(e, "serializer.Logic.EncodeMap")
	}
	return rawMsg, nil
}

func (u *jsonController) EncodeGetData(input []models.BeritaStruct) ([]byte, error) {
	rawMsg, e := json.Marshal(input)
	if e != nil {
		return nil, errors.Wrap(e, "serializer.Logic.EncodeGetData")
	}
	return rawMsg, nil
}
