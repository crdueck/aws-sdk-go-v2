// Code generated by smithy-go-codegen DO NOT EDIT.

package neptune

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/neptune/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the detailed parameter list for a particular DB parameter group.
func (c *Client) DescribeDBParameters(ctx context.Context, params *DescribeDBParametersInput, optFns ...func(*Options)) (*DescribeDBParametersOutput, error) {
	if params == nil {
		params = &DescribeDBParametersInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeDBParameters", params, optFns, c.addOperationDescribeDBParametersMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeDBParametersOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeDBParametersInput struct {

	// The name of a specific DB parameter group to return details for. Constraints:
	//   - If supplied, must match the name of an existing DBParameterGroup.
	//
	// This member is required.
	DBParameterGroupName *string

	// This parameter is not currently supported.
	Filters []types.Filter

	// An optional pagination token provided by a previous DescribeDBParameters
	// request. If this parameter is specified, the response includes only records
	// beyond the marker, up to the value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	// The parameter types to return. Default: All parameter types returned Valid
	// Values: user | system | engine-default
	Source *string

	noSmithyDocumentSerde
}

type DescribeDBParametersOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// A list of Parameter values.
	Parameters []types.Parameter

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeDBParametersMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeDBParameters{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeDBParameters{}, middleware.After)
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
	if err = addOpDescribeDBParametersValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeDBParameters(options.Region), middleware.Before); err != nil {
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

// DescribeDBParametersAPIClient is a client that implements the
// DescribeDBParameters operation.
type DescribeDBParametersAPIClient interface {
	DescribeDBParameters(context.Context, *DescribeDBParametersInput, ...func(*Options)) (*DescribeDBParametersOutput, error)
}

var _ DescribeDBParametersAPIClient = (*Client)(nil)

// DescribeDBParametersPaginatorOptions is the paginator options for
// DescribeDBParameters
type DescribeDBParametersPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeDBParametersPaginator is a paginator for DescribeDBParameters
type DescribeDBParametersPaginator struct {
	options   DescribeDBParametersPaginatorOptions
	client    DescribeDBParametersAPIClient
	params    *DescribeDBParametersInput
	nextToken *string
	firstPage bool
}

// NewDescribeDBParametersPaginator returns a new DescribeDBParametersPaginator
func NewDescribeDBParametersPaginator(client DescribeDBParametersAPIClient, params *DescribeDBParametersInput, optFns ...func(*DescribeDBParametersPaginatorOptions)) *DescribeDBParametersPaginator {
	if params == nil {
		params = &DescribeDBParametersInput{}
	}

	options := DescribeDBParametersPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeDBParametersPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeDBParametersPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeDBParameters page.
func (p *DescribeDBParametersPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeDBParametersOutput, error) {
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

	result, err := p.client.DescribeDBParameters(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeDBParameters(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "DescribeDBParameters",
	}
}
