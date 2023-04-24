// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Describes the specified opt-out list or all opt-out lists in your account. If
// you specify opt-out list names, the output includes information for only the
// specified opt-out lists. Opt-out lists include only those that meet the filter
// criteria. If you don't specify opt-out list names or filters, the output
// includes information for all opt-out lists. If you specify an opt-out list name
// that isn't valid, an Error is returned.
func (c *Client) DescribeOptOutLists(ctx context.Context, params *DescribeOptOutListsInput, optFns ...func(*Options)) (*DescribeOptOutListsOutput, error) {
	if params == nil {
		params = &DescribeOptOutListsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeOptOutLists", params, optFns, c.addOperationDescribeOptOutListsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeOptOutListsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeOptOutListsInput struct {

	// The maximum number of results to return per each request.
	MaxResults *int32

	// The token to be used for the next set of paginated results. You don't need to
	// supply a value for this field in the initial request.
	NextToken *string

	// The OptOutLists to show the details of. This is an array of strings that can be
	// either the OptOutListName or OptOutListArn.
	OptOutListNames []string

	noSmithyDocumentSerde
}

type DescribeOptOutListsOutput struct {

	// The token to be used for the next set of paginated results. If this field is
	// empty then there are no more results.
	NextToken *string

	// An array of OptOutListInformation objects that contain the details for the
	// requested OptOutLists.
	OptOutLists []types.OptOutListInformation

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeOptOutListsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeOptOutLists{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeOptOutLists{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeOptOutLists(options.Region), middleware.Before); err != nil {
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

// DescribeOptOutListsAPIClient is a client that implements the
// DescribeOptOutLists operation.
type DescribeOptOutListsAPIClient interface {
	DescribeOptOutLists(context.Context, *DescribeOptOutListsInput, ...func(*Options)) (*DescribeOptOutListsOutput, error)
}

var _ DescribeOptOutListsAPIClient = (*Client)(nil)

// DescribeOptOutListsPaginatorOptions is the paginator options for
// DescribeOptOutLists
type DescribeOptOutListsPaginatorOptions struct {
	// The maximum number of results to return per each request.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeOptOutListsPaginator is a paginator for DescribeOptOutLists
type DescribeOptOutListsPaginator struct {
	options   DescribeOptOutListsPaginatorOptions
	client    DescribeOptOutListsAPIClient
	params    *DescribeOptOutListsInput
	nextToken *string
	firstPage bool
}

// NewDescribeOptOutListsPaginator returns a new DescribeOptOutListsPaginator
func NewDescribeOptOutListsPaginator(client DescribeOptOutListsAPIClient, params *DescribeOptOutListsInput, optFns ...func(*DescribeOptOutListsPaginatorOptions)) *DescribeOptOutListsPaginator {
	if params == nil {
		params = &DescribeOptOutListsInput{}
	}

	options := DescribeOptOutListsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeOptOutListsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeOptOutListsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeOptOutLists page.
func (p *DescribeOptOutListsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeOptOutListsOutput, error) {
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

	result, err := p.client.DescribeOptOutLists(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeOptOutLists(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "DescribeOptOutLists",
	}
}
