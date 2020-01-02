package meeting_chain

import (
	"conscience-backend/bao"
	"conscience-backend/model"
	"fmt"
)

type FaceIDService struct {
	contract bao.IContract
}

func NewFaceIDService(contract bao.IContract) *FaceIDService {
	return &FaceIDService{contract: contract}
}

func (s *FaceIDService) RegisterFaceID(accessToken string, id *model.FaceID) error {
	var args []interface{}
	args = append(args, id)
	txId, _, err := s.contract.Invoke(accessToken,
		"FaceIDService.RegisterFaceID",
		args,
	)
	fmt.Println(txId)
	return err
}
