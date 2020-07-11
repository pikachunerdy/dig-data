package main

import (
	"flag"
	"github.com/diadata-org/diadata/pkg/dia"
	"sync"

	defiscraper "github.com/diadata-org/diadata/internal/pkg/defiscrapers"
	models "github.com/diadata-org/diadata/pkg/model"
	log "github.com/sirupsen/logrus"
)

// handleDefiInterestRate delegates rate information to Kafka
func handleDefiInterestRate(c chan *dia.DefiLendingRate, wg *sync.WaitGroup, ds models.Datastore) {
	defer wg.Done()
	// Pull from channel as long as not empty
	for {
		t, ok := <-c
		if !ok {
			log.Error("error")
			return
		}
		ds.SetDefiLendingRateInflux(t)
	}
}

func main() {
	rateType := flag.String("type", "DYDX", "Type of Defi rate")
	flag.Parse()

	wg := sync.WaitGroup{}
	ds, err := models.NewDataStore()

	if err != nil {
		log.Errorln("NewDataStore:", err)
	} else {

		sRate := defiscraper.SpawnDefiScraper(ds, *rateType)
		defer sRate.Close()

		// Send rates to the database while the scraper scrapes
		wg.Add(1)
		go handleDefiInterestRate(sRate.Channel(), &wg, ds)
		defer wg.Wait()
	}
}
