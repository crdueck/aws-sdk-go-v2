// Code generated by smithy-go-codegen DO NOT EDIT.

package panorama

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/panorama/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of application node instances.
func (c *Client) ListApplicationInstanceNodeInstances(ctx context.Context, params *ListApplicationInstanceNodeInstancesInput, optFns ...func(*Options)) (*ListApplicationInstanceNodeInstancesOutput, error) {
	if params == nil {
		params = &ListApplicationInstanceNodeInstancesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListApplicationInstanceNodeInstances", params, optFns, c.addOperationListApplicationInstanceNodeInstancesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListApplicationInstanceNodeInstancesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListApplicationInstanceNodeInstancesInput struct {

	// The node instances' application instance ID.
	//
	// This member is required.
	ApplicationInstanceId *string

	// The maximum number of node instances to return in one page of results.
	MaxResults int32

	// Specify the pagination token from a previous request to retrieve the next page
	// of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListApplicationInstanceNodeInstancesOutput struct {

	// A pagination token that's included if more results are available.
	NextToken *string

	// A list of node instances.
	NodeInstances []types.NodeInstance

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListApplicationInstanceNodeInstancesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListApplicationInstanceNodeInstances{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListApplicationInstanceNodeInstances{}, middleware.After)
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
	if err = addOpListApplicationInstanceNodeInstancesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListApplicationInstanceNodeInstances(options.Region), middleware.Before); err != nil {
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

// ListApplicationInstanceNodeInstancesAPIClient is a client that implements the
// ListApplicationInstanceNodeInstances operation.
type ListApplicationInstanceNodeInstancesAPIClient interface {
	ListApplicationInstanceNodeInstances(context.Context, *ListApplicationInstanceNodeInstancesInput, ...func(*Options)) (*ListApplicationInstanceNodeInstancesOutput, error)
}

var _ ListApplicationInstanceNodeInstancesAPIClient = (*Client)(nil)

// ListApplicationInstanceNodeInstancesPaginatorOptions is the paginator options
// for ListApplicationInstanceNodeInstances
type ListApplicationInstanceNodeInstancesPaginatorOptions struct {
	// The maximum number of node instances to return in one page of results.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListApplicationInstanceNodeInstancesPaginator is a paginator for
// ListApplicationInstanceNodeInstances
type ListApplicationInstanceNodeInstancesPaginator struct {
	options   ListApplicationInstanceNodeInstancesPaginatorOptions
	client    ListApplicationInstanceNodeInstancesAPIClient
	params    *ListApplicationInstanceNodeInstancesInput
	nextToken *string
	firstPage bool
}

// NewListApplicationInstanceNodeInstancesPaginator returns a new
// ListApplicationInstanceNodeInstancesPaginator
func NewListApplicationInstanceNodeInstancesPaginator(client ListApplicationInstanceNodeInstancesAPIClient, params *ListApplicationInstanceNodeInstancesInput, optFns ...func(*ListApplicationInstanceNodeInstancesPaginatorOptions)) *ListApplicationInstanceNodeInstancesPaginator {
	if params == nil {
		params = &ListApplicationInstanceNodeInstancesInput{}
	}

	options := ListApplicationInstanceNodeInstancesPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListApplicationInstanceNodeInstancesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListApplicationInstanceNodeInstancesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListApplicationInstanceNodeInstances page.
func (p *ListApplicationInstanceNodeInstancesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListApplicationInstanceNodeInstancesOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.ListApplicationInstanceNodeInstances(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListApplicationInstanceNodeInstances(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "panorama",
		OperationName: "ListApplicationInstanceNodeInstances",
	}
}
