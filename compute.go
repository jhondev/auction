package auction

// Compute will compute the auction and return the winning bidder.
// Provided bidders must be sorted by initial bid
func Compute(bidders []*Bidder) *Bidder {
	// track the final bids count
	finalBidsCount := 0
	currentBid := 0.0
	var winningBidder *Bidder
	for {
		// track the current bid
		for _, bidder := range bidders {
			if bidder.maxBidReached {
				continue
			}
			// if the bidder has not placed a bid, set the initial bid
			if bidder.currentBid == 0 {
				bidder.currentBid = bidder.InitialBid
			}

			// calculate the auto incremented bid
			incremented := bidder.currentBid + bidder.BidIncrement
			// if the bidder has reached the max bid, skip
			if incremented > bidder.MaxBid || (bidder == winningBidder && finalBidsCount == len(bidders)-1) {
				finalBidsCount++
				bidder.maxBidReached = true
				continue
			}

			// add dollar amount to bidder's current bid if is in a losing position
			if bidder != winningBidder && bidder.currentBid <= currentBid {
				bidder.currentBid = incremented
			}
			// if the bidder has the highest bid, set the current bid as the current winning bid
			if bidder.currentBid > currentBid {
				currentBid = bidder.currentBid
				winningBidder = bidder
			}
		}
		if finalBidsCount == len(bidders) {
			return winningBidder
		}
	}
}
