// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/mixed-file-directory/fern/core"
)

type RefundPaymentRequest struct {
	//	A unique string that identifies this `RefundPayment` request. The key can be any valid string
	//
	// but must be unique for every `RefundPayment` request.
	//
	// Keys are limited to a max of 45 characters - however, the number of allowed characters might be
	// less than 45, if multi-byte characters are used.
	//
	// For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key" url:"-"`
	// The amount of money to refund.
	//
	// This amount cannot be more than the `total_money` value of the payment minus the total
	// amount of all previously completed refunds for this payment.
	//
	// This amount must be specified in the smallest denomination of the applicable currency
	// (for example, US dollar amounts are specified in cents). For more information, see
	// [Working with Monetary Amounts](https://developer.squareup.com/docs/build-basics/working-with-monetary-amounts).
	//
	// The currency code must match the currency associated with the business
	// that is charging the card.
	AmountMoney *Money `json:"amount_money,omitempty" url:"-"`
	// The amount of money the developer contributes to help cover the refunded amount.
	// This amount is specified in the smallest denomination of the applicable currency (for example,
	// US dollar amounts are specified in cents).
	//
	// The value cannot be more than the `amount_money`.
	//
	// You can specify this parameter in a refund request only if the same parameter was also included
	// when taking the payment. This is part of the application fee scenario the API supports. For more
	// information, see [Take Payments and Collect Fees](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees).
	//
	// To set this field, `PAYMENTS_WRITE_ADDITIONAL_RECIPIENTS` OAuth permission is required.
	// For more information, see [Permissions](https://developer.squareup.com/docs/payments-api/take-payments-and-collect-fees#permissions).
	AppFeeMoney *Money `json:"app_fee_money,omitempty" url:"-"`
	// The unique ID of the payment being refunded. Must be provided and non-empty.
	PaymentId *string `json:"payment_id,omitempty" url:"-"`
	// A description of the reason for the refund.
	Reason *string `json:"reason,omitempty" url:"-"`
	//	Used for optimistic concurrency. This opaque token identifies the current `Payment`
	//
	// version that the caller expects. If the server has a different version of the Payment,
	// the update fails and a response with a VERSION_MISMATCH error is returned.
	// If the versions match, or the field is not provided, the refund proceeds as normal.
	PaymentVersionToken *string `json:"payment_version_token,omitempty" url:"-"`
	// An optional [TeamMember](entity:TeamMember) ID to associate with this refund.
	TeamMemberId *string `json:"team_member_id,omitempty" url:"-"`
}

type RefundsListRequest struct {
	// The timestamp for the beginning of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time minus one year.
	BeginTime *string `json:"-" url:"begin_time,omitempty"`
	// The timestamp for the end of the requested reporting period, in RFC 3339 format.
	//
	// Default: The current time.
	EndTime *string `json:"-" url:"end_time,omitempty"`
	// The order in which results are listed:
	//
	// - `ASC` - Oldest to newest.
	// - `DESC` - Newest to oldest (default).
	SortOrder *string `json:"-" url:"sort_order,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint.
	// Provide this cursor to retrieve the next set of results for the original query.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor *string `json:"-" url:"cursor,omitempty"`
	// Limit results to the location supplied. By default, results are returned
	// for all locations associated with the seller.
	LocationId *string `json:"-" url:"location_id,omitempty"`
	// If provided, only refunds with the given status are returned.
	// For a list of refund status values, see [PaymentRefund](entity:PaymentRefund).
	//
	// Default: If omitted, refunds are returned regardless of their status.
	Status *string `json:"-" url:"status,omitempty"`
	// If provided, only returns refunds whose payments have the indicated source type.
	// Current values include `CARD`, `BANK_ACCOUNT`, `WALLET`, `CASH`, and `EXTERNAL`.
	// For information about these payment source types, see
	// [Take Payments](https://developer.squareup.com/docs/payments-api/take-payments).
	//
	// Default: If omitted, refunds are returned regardless of the source type.
	SourceType *string `json:"-" url:"source_type,omitempty"`
	// The maximum number of results to be returned in a single page.
	//
	// It is possible to receive fewer results than the specified limit on a given page.
	//
	// If the supplied value is greater than 100, no more than 100 results are returned.
	//
	// Default: 100
	Limit *int `json:"-" url:"limit,omitempty"`
}

// Defines the response returned by [GetRefund](api-endpoint:Refunds-GetPaymentRefund).
//
// Note: If there are errors processing the request, the refund field might not be
// present or it might be present in a FAILED state.
type GetPaymentRefundResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The requested `PaymentRefund`.
	Refund *PaymentRefund `json:"refund,omitempty" url:"refund,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (g *GetPaymentRefundResponse) GetExtraProperties() map[string]interface{} {
	return g.extraProperties
}

func (g *GetPaymentRefundResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler GetPaymentRefundResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*g = GetPaymentRefundResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *g)
	if err != nil {
		return err
	}
	g.extraProperties = extraProperties

	g._rawJSON = json.RawMessage(data)
	return nil
}

func (g *GetPaymentRefundResponse) String() string {
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

// Defines the response returned by [ListPaymentRefunds](api-endpoint:Refunds-ListPaymentRefunds).
//
// Either `errors` or `refunds` is present in a given response (never both).
type ListPaymentRefundsResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The list of requested refunds.
	Refunds []*PaymentRefund `json:"refunds,omitempty" url:"refunds,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty,
	// this is the final response.
	//
	// For more information, see [Pagination](https://developer.squareup.com/docs/basics/api101/pagination).
	Cursor *string `json:"cursor,omitempty" url:"cursor,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (l *ListPaymentRefundsResponse) GetExtraProperties() map[string]interface{} {
	return l.extraProperties
}

func (l *ListPaymentRefundsResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler ListPaymentRefundsResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*l = ListPaymentRefundsResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *l)
	if err != nil {
		return err
	}
	l.extraProperties = extraProperties

	l._rawJSON = json.RawMessage(data)
	return nil
}

func (l *ListPaymentRefundsResponse) String() string {
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

// Defines the response returned by
// [RefundPayment](api-endpoint:Refunds-RefundPayment).
//
// If there are errors processing the request, the `refund` field might not be
// present, or it might be present with a status of `FAILED`.
type RefundPaymentResponse struct {
	// Information about errors encountered during the request.
	Errors []*Error `json:"errors,omitempty" url:"errors,omitempty"`
	// The successfully created `PaymentRefund`.
	Refund *PaymentRefund `json:"refund,omitempty" url:"refund,omitempty"`

	extraProperties map[string]interface{}
	_rawJSON        json.RawMessage
}

func (r *RefundPaymentResponse) GetExtraProperties() map[string]interface{} {
	return r.extraProperties
}

func (r *RefundPaymentResponse) UnmarshalJSON(data []byte) error {
	type unmarshaler RefundPaymentResponse
	var value unmarshaler
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*r = RefundPaymentResponse(value)

	extraProperties, err := core.ExtractExtraProperties(data, *r)
	if err != nil {
		return err
	}
	r.extraProperties = extraProperties

	r._rawJSON = json.RawMessage(data)
	return nil
}

func (r *RefundPaymentResponse) String() string {
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
