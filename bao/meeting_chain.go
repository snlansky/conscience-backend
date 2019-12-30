package bao

type MeetingChain struct {
	channel  string
	contract string
	api      *BlockchainAPI
}

func NewMeetingChain(api *BlockchainAPI) *MeetingChain {
	return &MeetingChain{
		channel:  "meetingchain",
		contract: "faceid",
		api:      api,
	}
}
