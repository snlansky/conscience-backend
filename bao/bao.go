package bao

import (
	"conscience-backend/config"
	"conscience-backend/model"
	"fmt"
	"github.com/pkg/errors"
	"time"

	"github.com/parnurzeal/gorequest"
)

var DefaultMeetingChain *MeetingChain
var DefaultStudyChain *StudyChain

func Init(conf config.BlockchainAPI) {
	api := &BlockchainAPI{
		network: "fabric",
		url:     conf.Url,
	}

	DefaultMeetingChain = NewMeetingChain(api)
	DefaultStudyChain = NewStudyChain(api)
}

type BlockchainAPI struct {
	network string
	url     string
}

func (api *BlockchainAPI) request(content interface{}, invoke bool) (*model.BlockchainResponse, error) {
	request := gorequest.New()
	var resp model.JsonResponse
	var bcResponse model.BlockchainResponse
	resp.Data = &bcResponse

	var url string
	if invoke {
		url = fmt.Sprintf("%s/api/v2/contract/invoke", api.url)
	} else {
		url = fmt.Sprintf("%s/api/v2/contract/query", api.url)
	}

	_, _, errs := request.Post(url).
		Send(content).
		Timeout(time.Second * 30).
		EndStruct(&resp)
	if len(errs) > 0 {
		return nil, errs[0]
	}
	if resp.ErrorCode != 0 {
		return nil, errors.New(resp.Description)
	}

	return &bcResponse, nil
}

func (api *BlockchainAPI) Invoke(accessToken, channel, contract, fn string, args []interface{}) (string, string, error) {
	request := model.RequestInvokeOrQueryContractV2{
		AccessToken: accessToken,
		NetworkId:   api.network,
		Channel:     channel,
		Contract:    contract,
		Sync:        true,
		Args: &model.Args{
			Method: fn,
			Params: args,
		},
	}
	blockchainResponse, err := api.request(request, true)
	if err != nil {
		return "", "", err
	}
	return blockchainResponse.TxID, blockchainResponse.Response, err
}
func (api *BlockchainAPI) Query(accessToken, channel, contract, fn string, args []interface{}) (string, error) {
	request := model.RequestInvokeOrQueryContractV2{
		AccessToken: accessToken,
		NetworkId:   api.network,
		Channel:     channel,
		Contract:    contract,
		Sync:        false,
		Args: &model.Args{
			Method: fn,
			Params: args,
		},
	}
	blockchainResponse, err := api.request(request, false)
	if err != nil {
		return "", err
	}
	return blockchainResponse.Response, err
}
