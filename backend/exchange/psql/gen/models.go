// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.15.0

package gen

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Exchanges string

const (
	ExchangesFtx      Exchanges = "ftx"
	ExchangesBinance  Exchanges = "binance"
	ExchangesBybit    Exchanges = "bybit"
	ExchangesDeribit  Exchanges = "deribit"
	ExchangesBitmex   Exchanges = "bitmex"
	ExchangesCoinbase Exchanges = "coinbase"
	ExchangesPhemex   Exchanges = "phemex"
)

func (e *Exchanges) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Exchanges(s)
	case string:
		*e = Exchanges(s)
	default:
		return fmt.Errorf("unsupported scan type for Exchanges: %T", src)
	}
	return nil
}

type NullExchanges struct {
	Exchanges Exchanges
	Valid     bool // Valid is true if String is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullExchanges) Scan(value interface{}) error {
	if value == nil {
		ns.Exchanges, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Exchanges.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullExchanges) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return ns.Exchanges, nil
}

type Binacelow struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Binancehigh struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Deribithigh struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Deribitlow struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Ftxhigh struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Ftxlow struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Index struct {
	IndexID int32
	Name    string
}

type Indexhigh struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Indexlow struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type MinuteChart struct {
	Starttime time.Time
	Open      float32
	High      float32
	Close     float32
	Low       float32
	Volume    float32
}

type MinuteManager struct {
	IndexID   int32
	Tablename string
	Dataarr   sql.NullString
}

type Ohclv struct {
	IndexID    int32
	Resolution int32
	Starttime  time.Time
	Open       float32
	High       float32
	Close      float32
	Low        float32
	Volume     float32
}

type Ticker struct {
	TickerID int32
	Exchange Exchanges
	Ticker   string
}

type TickerIndex struct {
	TickerID      int32
	IndexID       int32
	Weight        int32
	Excludevolume bool
}
