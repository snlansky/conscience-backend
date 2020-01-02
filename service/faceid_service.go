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

func (s *FaceIDService) RegisterFaceID(req *model.RequestRegisterFaceID) error {
	id := &model.FaceID{
		ID:         req.ID,
		SourceType: req.SourceType,
		SourceHash: req.SourceHash,
		Algorithm:  req.Algorithm,
		Labels:     req.Labels,
		Metadata:   req.Metadata,
		Timestamp:  req.Timestamp,
	}
	return meeting_chain.DefaultFaceIDService.RegisterFaceID(req.AccessToken, id)
}
