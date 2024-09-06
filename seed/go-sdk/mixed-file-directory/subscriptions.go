// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mixed-file-directory/fern/core"
)

type SwapPlanRequest struct {
	// The ID of the new subscription plan.
	NewPlanId string `json:"new_plan_id" url:"-"`
}

type CreateSubscriptionRequest struct {
	// A unique string that identifies this `CreateSubscription` request.
	// If you do not provide a unique string (or provide an empty string as the value),
	// the endpoint treats each request as independent.
	//
	// For more information, see [Idempotency keys](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey *string `json:"idempotency_key,omitempty" url:"-"`
	// The ID of the location the subscription is associated with.
	LocationId string `json:"location_id" url:"-"`
	// The ID of the subscription plan created using the Catalog API.
	// For more information, see
	// [Set Up and Manage a Subscription Plan](https://developer.squareup.com/docs/subscriptions-api/setup-plan) and
	// [Subscriptions Walkthrough](https://developer.squareup.com/docs/subscriptions-api/walkthrough).
	PlanId string `json:"plan_id" url:"-"`
	// The ID of the [customer](entity:Customer) subscribing to the subscription plan.
	CustomerId string `json:"customer_id" url:"-"`
	// The `YYYY-MM-DD`-formatted date to start the subscription.
	// If it is unspecified, the subscription starts immediately.
	StartDate *string `json:"start_date,omitempty" url:"-"`
	// The `YYYY-MM-DD`-formatted date when the newly created subscription is scheduled for cancellation.
	//
	// This date overrides the cancellation date set in the plan configuration.
	// If the cancellation date is earlier than the end date of a subscription cycle, the subscription stops
	// at the canceled date and the subscriber is sent a prorated invoice at the beginning of the canceled cycle.
	//
	// When the subscription plan of the newly created subscription has a fixed number of cycles and the `canceled_date`
	// occurs before the subscription plan expires, the specified `canceled_date` sets the date when the subscription
	// stops through the end of the last cycle.
	CanceledDate *string `json:"canceled_date,omitempty" url:"-"`
	// The tax to add when billing the subscription.
	// The percentage is expressed in decimal form, using a `'.'` as the decimal
	// separator and without a `'%'` sign. For example, a value of 7.5
	// corresponds to 7.5%.
	TaxPercentage *string `json:"tax_percentage,omitempty" url:"-"`
	// A custom price to apply for the subscription. If specified,
	// it overrides the price configured by the subscription plan.
	PriceOverrideMoney *Money `json:"price_override_money,omitempty" url:"-"`
	// The ID of the [subscriber's](entity:Customer) [card](entity:Card) to charge.
	// If it is not specified, the subscriber receives an invoice via email. For an example to
	// create a customer profile for a subscriber and add a card on file, see [Subscriptions Walkthrough](https://developer.squareup.com/docs/subscriptions-api/walkthrough).
	CardId *string `json:"card_id,omitempty" url:"-"`
	// The timezone that is used in date calculations for the subscription. If unset, defaults to
	// the location timezone. If a timezone is not configured for the location, defaults to "America/New_York".
	// Format: the IANA Timezone Database identifier for the location timezone. For
	// a list of time zones, see [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones).
	Timezone *string `json:"timezone,omitempty" url:"-"`
	// The origination details of the subscription.
	Source *SubscriptionSource `json:"source,omitempty" url:"-"`
}

type SubscriptionsGetRequest struct {
	// A query parameter to specify related information to be included in the response.
	//
	// The supported query parameter values are:
	//
	// - `actions`: to include scheduled actions on the targeted subscription.
	Include *string `json:"-" url:"include,omitempty"`
}

type SubscriptionsListEventsRequest struct {
	// When the total number of resulting subscription events exceeds the limit of a paged response,
	// specify the cursor returned from a preceding response here to fetch the next set of results.
	// If the cursor is unset, the response contains the last page of the results.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// The upper limit on the number of subscription events to return
	// in a paged response.
	Limit *int `json:"-" url:"limit,omitempty"`
}

type PauseSubscriptionRequest struct {
	// The `YYYY-MM-DD`-formatted date when the scheduled `PAUSE` action takes place on the subscription.
	//
	// When this date is unspecified or falls within the current billing cycle, the subscription is paused
	// on the starting date of the next billing cycle.
	PauseEffectiveDate *string `json:"pause_effective_date,omitempty" url:"-"`
	// The number of billing cycles the subscription will be paused before it is reactivated.
	//
	// When this is set, a `RESUME` action is also scheduled to take place on the subscription at
	// the end of the specified pause cycle duration. In this case, neither `resume_effective_date`
	// nor `resume_change_timing` may be specified.
	PauseCycleDuration *int64 `json:"pause_cycle_duration,omitempty" url:"-"`
	// The date when the subscription is reactivated by a scheduled `RESUME` action.
	// This date must be at least one billing cycle ahead of `pause_effective_date`.
	ResumeEffectiveDate *string `json:"resume_effective_date,omitempty" url:"-"`
	// The timing whether the subscription is reactivated immediately or at the end of the billing cycle, relative to
	// `resume_effective_date`.
	// See [ChangeTiming](#type-changetiming) for possible values
	ResumeChangeTiming *ChangeTiming `json:"resume_change_timing,omitempty" url:"-"`
	// The user-provided reason to pause the subscription.
	PauseReason *string `json:"pause_reason,omitempty" url:"-"`
}

type ResumeSubscriptionRequest struct {
	// The `YYYY-MM-DD`-formatted date when the subscription reactivated.
	ResumeEffectiveDate *string `json:"resume_effective_date,omitempty" url:"-"`
	// The timing to resume a subscription, relative to the specified
	// `resume_effective_date` attribute value.
	// See [ChangeTiming](#type-changetiming) for possible values
	ResumeChangeTiming *ChangeTiming `json:"resume_change_timing,omitempty" url:"-"`
}

type SearchSubscriptionsRequest struct {
	// When the total number of resulting subscriptions exceeds the limit of a paged response,
	// specify the cursor returned from a preceding response here to fetch the next set of results.
	// If the cursor is unset, the response contains the last page of the results.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"cursor,omitempty" url:"-"`
	// The upper limit on the number of subscriptions to return
	// in a paged response.
	Limit *int `json:"limit,omitempty" url:"-"`
	// A subscription query consisting of specified filtering conditions.
	//
	// If this `query` field is unspecified, the `SearchSubscriptions` call will return all subscriptions.
	Query *SearchSubscriptionsQuery `json:"query,omitempty" url:"-"`
	// An option to include related information in the response.
	//
	// The supported values are:
	//
	// - `actions`: to include scheduled actions on the targeted subscriptions.
	Include []string `json:"include,omitempty" url:"-"`
}

// Defines output parameters in a response from the
// [CancelSubscription](api-endpoint:Subscriptions-CancelSubscription) endpoint.
type CancelSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The specified subscription scheduled for cancellation according to the action created by the request.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`
	// A list of a single `CANCEL` action scheduled for the subscription.
	Actions []*SubscriptionAction `json:"actions,omitempty" url:"actions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CancelSubscriptionResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CancelSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CancelSubscriptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CancelSubscriptionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CancelSubscriptionResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Supported timings when a pending change, as an action, takes place to a subscription.
type ChangeTiming string

const (
	ChangeTimingDefaultChangeTimingTypeDoNotUse ChangeTiming = "DEFAULT_CHANGE_TIMING_TYPE_DO_NOT_USE"
	ChangeTimingImmediate                       ChangeTiming = "IMMEDIATE"
	ChangeTimingEndOfBillingCycle               ChangeTiming = "END_OF_BILLING_CYCLE"
)

func NewChangeTimingFromString(s string) (ChangeTiming, error) {
	switch s {
	case "DEFAULT_CHANGE_TIMING_TYPE_DO_NOT_USE":
		return ChangeTimingDefaultChangeTimingTypeDoNotUse, nil
	case "IMMEDIATE":
		return ChangeTimingImmediate, nil
	case "END_OF_BILLING_CYCLE":
		return ChangeTimingEndOfBillingCycle, nil
	}
	var t ChangeTiming
	return "", fmt.Errorf("%s is not a valid %T", s, t)
}

func (c ChangeTiming) Ptr() *ChangeTiming {
	return &c
}

// Defines output parameters in a response from the
// [CreateSubscription](api-endpoint:Subscriptions-CreateSubscription) endpoint.
type CreateSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The newly created subscription.
	//
	// For more information, see
	// [Subscription object](https://developer.squareup.com/docs/subscriptions-api/overview#subscription-object).
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (c *CreateSubscriptionResponse) GetExtraProperties() map[string]interface{} {
	return c.extraProperties
}

func (c *CreateSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler CreateSubscriptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*c = CreateSubscriptionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *c)
	if err != nil {
		return err
	}
	c.extraProperties = extraProperties

	c._rawJSON = json.RawMessage(data)
	return nil
}

func (c *CreateSubscriptionResponse) String() string {
	if len(c._rawJSON) > 0 {
		if value, err := core.StringifyJSON(c._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(c); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", c)
}

// Defines output parameters in a response of the [DeleteSubscriptionAction](api-endpoint:Subscriptions-DeleteSubscriptionAction)
// endpoint.
type DeleteSubscriptionActionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The subscription that has the specified action deleted.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (d *DeleteSubscriptionActionResponse) GetExtraProperties() map[string]interface{} {
	return d.extraProperties
}

func (d *DeleteSubscriptionActionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler DeleteSubscriptionActionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*d = DeleteSubscriptionActionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *d)
	if err != nil {
		return err
	}
	d.extraProperties = extraProperties

	d._rawJSON = json.RawMessage(data)
	return nil
}

func (d *DeleteSubscriptionActionResponse) String() string {
	if len(d._rawJSON) > 0 {
		if value, err := core.StringifyJSON(d._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(d); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", d)
}

// Defines output parameters in a response from the
// [RetrieveSubscription](api-endpoint:Subscriptions-RetrieveSubscription) endpoint.
type GetSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The subscription retrieved.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetSubscriptionResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetSubscriptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetSubscriptionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetSubscriptionResponse) String() string {
	if len(g._rawJSON) > 0 {
		if value, err := core.StringifyJSON(g._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(g); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", g)
}

// Defines output parameters in a response from the
// [ListSubscriptionEvents](api-endpoint:Subscriptions-ListSubscriptionEvents).
type ListSubscriptionEventsResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The retrieved subscription events.
	SubscriptionEvents []*SubscriptionEvent `json:"subscription_events,omitempty" url:"subscription_events,omitempty"`
	// When the total number of resulting subscription events exceeds the limit of a paged response,
	// the response includes a cursor for you to use in a subsequent request to fetch the next set of events.
	// If the cursor is unset, the response contains the last page of the results.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListSubscriptionEventsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListSubscriptionEventsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListSubscriptionEventsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListSubscriptionEventsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListSubscriptionEventsResponse) String() string {
	if len(l._rawJSON) > 0 {
		if value, err := core.StringifyJSON(l._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(l); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", l)
}

// Defines output parameters in a response from the
// [PauseSubscription](api-endpoint:Subscriptions-PauseSubscription) endpoint.
type PauseSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The subscription to be paused by the scheduled `PAUSE` action.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`
	// The list of a `PAUSE` action and a possible `RESUME` action created by the request.
	Actions []*SubscriptionAction `json:"actions,omitempty" url:"actions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (p *PauseSubscriptionResponse) GetExtraProperties() map[string]interface{} {
	return p.extraProperties
}

func (p *PauseSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler PauseSubscriptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*p = PauseSubscriptionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *p)
	if err != nil {
		return err
	}
	p.extraProperties = extraProperties

	p._rawJSON = json.RawMessage(data)
	return nil
}

func (p *PauseSubscriptionResponse) String() string {
	if len(p._rawJSON) > 0 {
		if value, err := core.StringifyJSON(p._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(p); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", p)
}

// Defines output parameters in a response from the
// [ResumeSubscription](api-endpoint:Subscriptions-ResumeSubscription) endpoint.
type ResumeSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The resumed subscription.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`
	// A list of `RESUME` actions created by the request and scheduled for the subscription.
	Actions []*SubscriptionAction `json:"actions,omitempty" url:"actions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *ResumeSubscriptionResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *ResumeSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ResumeSubscriptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = ResumeSubscriptionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *ResumeSubscriptionResponse) String() string {
	if len(r._rawJSON) > 0 {
		if value, err := core.StringifyJSON(r._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(r); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", r)
}

// Represents a query, consisting of specified query expressions, used to search for subscriptions.
type SearchSubscriptionsQuery struct {
	// A list of query expressions.
	Filter *SearchSubscriptionsFilter `json:"filter,omitempty" url:"filter,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchSubscriptionsQuery) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchSubscriptionsQuery) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchSubscriptionsQuery
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchSubscriptionsQuery(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchSubscriptionsQuery) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Defines output parameters in a response from the
// [SearchSubscriptions](api-endpoint:Subscriptions-SearchSubscriptions) endpoint.
type SearchSubscriptionsResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The subscriptions matching the specified query expressions.
	Subscriptions []*Subscription `json:"subscriptions,omitempty" url:"subscriptions,omitempty"`
	// When the total number of resulting subscription exceeds the limit of a paged response,
	// the response includes a cursor for you to use in a subsequent request to fetch the next set of results.
	// If the cursor is unset, the response contains the last page of the results.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SearchSubscriptionsResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SearchSubscriptionsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SearchSubscriptionsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SearchSubscriptionsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SearchSubscriptionsResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Represents a subscription to a subscription plan by a subscriber.
//
// For an overview of the `Subscription` type, see
// [Subscription object](https://developer.squareup.com/docs/subscriptions-api/overview#subscription-object-overview).
type Subscription struct {
	// The Square-assigned ID of the subscription.
	Id *string `json:"id,omitempty" url:"id,omitempty"`
	// The ID of the location associated with the subscription.
	LocationId *string `json:"location_id,omitempty" url:"location_id,omitempty"`
	// The ID of the subscribed-to [subscription plan](entity:CatalogSubscriptionPlan).
	PlanId *string `json:"plan_id,omitempty" url:"plan_id,omitempty"`
	// The ID of the subscribing [customer](entity:Customer) profile.
	CustomerId *string `json:"customer_id,omitempty" url:"customer_id,omitempty"`
	// The `YYYY-MM-DD`-formatted date (for example, 2013-01-15) to start the subscription.
	StartDate *string `json:"start_date,omitempty" url:"start_date,omitempty"`
	// The `YYYY-MM-DD`-formatted date (for example, 2013-01-15) to cancel the subscription,
	// when the subscription status changes to `CANCELED` and the subscription billing stops.
	//
	// If this field is not set, the subscription ends according its subscription plan.
	//
	// This field cannot be updated, other than being cleared.
	CanceledDate *string `json:"canceled_date,omitempty" url:"canceled_date,omitempty"`
	// The `YYYY-MM-DD`-formatted date up to when the subscriber is invoiced for the
	// subscription.
	//
	// After the invoice is sent for a given billing period,
	// this date will be the last day of the billing period.
	// For example,
	// suppose for the month of May a subscriber gets an invoice
	// (or charged the card) on May 1. For the monthly billing scenario,
	// this date is then set to May 31.
	ChargedThroughDate *string `json:"charged_through_date,omitempty" url:"charged_through_date,omitempty"`
	// The current status of the subscription.
	// See [SubscriptionStatus](#type-subscriptionstatus) for possible values
	Status *SubscriptionStatus `json:"status,omitempty" url:"status,omitempty"`
	// The tax amount applied when billing the subscription. The
	// percentage is expressed in decimal form, using a `'.'` as the decimal
	// separator and without a `'%'` sign. For example, a value of `7.5`
	// corresponds to 7.5%.
	TaxPercentage *string `json:"tax_percentage,omitempty" url:"tax_percentage,omitempty"`
	// The IDs of the [invoices](entity:Invoice) created for the
	// subscription, listed in order when the invoices were created
	// (newest invoices appear first).
	InvoiceIds []string `json:"invoice_ids,omitempty" url:"invoice_ids,omitempty"`
	// A custom price to apply for the subscription. If specified,
	// it overrides the price configured by the subscription plan.
	PriceOverrideMoney *Money `json:"price_override_money,omitempty" url:"price_override_money,omitempty"`
	// The version of the object. When updating an object, the version
	// supplied must match the version in the database, otherwise the write will
	// be rejected as conflicting.
	Version *int64 `json:"version,omitempty" url:"version,omitempty"`
	// The timestamp when the subscription was created, in RFC 3339 format.
	CreatedAt *string `json:"created_at,omitempty" url:"created_at,omitempty"`
	// The ID of the [subscriber's](entity:Customer) [card](entity:Card)
	// used to charge for the subscription.
	CardId *string `json:"card_id,omitempty" url:"card_id,omitempty"`
	// The `YYYY-MM-DD`-formatted date (for example, 2013-01-15) up to when the subscriber is invoiced for the
	// subscription.
	//
	// After the invoice is sent for a given billing period,
	// this date will be the last day of the billing period.
	// For example,
	// suppose for the month of May a subscriber gets an invoice
	// (or charged the card) on May 1. For the monthly billing scenario,
	// this date is then set to May 31.
	PaidUntilDate *string `json:"paid_until_date,omitempty" url:"paid_until_date,omitempty"`
	// Timezone that will be used in date calculations for the subscription.
	// Defaults to the timezone of the location based on `location_id`.
	// Format: the IANA Timezone Database identifier for the location timezone (for example, `America/Los_Angeles`).
	Timezone *string `json:"timezone,omitempty" url:"timezone,omitempty"`
	// The origination details of the subscription.
	Source *SubscriptionSource `json:"source,omitempty" url:"source,omitempty"`
	// The list of scheduled actions on this subscription. It is set only in the response from
	// [RetrieveSubscription](api-endpoint:Subscriptions-RetrieveSubscription) with the query parameter
	// of `include=actions` or from
	// [SearchSubscriptions](api-endpoint:Subscriptions-SearchSubscriptions) with the input parameter
	// of `include:["actions"]`.
	Actions []*SubscriptionAction `json:"actions,omitempty" url:"actions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *Subscription) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *Subscription) UnmarshalJSON(data []byte) error {
	type unmarshaler Subscription
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = Subscription(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *Subscription) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// The origination details of the subscription.
type SubscriptionSource struct {
	// The name used to identify the place (physical or digital) that
	// a subscription originates. If unset, the name defaults to the name
	// of the application that created the subscription.
	Name *string `json:"name,omitempty" url:"name,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SubscriptionSource) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SubscriptionSource) UnmarshalJSON(data []byte) error {
	type unmarshaler SubscriptionSource
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SubscriptionSource(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SubscriptionSource) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Defines output parameters in a response of the
// [SwapPlan](api-endpoint:Subscriptions-SwapPlan) endpoint.
type SwapPlanResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The subscription with the updated subscription plan.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`
	// A list of a `SWAP_PLAN` action created by the request.
	Actions []*SubscriptionAction `json:"actions,omitempty" url:"actions,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (s *SwapPlanResponse) GetExtraProperties() map[string]interface{} {
	return s.extraProperties
}

func (s *SwapPlanResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler SwapPlanResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*s = SwapPlanResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *s)
	if err != nil {
		return err
	}
	s.extraProperties = extraProperties

	s._rawJSON = json.RawMessage(data)
	return nil
}

func (s *SwapPlanResponse) String() string {
	if len(s._rawJSON) > 0 {
		if value, err := core.StringifyJSON(s._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(s); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", s)
}

// Defines output parameters in a response from the
// [UpdateSubscription](api-endpoint:Subscriptions-UpdateSubscription) endpoint.
type UpdateSubscriptionResponse struct {
	// Errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The updated subscription.
	Subscription *Subscription `json:"subscription,omitempty" url:"subscription,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (u *UpdateSubscriptionResponse) GetExtraProperties() map[string]interface{} {
	return u.extraProperties
}

func (u *UpdateSubscriptionResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler UpdateSubscriptionResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*u = UpdateSubscriptionResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *u)
	if err != nil {
		return err
	}
	u.extraProperties = extraProperties

	u._rawJSON = json.RawMessage(data)
	return nil
}

func (u *UpdateSubscriptionResponse) String() string {
	if len(u._rawJSON) > 0 {
		if value, err := core.StringifyJSON(u._rawJSON); err == nil {
			return value
		}
	}
	if value, err := core.StringifyJSON(u); err == nil {
		return value
	}
	return fmt.Sprintf("%#v", u)
}

type UpdateSubscriptionRequest struct {
	// The subscription object containing the current version, and fields to update.
	// Unset fields will be left at their current server values, and JSON `null` values will
	// be treated as a request to clear the relevant data.
	Subscription *Subscription `json:"subscription,omitempty" url:"-"`
}
