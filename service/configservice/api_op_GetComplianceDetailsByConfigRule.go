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

// Returns the evaluation results for the specified Config rule. The results
// indicate which Amazon Web Services resources were evaluated by the rule, when
// each resource was last evaluated, and whether each resource complies with the
// rule.
func (c *Client) GetComplianceDetailsByConfigRule(ctx context.Context, params *GetComplianceDetailsByConfigRuleInput, optFns ...func(*Options)) (*GetComplianceDetailsByConfigRuleOutput, error) {
	if params == nil {
		params = &GetComplianceDetailsByConfigRuleInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetComplianceDetailsByConfigRule", params, optFns, c.addOperationGetComplianceDetailsByConfigRuleMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetComplianceDetailsByConfigRuleOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetComplianceDetailsByConfigRuleInput struct {

	// The name of the Config rule for which you want compliance information.
	//
	// This member is required.
	ConfigRuleName *string

	// Filters the results by compliance. INSUFFICIENT_DATA is a valid ComplianceType
	// that is returned when an Config rule cannot be evaluated. However,
	// INSUFFICIENT_DATA cannot be used as a ComplianceType for filtering results.
	ComplianceTypes []types.ComplianceType

	// The maximum number of evaluation results returned on each page. The default is
	// 10. You cannot specify a number greater than 100. If you specify 0, Config uses
	// the default.
	Limit int32

	// The nextToken string returned on a previous page that you use to get the next
	// page of results in a paginated response.
	NextToken *string

	noSmithyDocumentSerde
}

type GetComplianceDetailsByConfigRuleOutput struct {

	// Indicates whether the Amazon Web Services resource complies with the specified
	// Config rule.
	EvaluationResults []types.EvaluationResult

	// The string that you use in a subsequent request to get the next page of results
	// in a paginated response.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetComplianceDetailsByConfigRuleMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetComplianceDetailsByConfigRule{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetComplianceDetailsByConfigRule{}, middleware.After)
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
	if err = addOpGetComplianceDetailsByConfigRuleValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetComplianceDetailsByConfigRule(options.Region), middleware.Before); err != nil {
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

// GetComplianceDetailsByConfigRuleAPIClient is a client that implements the
// GetComplianceDetailsByConfigRule operation.
type GetComplianceDetailsByConfigRuleAPIClient interface {
	GetComplianceDetailsByConfigRule(context.Context, *GetComplianceDetailsByConfigRuleInput, ...func(*Options)) (*GetComplianceDetailsByConfigRuleOutput, error)
}

var _ GetComplianceDetailsByConfigRuleAPIClient = (*Client)(nil)

// GetComplianceDetailsByConfigRulePaginatorOptions is the paginator options for
// GetComplianceDetailsByConfigRule
type GetComplianceDetailsByConfigRulePaginatorOptions struct {
	// The maximum number of evaluation results returned on each page. The default is
	// 10. You cannot specify a number greater than 100. If you specify 0, Config uses
	// the default.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetComplianceDetailsByConfigRulePaginator is a paginator for
// GetComplianceDetailsByConfigRule
type GetComplianceDetailsByConfigRulePaginator struct {
	options   GetComplianceDetailsByConfigRulePaginatorOptions
	client    GetComplianceDetailsByConfigRuleAPIClient
	params    *GetComplianceDetailsByConfigRuleInput
	nextToken *string
	firstPage bool
}

// NewGetComplianceDetailsByConfigRulePaginator returns a new
// GetComplianceDetailsByConfigRulePaginator
func NewGetComplianceDetailsByConfigRulePaginator(client GetComplianceDetailsByConfigRuleAPIClient, params *GetComplianceDetailsByConfigRuleInput, optFns ...func(*GetComplianceDetailsByConfigRulePaginatorOptions)) *GetComplianceDetailsByConfigRulePaginator {
	if params == nil {
		params = &GetComplianceDetailsByConfigRuleInput{}
	}

	options := GetComplianceDetailsByConfigRulePaginatorOptions{}
	if params.Limit != 0 {
		options.Limit = params.Limit
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetComplianceDetailsByConfigRulePaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetComplianceDetailsByConfigRulePaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetComplianceDetailsByConfigRule page.
func (p *GetComplianceDetailsByConfigRulePaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetComplianceDetailsByConfigRuleOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	params.Limit = p.options.Limit

	result, err := p.client.GetComplianceDetailsByConfigRule(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetComplianceDetailsByConfigRule(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "GetComplianceDetailsByConfigRule",
	}
}
