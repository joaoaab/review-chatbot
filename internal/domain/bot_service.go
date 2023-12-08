package domain

import (
	"context"
	"fmt"
	"github.com/grassmudhorses/vader-go"
	"log"
	"regexp"
	"review-bot/internal/infrastructure/repository/bot_log"
	"review-bot/internal/infrastructure/repository/customer"
	"review-bot/internal/models"
	"review-bot/internal/state"
	"review-bot/internal/structs"
	"strconv"
)

type BotService struct {
	StateModule      *state.StateMachineHandler
	ReviewRepository *customer.CustomerRepository
	BotLogRepository *bot_log.BotLogRepository
}

var numberExpression = regexp.MustCompile("[0-9]+")

func NewBotService(state *state.StateMachineHandler, repo *customer.CustomerRepository, botLogRepo *bot_log.BotLogRepository) *BotService {
	return &BotService{StateModule: state, ReviewRepository: repo, BotLogRepository: botLogRepo}
}

func (bs BotService) HandleMessage(message string, id uint) error {
	customerModel, err := bs.ReviewRepository.GetCustomer(id)
	if err != nil {
		return err
	}

	bs.BotLogRepository.AppendMessage(customerModel.Review.ID, structs.Message{
		IsUserMessage: true,
		Text:          message,
	})
	machine, _ := bs.StateModule.GetStateMachineFromChat(customerModel.Review.ID)
	currentState, _ := machine.State(context.Background())
	switch currentState {
	case state.StateWaitingSentiment:
		bs.handleWaitingSentiment(message, customerModel.Review)
	case state.StateWaitingReview:
		bs.handleWaitingReview(message, customerModel.Review)
	}

	return nil
}

func (bs BotService) handleWaitingSentiment(message string, customerReview models.CustomerReview) {
	analysis := vader.GetSentiment(message)
	if analysis.Compound >= 0.05 {
		bs.BotLogRepository.AppendMessage(customerReview.ID, structs.Message{
			IsUserMessage: false,
			Text:          fmt.Sprintf("Fantastic! On a scale of 1-5, how would you rate the %s?", customerReview.ProductName),
		})
		bs.BotLogRepository.ChangeState(customerReview.ID, state.StateWaitingReview)
	} else {
		bs.BotLogRepository.AppendMessage(customerReview.ID, structs.Message{
			IsUserMessage: false,
			Text:          fmt.Sprintf("Ok! hope you're enjoying your %s", customerReview.ProductName),
		})
		bs.BotLogRepository.ChangeState(customerReview.ID, state.StateGoodbyeMessage)
	}
}

func (bs BotService) handleWaitingReview(message string, customerReview models.CustomerReview) {
	review := numberExpression.FindString(message)
	reviewRating, err := strconv.Atoi(review)
	if err != nil || reviewRating > 5 || reviewRating < 1 {
		log.Println("Couldn't get a number from message")
		bs.BotLogRepository.AppendMessage(customerReview.ID, structs.Message{
			IsUserMessage: false,
			Text:          "I'm sorry, i didn't receive a number from 1-5",
		})
	} else {
		bs.ReviewRepository.SetReviewScore(customerReview.ID, reviewRating)
		bs.handleFinish(customerReview.ID)
	}
}

func (bs BotService) handleFinish(customerReviewID uint) {
	bs.BotLogRepository.ChangeState(customerReviewID, state.StateFinished)
	bs.BotLogRepository.AppendMessage(customerReviewID, structs.Message{
		IsUserMessage: false,
		Text:          "Thank you for sharing your feedback! If you have any more thoughts or need assistance with anything else, feel free to reach out!",
	})
}
