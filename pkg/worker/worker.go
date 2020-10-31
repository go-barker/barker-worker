package worker

import (
	"errors"
	"time"

	"github.com/corporateanon/barker-worker/pkg/sender"
	"github.com/corporateanon/barker/pkg/dao"
	"github.com/corporateanon/barker/pkg/types"
)

var ErrorBotEmpty = errors.New("Bot empty")
var ErrorDeliveryEmpty = errors.New("Delivery empty")
var ErrorSendFailed = errors.New("Send failed")

type Worker interface {
	Loop() error
}

type WorkerImpl struct {
	botDao      dao.BotDao
	deliveryDao dao.DeliveryDao
	sender      sender.Sender
}

func NewWorkerImpl(botDao dao.BotDao, deliveryDao dao.DeliveryDao, sender sender.Sender) Worker {
	return &WorkerImpl{
		botDao:      botDao,
		deliveryDao: deliveryDao,
		sender:      sender,
	}
}

func (w *WorkerImpl) Loop() error {
	for {
		w.tick()
		time.Sleep(500 * time.Millisecond)
	}
}

func (w *WorkerImpl) tick() error {
	bot, err := w.botDao.RRTake()
	if err != nil {
		return err
	}
	if bot == nil {
		return ErrorBotEmpty
	}

	deliveryTakeResult, err := w.deliveryDao.Take(bot.ID, 0, 0)
	if err != nil {
		return err
	}
	if deliveryTakeResult == nil {
		return ErrorDeliveryEmpty
	}

	err = w.sender.Send(bot, deliveryTakeResult.Campaign, deliveryTakeResult.User)
	if err != nil {
		w.deliveryDao.SetState(deliveryTakeResult.Delivery, types.DeliveryStateFail)
		return ErrorSendFailed
	}

	w.deliveryDao.SetState(deliveryTakeResult.Delivery, types.DeliveryStateSuccess)

	return err
}
