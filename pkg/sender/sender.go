package sender

import "github.com/corporateanon/barker/pkg/types"

type Sender interface {
	Send(bot *types.Bot, campaign *types.Campaign, user *types.User) error
}
