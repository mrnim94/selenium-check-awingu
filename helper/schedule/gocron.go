package schedule

import (
	"github.com/go-co-op/gocron"
	"time"
	"fmt"
)

type GoCron struct {
	GR *gocron.Scheduler
	ZoneName string
	SecondsOfUTC int
}

func (gc *GoCron) GetGoCron()  {
	gc.GR = gocron.NewScheduler(time.FixedZone(gc.ZoneName, gc.SecondsOfUTC))
	fmt.Println("Scheduler đã được khởi động")
}