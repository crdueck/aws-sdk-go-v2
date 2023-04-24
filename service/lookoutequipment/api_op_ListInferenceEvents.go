// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutequipment

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutequipment/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Lists all inference events that have been found for the specified inference
// scheduler.
func (c *Client) ListInferenceEvents(ctx context.Context, params *ListInferenceEventsInput, optFns ...func(*Options)) (*ListInferenceEventsOutput, error) {
	if params == nil {
		params = &ListInferenceEventsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListInferenceEvents", params, optFns, c.addOperationListInferenceEventsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListInferenceEventsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListInferenceEventsInput struct {

	// The name of the inference scheduler for the inference events listed.
	//
	// This member is required.
	InferenceSchedulerName *string

	// Returns all the inference events with an end start time equal to or greater
	// than less than the end time given
	//
	// This member is required.
	IntervalEndTime *time.Time

	// Lookout for Equipment will return all the inference events with an end time
	// equal to or greater than the start time given.
	//
	// This member is required.
	IntervalStartTime *time.Time

	// Specifies the maximum number of inference events to list.
	MaxResults *int32

	// An opaque pagination token indicating where to continue the listing of
	// inference events.
	NextToken *string

	noSmithyDocumentSerde
}

type ListInferenceEventsOutput struct {

	// Provides an array of information about the individual inference events returned
	// from the ListInferenceEvents operation, including scheduler used, event start
	// time, event end time, diagnostics, and so on.
	InferenceEventSummaries []types.InferenceEventSummary

	// An opaque pagination token indicating where to continue the listing of
	// inference executions.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListInferenceEventsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpListInferenceEvents{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpListInferenceEvents{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListInferenceEventsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListInferenceEvents(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

// ListInferenceEventsAPIClient is a client that implements the
// ListInferenceEvents operation.
type ListInferenceEventsAPIClient interface {
	ListInferenceEvents(context.Context, *ListInferenceEventsInput, ...func(*Options)) (*ListInferenceEventsOutput, error)
}

var _ ListInferenceEventsAPIClient = (*Client)(nil)

// ListInferenceEventsPaginatorOptions is the paginator options for
// ListInferenceEvents
type ListInferenceEventsPaginatorOptions struct {
	// Specifies the maximum number of inference events to list.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListInferenceEventsPaginator is a paginator for ListInferenceEvents
type ListInferenceEventsPaginator struct {
	options   ListInferenceEventsPaginatorOptions
	client    ListInferenceEventsAPIClient
	params    *ListInferenceEventsInput
	nextToken *string
	firstPage bool
}

// NewListInferenceEventsPaginator returns a new ListInferenceEventsPaginator
func NewListInferenceEventsPaginator(client ListInferenceEventsAPIClient, params *ListInferenceEventsInput, optFns ...func(*ListInferenceEventsPaginatorOptions)) *ListInferenceEventsPaginator {
	if params == nil {
		params = &ListInferenceEventsInput{}
	}

	options := ListInferenceEventsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListInferenceEventsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListInferenceEventsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListInferenceEvents page.
func (p *ListInferenceEventsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListInferenceEventsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxResults = limit

	result, err := p.client.ListInferenceEvents(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.NextToken

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opListInferenceEvents(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutequipment",
		OperationName: "ListInferenceEvents",
	}
}
