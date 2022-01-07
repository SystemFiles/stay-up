package tasks

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/systemfiles/stay-up/api/models"
	"github.com/systemfiles/stay-up/api/provider"
	"github.com/systemfiles/stay-up/api/util"
)

func InitBackgroundServiceRefresh(ctx context.Context, refreshTime time.Duration) error {
	select {
	case <- ctx.Done():
		return ctx.Err()
	default:
		// perform repeating synchronization
		for {
			var services []models.Service
			var wg sync.WaitGroup

			// get all services
			err := provider.GetAllServices(&services)
			if err != nil {
				return err
			}

			// for each service perform service tasks
			for _,svc := range services {
				wg.Add(1)
				svcLocal := svc // create a local copy of svc pointer

				go func() {
					defer wg.Done()
					serviceWorker(&svcLocal)
				}()
			}

			wg.Wait()

			// refresh time
			time.Sleep(5000 * time.Millisecond)
		}
	}
}

func serviceWorker(svc *models.Service) {
	db, err := util.GetDBInstance()
	if err != nil {
		fmt.Printf("Failed to get database instance. %s", err.Error())
		return
	}

	if err := svc.CheckAndUpdateStatus(); err != nil {
		fmt.Printf("Failed to get status for service: %s. Reason: %s", svc.Name, err.Error())
		return
	}

	db.Save(&svc)
}