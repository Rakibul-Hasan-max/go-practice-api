package database

var reviewsList []Review

type Review struct {
	ID        int     `json:"id"`
	ProductID int     `json:"productId"`
	Reviewer  string  `json:"reviewer"`
	Rating    float64 `json:"rating"`
	Comment   string  `json:"comment"`
}

// Store new review
func StoreReview(r Review) Review {
	r.ID = len(reviewsList) + 1
	reviewsList = append(reviewsList, r)
	return r
}

// List all reviews
func ListReviews() []Review {
	return reviewsList
}

// Update review
func UpdateReview(review Review) {
	for idx, r := range reviewsList {
		if r.ID == review.ID {
			reviewsList[idx] = review
		}
	}
}

// Delete review by ID
func DeleteReview(reviewID int) {
	var tempList []Review
	for _, r := range reviewsList {
		if r.ID != reviewID {
			tempList = append(tempList, r)
		}
	}
	reviewsList = tempList
}

// Demo reviews init
func init() {
	r1 := Review{
		ID:        1,
		ProductID: 1,
		Reviewer:  "Alice",
		Comment:   "Really fresh and tasty orange!",
		Rating:    5,
	}

	r2 := Review{
		ID:        2,
		ProductID: 2,
		Reviewer:  "Bob",
		Comment:   "Apple was good but a bit expensive.",
		Rating:    4,
	}

	r3 := Review{
		ID:        3,
		ProductID: 3,
		Reviewer:  "Charlie",
		Comment:   "Banana was ripe and delicious.",
		Rating:    5,
	}

	reviewsList = append(reviewsList, r1, r2, r3)
}
