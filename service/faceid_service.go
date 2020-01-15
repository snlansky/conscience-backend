package service

import (
	"conscience-backend/bao/meeting_chain"
	"conscience-backend/model"
)

type FaceIDService struct {
}

func NewFaceIDService() *FaceIDService {
	return &FaceIDService{}
}

func (s *FaceIDService) RegisterFaceID(req *model.RequestRegisterFaceID) (string, error) {
	return meeting_chain.DefaultFaceIDService.RegisterFaceID(req.AccessToken, &req.FaceID)
}

func (s *FaceIDService) RegisterCertificate(req *model.RequestRegisterCertificate) (string, error) {
	return meeting_chain.DefaultFaceIDService.RegisterCertificate(req.AccessToken, &req.FaceID)
}

func (s *FaceIDService) Record(req *model.RequestRecord) (string, error) {
	return meeting_chain.DefaultFaceIDService.Record(req.AccessToken, &req.FaceID)
}

func (s *FaceIDService) GetUser(req *model.RequestGetUser) (*model.RegisterUser, error) {
	return meeting_chain.DefaultFaceIDService.GetUser(req.AccessToken)
}

func (s *FaceIDService) HistoryFaceIDs(req *model.RequestHistoryFaceIDs) ([]*model.FaceID, error) {
	return meeting_chain.DefaultFaceIDService.HistoryFaceIDs(req.AccessToken, &req.RequestFaceIDHistory)
}
