// Code generated by smithy-go-codegen DO NOT EDIT.

package mediaconnect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/mediaconnect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Displays a list of instances associated with the AWS account. This request
// returns a paginated result. You can use the filterArn property to display only
// the instances associated with the selected Gateway Amazon Resource Name (ARN).
func (c *Client) ListGatewayInstances(ctx context.Context, params *ListGatewayInstancesInput, optFns ...func(*Options)) (*ListGatewayInstancesOutput, error) {
	if params == nil {
		params = &ListGatewayInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListGatewayInstances", params, optFns, c.addOperationListGatewayInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListGatewayInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListGatewayInstancesInput struct {

	// Filter the list results to display only the instances associated with the
	// selected Gateway Amazon Resource Name (ARN).
	FilterArn *string

	// The maximum number of results to return per API request. For example, you
	// submit a ListInstances request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	MaxResults int32

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListInstances request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListInstances request a second
	// time and specify the NextToken value.
	NextToken *string

	noSmithyDocumentSerde
}

type ListGatewayInstancesOutput struct {

	// A list of instance summaries.
	Instances []types.ListedGatewayInstance

	// The token that identifies which batch of results that you want to see. For
	// example, you submit a ListInstances request with MaxResults set at 5. The
	// service returns the first batch of results (up to 5) and a NextToken value. To
	// see the next batch of results, you can submit the ListInstances request a second
	// time and specify the NextToken value.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListGatewayInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListGatewayInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListGatewayInstances{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListGatewayInstances(options.Region), middleware.Before); err != nil {
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

// ListGatewayInstancesAPIClient is a client that implements the
// ListGatewayInstances operation.
type ListGatewayInstancesAPIClient interface {
	ListGatewayInstances(context.Context, *ListGatewayInstancesInput, ...func(*Options)) (*ListGatewayInstancesOutput, error)
}

var _ ListGatewayInstancesAPIClient = (*Client)(nil)

// ListGatewayInstancesPaginatorOptions is the paginator options for
// ListGatewayInstances
type ListGatewayInstancesPaginatorOptions struct {
	// The maximum number of results to return per API request. For example, you
	// submit a ListInstances request with MaxResults set at 5. Although 20 items match
	// your request, the service returns no more than the first 5 items. (The service
	// also returns a NextToken value that you can use to fetch the next batch of
	// results.) The service might return fewer results than the MaxResults value. If
	// MaxResults is not included in the request, the service defaults to pagination
	// with a maximum of 10 results per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListGatewayInstancesPaginator is a paginator for ListGatewayInstances
type ListGatewayInstancesPaginator struct {
	options   ListGatewayInstancesPaginatorOptions
	client    ListGatewayInstancesAPIClient
	params    *ListGatewayInstancesInput
	nextToken *string
	firstPage bool
}

// NewListGatewayInstancesPaginator returns a new ListGatewayInstancesPaginator
func NewListGatewayInstancesPaginator(client ListGatewayInstancesAPIClient, params *ListGatewayInstancesInput, optFns ...func(*ListGatewayInstancesPaginatorOptions)) *ListGatewayInstancesPaginator {
	if params == nil {
		params = &ListGatewayInstancesInput{}
	}

	options := ListGatewayInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListGatewayInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListGatewayInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListGatewayInstances page.
func (p *ListGatewayInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListGatewayInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListGatewayInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListGatewayInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "mediaconnect",
		OperationName: "ListGatewayInstances",
	}
}
