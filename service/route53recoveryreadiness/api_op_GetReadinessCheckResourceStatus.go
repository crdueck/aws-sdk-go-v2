// Code generated by smithy-go-codegen DO NOT EDIT.

package route53recoveryreadiness

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/route53recoveryreadiness/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets individual readiness status for a readiness check. To see the overall
// readiness status for a recovery group, that considers the readiness status for
// all the readiness checks in the recovery group, use
// GetRecoveryGroupReadinessSummary.
func (c *Client) GetReadinessCheckResourceStatus(ctx context.Context, params *GetReadinessCheckResourceStatusInput, optFns ...func(*Options)) (*GetReadinessCheckResourceStatusOutput, error) {
	if params == nil {
		params = &GetReadinessCheckResourceStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetReadinessCheckResourceStatus", params, optFns, c.addOperationGetReadinessCheckResourceStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetReadinessCheckResourceStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetReadinessCheckResourceStatusInput struct {

	// Name of a readiness check.
	//
	// This member is required.
	ReadinessCheckName *string

	// The resource identifier, which is the Amazon Resource Name (ARN) or the
	// identifier generated for the resource by Application Recovery Controller (for
	// example, for a DNS target resource).
	//
	// This member is required.
	ResourceIdentifier *string

	// The number of objects that you want to return with this call.
	MaxResults int32

	// The token that identifies which batch of results you want to see.
	NextToken *string

	noSmithyDocumentSerde
}

type GetReadinessCheckResourceStatusOutput struct {

	// The token that identifies which batch of results you want to see.
	NextToken *string

	// The readiness at a rule level.
	Readiness types.Readiness

	// Details of the rule's results.
	Rules []types.RuleResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetReadinessCheckResourceStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetReadinessCheckResourceStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetReadinessCheckResourceStatus{}, middleware.After)
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
	if err = addOpGetReadinessCheckResourceStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetReadinessCheckResourceStatus(options.Region), middleware.Before); err != nil {
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

// GetReadinessCheckResourceStatusAPIClient is a client that implements the
// GetReadinessCheckResourceStatus operation.
type GetReadinessCheckResourceStatusAPIClient interface {
	GetReadinessCheckResourceStatus(context.Context, *GetReadinessCheckResourceStatusInput, ...func(*Options)) (*GetReadinessCheckResourceStatusOutput, error)
}

var _ GetReadinessCheckResourceStatusAPIClient = (*Client)(nil)

// GetReadinessCheckResourceStatusPaginatorOptions is the paginator options for
// GetReadinessCheckResourceStatus
type GetReadinessCheckResourceStatusPaginatorOptions struct {
	// The number of objects that you want to return with this call.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetReadinessCheckResourceStatusPaginator is a paginator for
// GetReadinessCheckResourceStatus
type GetReadinessCheckResourceStatusPaginator struct {
	options   GetReadinessCheckResourceStatusPaginatorOptions
	client    GetReadinessCheckResourceStatusAPIClient
	params    *GetReadinessCheckResourceStatusInput
	nextToken *string
	firstPage bool
}

// NewGetReadinessCheckResourceStatusPaginator returns a new
// GetReadinessCheckResourceStatusPaginator
func NewGetReadinessCheckResourceStatusPaginator(client GetReadinessCheckResourceStatusAPIClient, params *GetReadinessCheckResourceStatusInput, optFns ...func(*GetReadinessCheckResourceStatusPaginatorOptions)) *GetReadinessCheckResourceStatusPaginator {
	if params == nil {
		params = &GetReadinessCheckResourceStatusInput{}
	}

	options := GetReadinessCheckResourceStatusPaginatorOptions{}
	if params.MaxResults != 0 {
		options.Limit = params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetReadinessCheckResourceStatusPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetReadinessCheckResourceStatusPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetReadinessCheckResourceStatus page.
func (p *GetReadinessCheckResourceStatusPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetReadinessCheckResourceStatusOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.MaxResults = p.options.Limit

	result, err := p.client.GetReadinessCheckResourceStatus(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetReadinessCheckResourceStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "route53-recovery-readiness",
		OperationName: "GetReadinessCheckResourceStatus",
	}
}
