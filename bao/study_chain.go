package bao

import "fmt"

type StudyChain struct {
	channel  string
	contract string
	api      *BlockchainAPI
}

func NewStudyChain(api *BlockchainAPI) *StudyChain {
	// TODO
	return &StudyChain{
		channel:  "studychain",
		contract: "",
		api:      api,
	}
}

func (c *StudyChain) AddPoints(accessToken string, points int) error {
	var args []interface{}
	args = append(args, points)

	txID, _, err := c.api.Invoke(accessToken, c.channel, c.contract, "PointsService.AddPoints", args)
	if err != nil {
		return err
	}
	fmt.Println(txID)
	return nil
}
