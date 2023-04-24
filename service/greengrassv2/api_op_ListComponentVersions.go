// Code generated by smithy-go-codegen DO NOT EDIT.

package greengrassv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/greengrassv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a paginated list of all versions for a component. Greater versions
// are listed first.
func (c *Client) ListComponentVersions(ctx context.Context, params *ListComponentVersionsInput, optFns ...func(*Options)) (*ListComponentVersionsOutput, error) {
	if params == nil {
		params = &ListComponentVersionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListComponentVersions", params, optFns, c.addOperationListComponentVersionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListComponentVersionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListComponentVersionsInput struct {

	// The ARN (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the component.
	//
	// This member is required.
	Arn *string

	// The maximum number of results to be returned per paginated request.
	MaxResults *int32

	// The token to be used for the next set of paginated results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListComponentVersionsOutput struct {

	// A list of versions that exist for the component.
	ComponentVersions []types.ComponentVersionListItem

	// The token for the next set of results, or null if there are no additional
	// results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListComponentVersionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListComponentVersions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListComponentVersions{}, middleware.After)
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
	if err = addOpListComponentVersionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListComponentVersions(options.Region), middleware.Before); err != nil {
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

// ListComponentVersionsAPIClient is a client that implements the
// ListComponentVersions operation.
type ListComponentVersionsAPIClient interface {
	ListComponentVersions(context.Context, *ListComponentVersionsInput, ...func(*Options)) (*ListComponentVersionsOutput, error)
}

var _ ListComponentVersionsAPIClient = (*Client)(nil)

// ListComponentVersionsPaginatorOptions is the paginator options for
// ListComponentVersions
type ListComponentVersionsPaginatorOptions struct {
	// The maximum number of results to be returned per paginated request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListComponentVersionsPaginator is a paginator for ListComponentVersions
type ListComponentVersionsPaginator struct {
	options   ListComponentVersionsPaginatorOptions
	client    ListComponentVersionsAPIClient
	params    *ListComponentVersionsInput
	nextToken *string
	firstPage bool
}

// NewListComponentVersionsPaginator returns a new ListComponentVersionsPaginator
func NewListComponentVersionsPaginator(client ListComponentVersionsAPIClient, params *ListComponentVersionsInput, optFns ...func(*ListComponentVersionsPaginatorOptions)) *ListComponentVersionsPaginator {
	if params == nil {
		params = &ListComponentVersionsInput{}
	}

	options := ListComponentVersionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListComponentVersionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListComponentVersionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListComponentVersions page.
func (p *ListComponentVersionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListComponentVersionsOutput, error) {
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

	result, err := p.client.ListComponentVersions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListComponentVersions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "greengrass",
		OperationName: "ListComponentVersions",
	}
}
