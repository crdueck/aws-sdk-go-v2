// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Indicates whether the specified Amazon Web Services resources are compliant. If
// a resource is noncompliant, this action returns the number of Config rules that
// the resource does not comply with. A resource is compliant if it complies with
// all the Config rules that evaluate it. It is noncompliant if it does not comply
// with one or more of these rules. If Config has no current evaluation results for
// the resource, it returns INSUFFICIENT_DATA . This result might indicate one of
// the following conditions about the rules that evaluate the resource:
//   - Config has never invoked an evaluation for the rule. To check whether it
//     has, use the DescribeConfigRuleEvaluationStatus action to get the
//     LastSuccessfulInvocationTime and LastFailedInvocationTime .
//   - The rule's Lambda function is failing to send evaluation results to Config.
//     Verify that the role that you assigned to your configuration recorder includes
//     the config:PutEvaluations permission. If the rule is a custom rule, verify
//     that the Lambda execution role includes the config:PutEvaluations permission.
//   - The rule's Lambda function has returned NOT_APPLICABLE for all evaluation
//     results. This can occur if the resources were deleted or removed from the rule's
//     scope.
func (c *Client) DescribeComplianceByResource(ctx context.Context, params *DescribeComplianceByResourceInput, optFns ...func(*Options)) (*DescribeComplianceByResourceOutput, error) {
	if params == nil {
		params = &DescribeComplianceByResourceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeComplianceByResource", params, optFns, c.addOperationDescribeComplianceByResourceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeComplianceByResourceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeComplianceByResourceInput struct {

	// Filters the results by compliance.
	ComplianceTypes []types.ComplianceType

	// The maximum number of evaluation results returned on each page. The default is
	// 10. You cannot specify a number greater than 100. If you specify 0, Config uses
	// the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	// The ID of the Amazon Web Services resource for which you want compliance
	// information. You can specify only one resource ID. If you specify a resource ID,
	// you must also specify a type for ResourceType .
	ResourceId *string

	// The types of Amazon Web Services resources for which you want compliance
	// information (for example, AWS::EC2::Instance ). For this action, you can specify
	// that the resource type is an Amazon Web Services account by specifying
	// AWS::::Account .
	ResourceType *string

	noSmithyDocumentSerde
}

type DescribeComplianceByResourceOutput struct {

	// Indicates whether the specified Amazon Web Services resource complies with all
	// of the Config rules that evaluate it.
	ComplianceByResources []types.ComplianceByResource

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeComplianceByResourceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeComplianceByResource{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeComplianceByResource{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeComplianceByResource(options.Region), middleware.Before); err != nil {
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

// DescribeComplianceByResourceAPIClient is a client that implements the
// DescribeComplianceByResource operation.
type DescribeComplianceByResourceAPIClient interface {
	DescribeComplianceByResource(context.Context, *DescribeComplianceByResourceInput, ...func(*Options)) (*DescribeComplianceByResourceOutput, error)
}

var _ DescribeComplianceByResourceAPIClient = (*Client)(nil)

// DescribeComplianceByResourcePaginatorOptions is the paginator options for
// DescribeComplianceByResource
type DescribeComplianceByResourcePaginatorOptions struct {
	// The maximum number of evaluation results returned on each page. The default is
	// 10. You cannot specify a number greater than 100. If you specify 0, Config uses
	// the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeComplianceByResourcePaginator is a paginator for
// DescribeComplianceByResource
type DescribeComplianceByResourcePaginator struct {
	options   DescribeComplianceByResourcePaginatorOptions
	client    DescribeComplianceByResourceAPIClient
	params    *DescribeComplianceByResourceInput
	nextToken *string
	firstPage bool
}

// NewDescribeComplianceByResourcePaginator returns a new
// DescribeComplianceByResourcePaginator
func NewDescribeComplianceByResourcePaginator(client DescribeComplianceByResourceAPIClient, params *DescribeComplianceByResourceInput, optFns ...func(*DescribeComplianceByResourcePaginatorOptions)) *DescribeComplianceByResourcePaginator {
	if params == nil {
		params = &DescribeComplianceByResourceInput{}
	}

	options := DescribeComplianceByResourcePaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeComplianceByResourcePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeComplianceByResourcePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeComplianceByResource page.
func (p *DescribeComplianceByResourcePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeComplianceByResourceOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.DescribeComplianceByResource(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeComplianceByResource(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "DescribeComplianceByResource",
	}
}
