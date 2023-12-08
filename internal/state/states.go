package state

import (
	"github.com/qmuntal/stateless"
	"review-bot/internal/infrastructure/repository/customer"
)

const (
	StateStarted          = "started"
	StateWaitingSentiment = "waiting_sentiment"
	StateWaitingReview    = "waiting_review"
	StateGoodbyeMessage   = "goodbye"
	StateFinished         = "finished"
)

const (
	TriggerWaitSentiment = "wait_sentiment"
	TriggerWaitReview    = "wait_review"
	TriggerGoodbye       = "go_to_goodbye"
	TriggerFinish        = "finish"
)

type StateMachineHandler struct {
	repository *customer.CustomerRepository
}

func NewStateMachineHandler(repo *customer.CustomerRepository) *StateMachineHandler {
	return &StateMachineHandler{repository: repo}
}

func (smh StateMachineHandler) GetStateMachineFromChat(customerReviewID uint) (*stateless.StateMachine, error) {
	review, err := smh.repository.GetReview(customerReviewID)
	if err != nil {
		return &stateless.StateMachine{}, err
	}

	currentState := review.Log.State
	machine := stateless.NewStateMachine(currentState)
	machine.Configure(StateStarted).Permit(TriggerWaitSentiment, StateWaitingSentiment)
	machine.Configure(StateWaitingSentiment).Permit(TriggerGoodbye, StateGoodbyeMessage)
	machine.Configure(StateWaitingSentiment).Permit(TriggerWaitReview, StateWaitingReview)
	machine.Configure(StateWaitingReview).Permit(TriggerFinish, StateFinished)
	return machine, nil
}
