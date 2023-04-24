// Code generated by smithy-go-codegen DO NOT EDIT.

package elasticache

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/elasticache/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of cache parameter group descriptions. If a cache parameter
// group name is specified, the list contains only the descriptions for that group.
func (c *Client) DescribeCacheParameterGroups(ctx context.Context, params *DescribeCacheParameterGroupsInput, optFns ...func(*Options)) (*DescribeCacheParameterGroupsOutput, error) {
	if params == nil {
		params = &DescribeCacheParameterGroupsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeCacheParameterGroups", params, optFns, c.addOperationDescribeCacheParameterGroupsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeCacheParameterGroupsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a DescribeCacheParameterGroups operation.
type DescribeCacheParameterGroupsInput struct {

	// The name of a specific cache parameter group to return details for.
	CacheParameterGroupName *string

	// An optional marker returned from a prior request. Use this marker for
	// pagination of results from this operation. If this parameter is specified, the
	// response includes only records beyond the marker, up to the value specified by
	// MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

// Represents the output of a DescribeCacheParameterGroups operation.
type DescribeCacheParameterGroupsOutput struct {

	// A list of cache parameter groups. Each element in the list contains detailed
	// information about one cache parameter group.
	CacheParameterGroups []types.CacheParameterGroup

	// Provides an identifier to allow retrieval of paginated results.
	Marker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeCacheParameterGroupsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeCacheParameterGroups{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeCacheParameterGroups{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeCacheParameterGroups(options.Region), middleware.Before); err != nil {
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

// DescribeCacheParameterGroupsAPIClient is a client that implements the
// DescribeCacheParameterGroups operation.
type DescribeCacheParameterGroupsAPIClient interface {
	DescribeCacheParameterGroups(context.Context, *DescribeCacheParameterGroupsInput, ...func(*Options)) (*DescribeCacheParameterGroupsOutput, error)
}

var _ DescribeCacheParameterGroupsAPIClient = (*Client)(nil)

// DescribeCacheParameterGroupsPaginatorOptions is the paginator options for
// DescribeCacheParameterGroups
type DescribeCacheParameterGroupsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a marker is included in the response so
	// that the remaining results can be retrieved. Default: 100 Constraints: minimum
	// 20; maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeCacheParameterGroupsPaginator is a paginator for
// DescribeCacheParameterGroups
type DescribeCacheParameterGroupsPaginator struct {
	options   DescribeCacheParameterGroupsPaginatorOptions
	client    DescribeCacheParameterGroupsAPIClient
	params    *DescribeCacheParameterGroupsInput
	nextToken *string
	firstPage bool
}

// NewDescribeCacheParameterGroupsPaginator returns a new
// DescribeCacheParameterGroupsPaginator
func NewDescribeCacheParameterGroupsPaginator(client DescribeCacheParameterGroupsAPIClient, params *DescribeCacheParameterGroupsInput, optFns ...func(*DescribeCacheParameterGroupsPaginatorOptions)) *DescribeCacheParameterGroupsPaginator {
	if params == nil {
		params = &DescribeCacheParameterGroupsInput{}
	}

	options := DescribeCacheParameterGroupsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeCacheParameterGroupsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeCacheParameterGroupsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeCacheParameterGroups page.
func (p *DescribeCacheParameterGroupsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeCacheParameterGroupsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.Marker = p.nextToken

	var limit *int32
	if p.options.Limit > 0 {
		limit = &p.options.Limit
	}
	params.MaxRecords = limit

	result, err := p.client.DescribeCacheParameterGroups(ctx, &params, optFns...)
	if err != nil {
		return nil, err
	}
	p.firstPage = false

	prevToken := p.nextToken
	p.nextToken = result.Marker

	if p.options.StopOnDuplicateToken &&
		prevToken != nil &&
		p.nextToken != nil &&
		*prevToken == *p.nextToken {
		p.nextToken = nil
	}

	return result, nil
}

func newServiceMetadataMiddleware_opDescribeCacheParameterGroups(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticache",
		OperationName: "DescribeCacheParameterGroups",
	}
}
