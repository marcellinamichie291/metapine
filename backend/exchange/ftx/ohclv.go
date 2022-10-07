package ftx

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/DawnKosmos/metapine/backend/exchange"
)

func (f *FTX) OHCLV(ticker string, resolution int64, startTime time.Time, endTime time.Time) ([]exchange.Candle, error) {
	var hp []exchange.Candle
	st, et := startTime.Unix(), endTime.Unix()
	var end int64 = et

	if time.Now().Unix() < end {
		end = time.Now().Unix()
	}
	newRes := checkResolution(resolution)
	for {
		c, err := f.getOHCLV(ticker, newRes, st, end)
		if err != nil {
			fmt.Println(st, end)
			log.Printf("Error OHCLV FTX %v", err)
			return hp, err
		}
		fmt.Println(len(c))
		if len(c) == 0 {
			return hp, nil
		}

		hp = append(c, hp...)
		end = hp[0].StartTime.Unix() - 1
		time.Sleep(100 * time.Microsecond)
	}
	return hp, nil

}

/*
func (f *FTX) OHCLV(ticker string, resolution int64, st time.Time, et time.Time) ([]exchange.Candle, error) {
	var historicalPrices []exchange.Candle
	var end int64 = 0
	startTime := st.Unix()
	endTime := et.Unix()

	if time.Now().Unix() < endTime {
		endTime = time.Now().Unix()
	}

	newRes := checkResolution(resolution)

	for startTime < endTime {
		end = startTime + newRes*1500
		if end >= endTime {
			c, err := f.getOHCLV(ticker, int64(newRes), startTime, endTime)
			if err != nil {
				log.Printf("Error OHCLV FTX %v", err)
				return historicalPrices, err
			}
			historicalPrices = append(historicalPrices, c...)
		} else {
			c, err := f.getOHCLV(ticker, int64(newRes), startTime, end)
			if err != nil {
				log.Printf("Error OHCLV FTX %v", err)
				return historicalPrices, err
			}
			historicalPrices = append(historicalPrices, c...)

		}
		startTime = startTime + int64(newRes*1501)
	}

	kek, err := exchange.ConvertChartResolution(newRes, resolution, historicalPrices)
	return kek, err
}
*/

type HistoricalPrices struct {
	Success bool              `json:"success"`
	Result  []exchange.Candle `json:"result"`
}

func (f *FTX) getOHCLV(ticker string, res int64, st int64, et int64) ([]exchange.Candle, error) {
	var h HistoricalPrices
	resp, err := f.get(
		"markets/"+ticker+
			"/candles?resolution="+strconv.FormatInt(res, 10)+
			"&start_time="+strconv.FormatInt(st, 10)+
			"&end_time="+strconv.FormatInt(et, 10),
		[]byte(""))
	if err != nil {
		log.Printf("Error OHCLV FTX %v", err)
		return h.Result, err
	}
	err = processResponse(resp, &h)
	return h.Result, nil

}

// checkResolution looking if the asked resolution is a valid one
func checkResolution(res int64) int64 {
	fn := exchange.GenerateResolutionFunc(86400*7, 86400*2, 86400, 14400,
		3600, 900, 300, 60, 15)
	return fn(res)
	/*
		var newRes int64
		if res == 3600 || res == 14400 || res == 86400 || res == 300 || res == 60 || res == 900 {
			newRes = res
			return newRes
		}

		if res >= 86400 && res%86400 == 0 {
			return 86400
		}

		if res >= 14400 && res%14400 == 0 {
			return 14400
		}

		if res >= 3600 && res%3600 == 0 {
			return 3600
		}
		if res >= 300 && res%300 == 0 {
			return 300
		}
		if res >= 900 && res%900 == 0 {
			return 900
		}

		if res >= 15 && res%15 == 0 {
			return 15
		}

		return 3600
	*/
}
