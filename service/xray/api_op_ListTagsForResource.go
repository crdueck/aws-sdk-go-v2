// Code generated by smithy-go-codegen DO NOT EDIT.

package xray

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/xray/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of tags that are applied to the specified Amazon Web Services
// X-Ray group or sampling rule.
func (c *Client) ListTagsForResource(ctx context.Context, params *ListTagsForResourceInput, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListTagsForResource", params, optFns, c.addOperationListTagsForResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListTagsForResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListTagsForResourceInput struct {

	// The Amazon Resource Number (ARN) of an X-Ray group or sampling rule.
	//
	// This member is required.
	ResourceARN *string

	// A pagination token. If multiple pages of results are returned, use the NextToken
	// value returned with the current page of results as the value of this parameter
	// to get the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListTagsForResourceOutput struct {

	// A pagination token. If multiple pages of results are returned, use the NextToken
	// value returned with the current page of results to get the next page of results.
	NextToken *string

	// A list of tags, as key and value pairs, that is associated with the specified
	// X-Ray group or sampling rule.
	Tags []types.Tag

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListTagsForResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListTagsForResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListTagsForResource{}, middleware.After)
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
	if err = addOpListTagsForResourceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListTagsForResource(options.Region), middleware.Before); err != nil {
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

// ListTagsForResourceAPIClient is a client that implements the
// ListTagsForResource operation.
type ListTagsForResourceAPIClient interface {
	ListTagsForResource(context.Context, *ListTagsForResourceInput, ...func(*Options)) (*ListTagsForResourceOutput, error)
}

var _ ListTagsForResourceAPIClient = (*Client)(nil)

// ListTagsForResourcePaginatorOptions is the paginator options for
// ListTagsForResource
type ListTagsForResourcePaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListTagsForResourcePaginator is a paginator for ListTagsForResource
type ListTagsForResourcePaginator struct {
	options   ListTagsForResourcePaginatorOptions
	client    ListTagsForResourceAPIClient
	params    *ListTagsForResourceInput
	nextToken *string
	firstPage bool
}

// NewListTagsForResourcePaginator returns a new ListTagsForResourcePaginator
func NewListTagsForResourcePaginator(client ListTagsForResourceAPIClient, params *ListTagsForResourceInput, optFns ...func(*ListTagsForResourcePaginatorOptions)) *ListTagsForResourcePaginator {
	if params == nil {
		params = &ListTagsForResourceInput{}
	}

	options := ListTagsForResourcePaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListTagsForResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListTagsForResourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListTagsForResource page.
func (p *ListTagsForResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListTagsForResourceOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListTagsForResource(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListTagsForResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "xray",
		OperationName: "ListTagsForResource",
	}
}
