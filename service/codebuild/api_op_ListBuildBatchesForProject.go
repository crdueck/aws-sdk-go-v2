// Code generated by smithy-go-codegen DO NOT EDIT.

package codebuild

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codebuild/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the identifiers of the build batches for a specific project.
func (c *Client) ListBuildBatchesForProject(ctx context.Context, params *ListBuildBatchesForProjectInput, optFns ...func(*Options)) (*ListBuildBatchesForProjectOutput, error) {
	if params == nil {
		params = &ListBuildBatchesForProjectInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBuildBatchesForProject", params, optFns, c.addOperationListBuildBatchesForProjectMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBuildBatchesForProjectOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBuildBatchesForProjectInput struct {

	// A BuildBatchFilter object that specifies the filters for the search.
	Filter *types.BuildBatchFilter

	// The maximum number of results to return.
	MaxResults *int32

	// The nextToken value returned from a previous call to ListBuildBatchesForProject
	// . This specifies the next item to return. To return the beginning of the list,
	// exclude this parameter.
	NextToken *string

	// The name of the project.
	ProjectName *string

	// Specifies the sort order of the returned items. Valid values include:
	//   - ASCENDING : List the batch build identifiers in ascending order by
	//   identifier.
	//   - DESCENDING : List the batch build identifiers in descending order by
	//   identifier.
	SortOrder types.SortOrderType

	noSmithyDocumentSerde
}

type ListBuildBatchesForProjectOutput struct {

	// An array of strings that contains the batch build identifiers.
	Ids []string

	// If there are more items to return, this contains a token that is passed to a
	// subsequent call to ListBuildBatchesForProject to retrieve the next set of items.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBuildBatchesForProjectMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListBuildBatchesForProject{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListBuildBatchesForProject{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBuildBatchesForProject(options.Region), middleware.Before); err != nil {
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

// ListBuildBatchesForProjectAPIClient is a client that implements the
// ListBuildBatchesForProject operation.
type ListBuildBatchesForProjectAPIClient interface {
	ListBuildBatchesForProject(context.Context, *ListBuildBatchesForProjectInput, ...func(*Options)) (*ListBuildBatchesForProjectOutput, error)
}

var _ ListBuildBatchesForProjectAPIClient = (*Client)(nil)

// ListBuildBatchesForProjectPaginatorOptions is the paginator options for
// ListBuildBatchesForProject
type ListBuildBatchesForProjectPaginatorOptions struct {
	// The maximum number of results to return.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBuildBatchesForProjectPaginator is a paginator for
// ListBuildBatchesForProject
type ListBuildBatchesForProjectPaginator struct {
	options   ListBuildBatchesForProjectPaginatorOptions
	client    ListBuildBatchesForProjectAPIClient
	params    *ListBuildBatchesForProjectInput
	nextToken *string
	firstPage bool
}

// NewListBuildBatchesForProjectPaginator returns a new
// ListBuildBatchesForProjectPaginator
func NewListBuildBatchesForProjectPaginator(client ListBuildBatchesForProjectAPIClient, params *ListBuildBatchesForProjectInput, optFns ...func(*ListBuildBatchesForProjectPaginatorOptions)) *ListBuildBatchesForProjectPaginator {
	if params == nil {
		params = &ListBuildBatchesForProjectInput{}
	}

	options := ListBuildBatchesForProjectPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBuildBatchesForProjectPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBuildBatchesForProjectPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBuildBatchesForProject page.
func (p *ListBuildBatchesForProjectPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBuildBatchesForProjectOutput, error) {
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

	result, err := p.client.ListBuildBatchesForProject(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBuildBatchesForProject(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codebuild",
		OperationName: "ListBuildBatchesForProject",
	}
}
