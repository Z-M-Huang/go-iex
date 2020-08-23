package iex

//Book https://iexcloud.io/docs/api/#book
type Book struct {
	Quote       Quote       `json:"quote"`
	Bids        []BidAsk    `json:"bids"`
	Asks        []BidAsk    `json:"asks"`
	Trades      []Trade     `json:"trades"`
	SystemEvent SystemEvent `json:"systemEvent"`
}

//Quote https://iexcloud.io/docs/api/#quote
type Quote struct {
	Symbol                 string     `json:"symbol"`
	CompanyName            string     `json:"companyName"`
	CalculationPrice       string     `json:"calculationPrice"`
	Open                   float64    `json:"open"`
	OpenTime               EpochTime  `json:"openTime"`
	OpenSource             string     `json:"openSource"`
	Close                  float64    `json:"close"`
	CloseTime              EpochTime  `json:"closeTime"`
	CloseSource            string     `json:"closeSource"`
	High                   float64    `json:"high"`
	HighTime               EpochTime  `json:"highTime"`
	HighSource             string     `json:"highSource"`
	Low                    float64    `json:"low"`
	LowTime                EpochTime  `json:"lowTime"`
	LowSource              string     `json:"lowSource"`
	LatestPrice            float64    `json:"latestPrice"`
	LatestSource           string     `json:"latestSource"`
	LatestTime             string     `json:"latestTime"`
	LatestUpdate           EpochTime  `json:"latestUpdate"`
	LatestVolume           int        `json:"latestVolume"`
	Volume                 int        `json:"volume"`
	IexRealtimePrice       float64    `json:"iexRealtimePrice"`
	IexRealtimeSize        int        `json:"iexRealtimeSize"`
	IexLastUpdated         EpochTime  `json:"iexLastUpdated"`
	DelayedPrice           float64    `json:"delayedPrice"`
	DelayedPriceTime       EpochTime  `json:"delayedPriceTime"`
	OddLotDelayedPrice     float64    `json:"oddLotDelayedPrice"`
	OddLotDelayedPriceTime EpochTime  `json:"oddLotDelayedPriceTime"`
	ExtendedPrice          float64    `json:"extendedPrice"`
	ExtendedChange         float64    `json:"extendedChange"`
	ExtendedChangePercent  float64    `json:"extendedChangePercent"`
	ExtendedPriceTime      EpochTime  `json:"extendedPriceTime"`
	PreviousClose          float64    `json:"previousClose"`
	PreviousVolume         int        `json:"previousVolume"`
	Change                 float64    `json:"change"`
	ChangePercent          float64    `json:"changePercent"`
	IexMarketPercent       *float64   `json:"iexMarketPercent"`
	IexVolume              *int       `json:"iexVolume"`
	AvgTotalVolume         int        `json:"avgTotalVolume"`
	IexBidPrice            *float64   `json:"iexBidPrice"`
	IexBidSize             *int       `json:"iexBidSize"`
	IexAskPrice            *float64   `json:"iexAskPrice"`
	IexAskSize             *int       `json:"iexAskSize"`
	IexOpen                *float64   `json:"iexOpen"`
	IexOpenTime            *EpochTime `json:"iexOpenTime"`
	IexClose               float64    `json:"iexClose"`
	IexCloseTime           EpochTime  `json:"iexCloseTime"`
	MarketCap              int64      `json:"marketCap"`
	Week52High             float64    `json:"week52High"`
	Week52Low              float64    `json:"week52Low"`
	YtdChange              float64    `json:"ytdChange"`
	PeRatio                float64    `json:"peRatio"`
	LastTradeTime          EpochTime  `json:"lastTradeTime"`
	IsUSMarketOpen         bool       `json:"isUSMarketOpen"`
}

// BidAsk models a bid or an ask for a quote.
type BidAsk struct {
	Price     float64   `json:"price"`
	Size      int       `json:"size"`
	Timestamp EpochTime `json:"timestamp"`
}

//Trade models a trade for a quote.
type Trade struct {
	Price                 float64   `json:"price"`
	Size                  int       `json:"size"`
	TradeID               int       `json:"tradeId"`
	IsISO                 bool      `json:"isISO"`
	IsOddLot              bool      `json:"isOddLot"`
	IsOutsideRegularHours bool      `json:"isOutsideRegularHours"`
	IsSinglePriceCross    bool      `json:"isSinglePriceCross"`
	IsTradeThroughExempt  bool      `json:"isTradeThroughExempt"`
	Timestamp             EpochTime `json:"timestamp"`
}

//HistoricalPrice for https://iexcloud.io/docs/api/#historical-prices
type HistoricalPrice struct {
	Date           string  `json:"date"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Volume         int     `json:"volume"`
	UOpen          float64 `json:"uOpen"`
	UClose         float64 `json:"uClose"`
	UHigh          float64 `json:"uHigh"`
	ULow           float64 `json:"uLow"`
	UVolume        int     `json:"uVolume"`
	Change         int     `json:"change"`
	ChangePercent  int     `json:"changePercent"`
	Label          string  `json:"label"`
	ChangeOverTime float64 `json:"changeOverTime"`
}

//IntradayPrice for https://iexcloud.io/docs/api/#intraday-prices
type IntradayPrice struct {
	Date                 string   `json:"date"`
	Minute               string   `json:"minute"`
	MarketAverate        *float64 `json:"marketAverage"`
	MarketNotional       *float64 `json:"marketNotional"`
	MarketNumberOfTrades *float64 `json:"marketNumberOfTrades"`
	MarketOpen           *float64 `json:"marketOpen"`
	MarketClose          *float64 `json:"marketClose"`
	MarketHigh           *float64 `json:"marketHigh"`
	MarketLow            *float64 `json:"marketLow"`
	MarketVolume         *int     `json:"marketVolume"`
	MarketChangeOverTime *float64 `json:"marketChangeOverTime"`
	ChangeOverTime       *float64 `json:"changeOverTime"`
	Label                string   `json:"label"`
	High                 float64  `json:"high"`
	Low                  float64  `json:"low"`
	Open                 float64  `json:"open"`
	Close                float64  `json:"close"`
	Average              float64  `json:"average"`
	Volume               int      `json:"volume"`
	Notional             float64  `json:"notional"`
	NumberOfTrades       int      `json:"numberOfTrades"`
}

//DelayedQuote https://iexcloud.io/docs/api/#delayed-quote
type DelayedQuote struct {
	Symbol           string    `json:"symbol"`
	DelayedPrice     float64   `json:"delayedPrice"`
	DelayedSize      int       `json:"delayedSize"`
	DelayedPriceTime EpochTime `json:"delayedPriceTime"`
	High             float64   `json:"high"`
	Low              float64   `json:"low"`
	TotalVolume      int       `json:"totalVolume"`
	ProcessedTime    EpochTime `json:"processedTime"`
}

//LargestTrade for https://iexcloud.io/docs/api/#largest-trades
type LargestTrade struct {
	Price     float64   `json:"price"`
	Size      int       `json:"size"`
	Time      EpochTime `json:"time"`
	TimeLabel string    `json:"timeLabel"`
	Venue     string    `json:"venue"`
	VenueName string    `json:"venueName"`
}

//OHLC for https://iexcloud.io/docs/api/#open-close-price
type OHLC struct {
	Open   OpenClose `json:"open"`
	Close  OpenClose `json:"close"`
	High   float64   `json:"high"`
	Low    float64   `json:"low"`
	Volume int       `json:"volume"`
	Symbol string    `json:"symbol"`
}

//OpenClose open/close price
type OpenClose struct {
	Price float64   `json:"price"`
	Time  EpochTime `json:"time"`
}

//PreviousDayPrice for https://iexcloud.io/docs/api/#previous-day-price
type PreviousDayPrice struct {
	Date           string  `json:"date"`
	Open           float64 `json:"open"`
	Close          float64 `json:"close"`
	High           float64 `json:"high"`
	Low            float64 `json:"low"`
	Volume         int     `json:"volume"`
	UOpen          float64 `json:"uOpen"`
	UClose         float64 `json:"uClose"`
	UHigh          float64 `json:"uHigh"`
	ULow           float64 `json:"uLow"`
	UVolume        int     `json:"uVolume"`
	Change         int     `json:"change"`
	ChangePercent  int     `json:"changePercent"`
	ChangeOverTime int     `json:"changeOverTime"`
	Symbol         string  `json:"symbol"`
}

//VolumeByVenue for https://iexcloud.io/docs/api/#volume-by-venue
type VolumeByVenue struct {
	Volume           int     `json:"volume"`
	Venue            string  `json:"venue"`
	VenueName        string  `json:"venueName"`
	Date             string  `json:"date"`
	MarketPercent    float64 `json:"marketPercent"`
	AvgMarketPercent float64 `json:"avgMarketPercent"`
}
