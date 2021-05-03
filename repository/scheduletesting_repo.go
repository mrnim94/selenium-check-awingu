package repository

import (
	"context"
	"selenium-check-awingu/model"
)

type ScheduleTestingRepo interface {
	SaveSignalSchedule(context context.Context, signalKey model.SignalKey) (model.SignalKey, error)
	SelectSignalScheduleBySignId(context context.Context, signalId string) (string, error)
	UpdateSignalKey(context context.Context, signalKey model.SignalKey) (string, error)
	SaveScheduleTesting(context context.Context, scheduleT model.ScheduleTesting) (model.ScheduleTesting, error)

	SelectAllScheduleDifferenceSignalKey(context context.Context, signalKey string) ([]model.ScheduleTesting, error)
	UpdateSignalKeyForSchedule(context context.Context, newSignalKey string, scheduleT model.ScheduleTesting) (model.ScheduleTesting, error)
	SelectAllScheduleTesting(context context.Context) ([]model.ScheduleTesting, error)
	RemoveScheduleById(context context.Context, scheduleId string) error
}
