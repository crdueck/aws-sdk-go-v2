// Code generated by smithy-go-codegen DO NOT EDIT.

package connect

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/connect/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This API is in preview release for Amazon Connect and is subject to change.
// Searches security profiles in an Amazon Connect instance, with optional
// filtering.
func (c *Client) SearchSecurityProfiles(ctx context.Context, params *SearchSecurityProfilesInput, optFns ...func(*Options)) (*SearchSecurityProfilesOutput, error) {
	if params == nil {
		params = &SearchSecurityProfilesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SearchSecurityProfiles", params, optFns, c.addOperationSearchSecurityProfilesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SearchSecurityProfilesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SearchSecurityProfilesInput struct {

	// The identifier of the Amazon Connect instance. You can find the instanceId in
	// the ARN of the instance.
	//
	// This member is required.
	InstanceId *string

	// The maximum number of results to return per page.
	MaxResults *int32

	// The token for the next set of results. Use the value returned in the previous
	// response in the next request to retrieve the next set of results.
	NextToken *string

	// The search criteria to be used to return security profiles. The name field
	// support "contains" queries with a minimum of 2 characters and maximum of 25
	// characters. Any queries with character lengths outside of this range will throw
	// invalid results. The currently supported value for FieldName: name
	SearchCriteria *types.SecurityProfileSearchCriteria

	// Filters to be applied to search results.
	SearchFilter *types.SecurityProfilesSearchFilter

	noSmithyDocumentSerde
}

type SearchSecurityProfilesOutput struct {

	// The total number of security profiles which matched your search query.
	ApproximateTotalCount *int64

	// If there are additional results, this is the token for the next set of results.
	NextToken *string

	// Information about the security profiles.
	SecurityProfiles []types.SecurityProfileSearchSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSearchSecurityProfilesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSearchSecurityProfiles{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSearchSecurityProfiles{}, middleware.After)
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
	if err = addOpSearchSecurityProfilesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSearchSecurityProfiles(options.Region), middleware.Before); err != nil {
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

// SearchSecurityProfilesAPIClient is a client that implements the
// SearchSecurityProfiles operation.
type SearchSecurityProfilesAPIClient interface {
	SearchSecurityProfiles(context.Context, *SearchSecurityProfilesInput, ...func(*Options)) (*SearchSecurityProfilesOutput, error)
}

var _ SearchSecurityProfilesAPIClient = (*Client)(nil)

// SearchSecurityProfilesPaginatorOptions is the paginator options for
// SearchSecurityProfiles
type SearchSecurityProfilesPaginatorOptions struct {
	// The maximum number of results to return per page.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// SearchSecurityProfilesPaginator is a paginator for SearchSecurityProfiles
type SearchSecurityProfilesPaginator struct {
	options   SearchSecurityProfilesPaginatorOptions
	client    SearchSecurityProfilesAPIClient
	params    *SearchSecurityProfilesInput
	nextToken *string
	firstPage bool
}

// NewSearchSecurityProfilesPaginator returns a new SearchSecurityProfilesPaginator
func NewSearchSecurityProfilesPaginator(client SearchSecurityProfilesAPIClient, params *SearchSecurityProfilesInput, optFns ...func(*SearchSecurityProfilesPaginatorOptions)) *SearchSecurityProfilesPaginator {
	if params == nil {
		params = &SearchSecurityProfilesInput{}
	}

	options := SearchSecurityProfilesPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &SearchSecurityProfilesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *SearchSecurityProfilesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next SearchSecurityProfiles page.
func (p *SearchSecurityProfilesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*SearchSecurityProfilesOutput, error) {
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

	result, err := p.client.SearchSecurityProfiles(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opSearchSecurityProfiles(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "connect",
		OperationName: "SearchSecurityProfiles",
	}
}
