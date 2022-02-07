package model

type Preferences struct {
	// UserID        int      `json:"user_id"`
	TrackedTokens []string `json:"tracked_tokens"` // tracked crypto currencies
	FiatCurrency  string   `json:"fiat_currency"`  // e.g. "EUR", "UAH", the one that user uses to convert to
}
