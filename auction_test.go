package auction_test

import (
	"testing"

	"github.com/jhondev/auction"
)

func TestAuction(t *testing.T) {
	tests := []struct {
		name     string
		bidders  []*auction.Bidder
		wantName string
		wantBid  float64
	}{
		{
			name: "Auction #1",
			bidders: []*auction.Bidder{
				{Name: "Sasha", InitialBid: 50.00, MaxBid: 80.00, BidIncrement: 3.00},
				{Name: "John", InitialBid: 60.00, MaxBid: 82.00, BidIncrement: 2.00},
				{Name: "Pat", InitialBid: 55.00, MaxBid: 85.00, BidIncrement: 5.00},
			},
			wantName: "Pat",
			wantBid:  85.00,
		},
		{
			name: "Auction #2",
			bidders: []*auction.Bidder{
				{Name: "Riley", InitialBid: 700.00, MaxBid: 725.00, BidIncrement: 2.00},
				{Name: "Morgan", InitialBid: 599.00, MaxBid: 725.00, BidIncrement: 15.00},
				{Name: "Charlie", InitialBid: 625.00, MaxBid: 725.00, BidIncrement: 8.00},
			},
			wantName: "Riley",
			wantBid:  722.00,
		},
		{
			name: "Auction #3",
			bidders: []*auction.Bidder{
				{Name: "Alex", InitialBid: 2500.00, MaxBid: 3000.00, BidIncrement: 500.00},
				{Name: "Jesse", InitialBid: 2800.00, MaxBid: 3100.00, BidIncrement: 201.00},
				{Name: "Drew", InitialBid: 2501.00, MaxBid: 3200.00, BidIncrement: 247.00},
			},
			wantName: "Jesse",
			wantBid:  3001.00,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := auction.New(tt.bidders)
			a.Start()
			if a.GetWinningBidder().Name != tt.wantName {
				t.Errorf("got winner %s, want %s", a.GetWinningBidder().Name, tt.wantName)
			}
			if a.GetWinningBid() != tt.wantBid {
				t.Errorf("got winning bid %.2f, want %.2f", a.GetWinningBid(), tt.wantBid)
			}
		})
	}
}
