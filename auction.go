package auction

import (
	"sort"
)

// Bidder represents a bidder in the auction
type Bidder struct {
	Name          string
	InitialBid    float64
	MaxBid        float64
	BidIncrement  float64
	currentBid    float64
	maxBidReached bool
}

// Auction handles the auction state and computation
type Auction struct {
	Bidders       []*Bidder
	WinningBidder *Bidder
	Computed      bool
}

// New creates a new auction with the provided bidders
func New(bidders []*Bidder) *Auction {
	return &Auction{
		Bidders: bidders,
	}
}

// Start will start the auction with the bidders sorted and compute the winning bidder
func (a *Auction) Start() {
	// sort the bidders by initial bid to start the auction
	sort.Slice(a.Bidders, func(i, j int) bool {
		return a.Bidders[i].InitialBid < a.Bidders[j].InitialBid
	})
	a.WinningBidder = Compute(a.Bidders)
	a.Computed = true
}

// GetWinningBid will return the winning bid, it will start the auction if it has not been started
func (a *Auction) GetWinningBid() float64 {
	if !a.Computed {
		a.Start()
	}
	return a.WinningBidder.currentBid
}

// GetWinningBidder will return the winning bidder, it will start the auction if it has not been started
func (a *Auction) GetWinningBidder() *Bidder {
	if !a.Computed {
		a.Start()
	}
	return a.WinningBidder
}
