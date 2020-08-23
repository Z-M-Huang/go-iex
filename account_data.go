package iex

//Metadata result from /account/metadata
//https://iexcloud.io/docs/api/#metadata
type Metadata struct {
	PayAsYouGoEnabled    bool    `json:"payAsYouGoEnabled"`
	EffectiveDate        int64   `json:"effectiveDate"`
	SubscriptionTermType string  `json:"subscriptionTermType"`
	TierName             string  `json:"tierName"`
	MessageLimit         int     `json:"messageLimit"`
	MessagesUsed         int     `json:"messagesUsed"`
	CircuitBreaker       *uint64 `json:"circuitBreaker"`
}
