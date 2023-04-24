// Code generated by smithy-go-codegen DO NOT EDIT.

package comprehend

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/comprehend/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets a list of summaries of the document classifiers that you have created
func (c *Client) ListDocumentClassifierSummaries(ctx context.Context, params *ListDocumentClassifierSummariesInput, optFns ...func(*Options)) (*ListDocumentClassifierSummariesOutput, error) {
	if params == nil {
		params = &ListDocumentClassifierSummariesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListDocumentClassifierSummaries", params, optFns, c.addOperationListDocumentClassifierSummariesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListDocumentClassifierSummariesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListDocumentClassifierSummariesInput struct {

	// The maximum number of results to return on each page. The default is 100.
	MaxResults *int32

	// Identifies the next page of results to return.
	NextToken *string

	noSmithyDocumentSerde
}

type ListDocumentClassifierSummariesOutput struct {

	// The list of summaries of document classifiers.
	DocumentClassifierSummariesList []types.DocumentClassifierSummary

	// Identifies the next page of results to return.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListDocumentClassifierSummariesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListDocumentClassifierSummaries{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListDocumentClassifierSummaries{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListDocumentClassifierSummaries(options.Region), middleware.Before); err != nil {
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

// ListDocumentClassifierSummariesAPIClient is a client that implements the
// ListDocumentClassifierSummaries operation.
type ListDocumentClassifierSummariesAPIClient interface {
	ListDocumentClassifierSummaries(context.Context, *ListDocumentClassifierSummariesInput, ...func(*Options)) (*ListDocumentClassifierSummariesOutput, error)
}

var _ ListDocumentClassifierSummariesAPIClient = (*Client)(nil)

// ListDocumentClassifierSummariesPaginatorOptions is the paginator options for
// ListDocumentClassifierSummaries
type ListDocumentClassifierSummariesPaginatorOptions struct {
	// The maximum number of results to return on each page. The default is 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListDocumentClassifierSummariesPaginator is a paginator for
// ListDocumentClassifierSummaries
type ListDocumentClassifierSummariesPaginator struct {
	options   ListDocumentClassifierSummariesPaginatorOptions
	client    ListDocumentClassifierSummariesAPIClient
	params    *ListDocumentClassifierSummariesInput
	nextToken *string
	firstPage bool
}

// NewListDocumentClassifierSummariesPaginator returns a new
// ListDocumentClassifierSummariesPaginator
func NewListDocumentClassifierSummariesPaginator(client ListDocumentClassifierSummariesAPIClient, params *ListDocumentClassifierSummariesInput, optFns ...func(*ListDocumentClassifierSummariesPaginatorOptions)) *ListDocumentClassifierSummariesPaginator {
	if params == nil {
		params = &ListDocumentClassifierSummariesInput{}
	}

	options := ListDocumentClassifierSummariesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListDocumentClassifierSummariesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListDocumentClassifierSummariesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListDocumentClassifierSummaries page.
func (p *ListDocumentClassifierSummariesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListDocumentClassifierSummariesOutput, error) {
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

	result, err := p.client.ListDocumentClassifierSummaries(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListDocumentClassifierSummaries(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "comprehend",
		OperationName: "ListDocumentClassifierSummaries",
	}
}
