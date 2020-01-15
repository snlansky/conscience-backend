package meeting_chain

import (
	"conscience-backend/bao"
	"conscience-backend/model"
	"encoding/json"
	"github.com/pkg/errors"
)

type FaceIDService struct {
	contract bao.IContract
}

func NewFaceIDService(contract bao.IContract) *FaceIDService {
	return &FaceIDService{contract: contract}
}

func (s *FaceIDService) RegisterFaceID(accessToken string, id *model.FaceID) (string, error) {
	var args []interface{}
	args = append(args, id)
	txId, _, err := s.contract.Invoke(accessToken,
		"FaceIDService.RegisterFaceID",
		args,
	)
	return txId, err
}

func (s *FaceIDService) RegisterCertificate(accessToken string, id *model.FaceID) (string, error) {
	var args []interface{}
	args = append(args, id)
	txId, _, err := s.contract.Invoke(accessToken,
		"FaceIDService.RegisterCertificate",
		args,
	)
	return txId, err
}

func (s *FaceIDService) Record(accessToken string, id *model.FaceID) (string, error) {
	var args []interface{}
	args = append(args, id)
	txId, _, err := s.contract.Invoke(accessToken,
		"FaceIDService.Record",
		args,
	)
	return txId, err
}

func (s *FaceIDService) GetUser(accessToken string) (*model.RegisterUser, error) {
	var args []interface{}
	res, err := s.contract.Query(accessToken,
		"FaceIDService.GetUser",
		args,
	)
	if err != nil {
		return nil, errors.WithMessage(err, "contract query error")
	}

	var u model.RegisterUser
	err = json.Unmarshal([]byte(res), &u)
	if err != nil {
		return nil, errors.WithMessage(err, "json unmarshal error")
	}
	return &u, err
}

func (s *FaceIDService) HistoryFaceIDs(accessToken string, id *model.HistoryFaceIDs) ([]*model.FaceID, error) {
	var args []interface{}
	args = append(args, id)
	res, err := s.contract.Query(accessToken,
		"FaceIDService.HistoryFaceIDs",
		args,
	)
	if err != nil {
		return nil, errors.WithMessage(err, "contract query error")
	}
	var list []*model.FaceID
	err = json.Unmarshal([]byte(res), &list)
	if err != nil {
		return nil, errors.WithMessage(err, "json unmarshal error")
	}

	return list, err
}
