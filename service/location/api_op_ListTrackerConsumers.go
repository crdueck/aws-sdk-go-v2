// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists geofence collections currently associated to the given tracker resource.
func (c *Client) ListTrackerConsumers(ctx context.Context, params *ListTrackerConsumersInput, optFns ...func(*Options)) (*ListTrackerConsumersOutput, error) {
	if params == nil {
		params = &ListTrackerConsumersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTrackerConsumers", params, optFns, c.addOperationListTrackerConsumersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTrackerConsumersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTrackerConsumersInput struct {

	// The tracker resource whose associated geofence collections you want to list.
	//
	// This member is required.
	TrackerName *string

	// An optional limit for the number of resources returned in a single call.
	// Default value: 100
	MaxResults *int32

	// The pagination token specifying which page of results to return in the
	// response. If no token is provided, the default page is the first page. Default
	// value: null
	NextToken *string

	noSmithyDocumentSerde
}

type ListTrackerConsumersOutput struct {

	// Contains the list of geofence collection ARNs associated to the tracker
	// resource.
	//
	// This member is required.
	ConsumerArns []string

	// A pagination token indicating there are additional pages available. You can use
	// the token in a following request to fetch the next set of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTrackerConsumersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTrackerConsumers{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTrackerConsumers{}, middleware.After)
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
	if err = addEndpointPrefix_opListTrackerConsumersMiddleware(stack); err != nil {
		return err
	}
	if err = addOpListTrackerConsumersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTrackerConsumers(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opListTrackerConsumersMiddleware struct {
}

func (*endpointPrefix_opListTrackerConsumersMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opListTrackerConsumersMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "tracking." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opListTrackerConsumersMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opListTrackerConsumersMiddleware{}, `OperationSerializer`, middleware.After)
}

// ListTrackerConsumersAPIClient is a client that implements the
// ListTrackerConsumers operation.
type ListTrackerConsumersAPIClient interface {
	ListTrackerConsumers(context.Context, *ListTrackerConsumersInput, ...func(*Options)) (*ListTrackerConsumersOutput, error)
}

var _ ListTrackerConsumersAPIClient = (*Client)(nil)

// ListTrackerConsumersPaginatorOptions is the paginator options for
// ListTrackerConsumers
type ListTrackerConsumersPaginatorOptions struct {
	// An optional limit for the number of resources returned in a single call.
	// Default value: 100
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTrackerConsumersPaginator is a paginator for ListTrackerConsumers
type ListTrackerConsumersPaginator struct {
	options   ListTrackerConsumersPaginatorOptions
	client    ListTrackerConsumersAPIClient
	params    *ListTrackerConsumersInput
	nextToken *string
	firstPage bool
}

// NewListTrackerConsumersPaginator returns a new ListTrackerConsumersPaginator
func NewListTrackerConsumersPaginator(client ListTrackerConsumersAPIClient, params *ListTrackerConsumersInput, optFns ...func(*ListTrackerConsumersPaginatorOptions)) *ListTrackerConsumersPaginator {
	if params == nil {
		params = &ListTrackerConsumersInput{}
	}

	options := ListTrackerConsumersPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTrackerConsumersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTrackerConsumersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTrackerConsumers page.
func (p *ListTrackerConsumersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTrackerConsumersOutput, error) {
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

	result, err := p.client.ListTrackerConsumers(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTrackerConsumers(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "ListTrackerConsumers",
	}
}
