package sentinelwatchlist


type SentinelWatchlistTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_watchlist#create SentinelWatchlist#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_watchlist#delete SentinelWatchlist#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/sentinel_watchlist#read SentinelWatchlist#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

