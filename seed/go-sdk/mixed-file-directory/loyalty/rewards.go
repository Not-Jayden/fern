// This file was auto-generated by Fern from our API Definition.

package loyalty

import (
	fern "github.com/mixed-file-directory/fern"
)

type CreateLoyaltyRewardRequest struct {
	// The reward to create.
	Reward *fern.LoyaltyReward `json:"reward,omitempty" url:"-"`
	// A unique string that identifies this `CreateLoyaltyReward` request.
	// Keys can be any valid string, but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
}

type RedeemLoyaltyRewardRequest struct {
	// A unique string that identifies this `RedeemLoyaltyReward` request.
	// Keys can be any valid string, but must be unique for every request.
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The ID of the [location](entity:Location) where the reward is redeemed.
	LocationId string `json:"location_id" url:"-"`
}

type SearchLoyaltyRewardsRequest struct {
	// The search criteria for the request.
	// If empty, the endpoint retrieves all loyalty rewards in the loyalty program.
	Query *fern.SearchLoyaltyRewardsRequestLoyaltyRewardQuery `json:"query,omitempty" url:"-"`
	// The maximum number of results to return in the response. The default value is 30.
	Limit *int `json:"limit,omitempty" url:"-"`
	// A pagination cursor returned by a previous call to
	// this endpoint. Provide this to retrieve the next set of
	// results for the original query.
	// For more information,
	// see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
}
