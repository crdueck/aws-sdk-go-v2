// Code generated by smithy-go-codegen DO NOT EDIT.

package iotthingsgraph

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotthingsgraph/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets revisions of the specified workflow. Only the last 100 revisions are
// stored. If the workflow has been deprecated, this action will return revisions
// that occurred before the deprecation. This action won't work for workflows that
// have been deleted.
//
// Deprecated: since: 2022-08-30
func (c *Client) GetFlowTemplateRevisions(ctx context.Context, params *GetFlowTemplateRevisionsInput, optFns ...func(*Options)) (*GetFlowTemplateRevisionsOutput, error) {
	if params == nil {
		params = &GetFlowTemplateRevisionsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetFlowTemplateRevisions", params, optFns, c.addOperationGetFlowTemplateRevisionsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetFlowTemplateRevisionsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetFlowTemplateRevisionsInput struct {

	// The ID of the workflow. The ID should be in the following format.
	// urn:tdm:REGION/ACCOUNT ID/default:workflow:WORKFLOWNAME
	//
	// This member is required.
	Id *string

	// The maximum number of results to return in the response.
	MaxResults *int32

	// The string that specifies the next page of results. Use this when you're
	// paginating results.
	NextToken *string

	noSmithyDocumentSerde
}

type GetFlowTemplateRevisionsOutput struct {

	// The string to specify as nextToken when you request the next page of results.
	NextToken *string

	// An array of objects that provide summary data about each revision.
	Summaries []types.FlowTemplateSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetFlowTemplateRevisionsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetFlowTemplateRevisions{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetFlowTemplateRevisions{}, middleware.After)
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
	if err = addOpGetFlowTemplateRevisionsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetFlowTemplateRevisions(options.Region), middleware.Before); err != nil {
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

// GetFlowTemplateRevisionsAPIClient is a client that implements the
// GetFlowTemplateRevisions operation.
type GetFlowTemplateRevisionsAPIClient interface {
	GetFlowTemplateRevisions(context.Context, *GetFlowTemplateRevisionsInput, ...func(*Options)) (*GetFlowTemplateRevisionsOutput, error)
}

var _ GetFlowTemplateRevisionsAPIClient = (*Client)(nil)

// GetFlowTemplateRevisionsPaginatorOptions is the paginator options for
// GetFlowTemplateRevisions
type GetFlowTemplateRevisionsPaginatorOptions struct {
	// The maximum number of results to return in the response.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// GetFlowTemplateRevisionsPaginator is a paginator for GetFlowTemplateRevisions
type GetFlowTemplateRevisionsPaginator struct {
	options   GetFlowTemplateRevisionsPaginatorOptions
	client    GetFlowTemplateRevisionsAPIClient
	params    *GetFlowTemplateRevisionsInput
	nextToken *string
	firstPage bool
}

// NewGetFlowTemplateRevisionsPaginator returns a new
// GetFlowTemplateRevisionsPaginator
func NewGetFlowTemplateRevisionsPaginator(client GetFlowTemplateRevisionsAPIClient, params *GetFlowTemplateRevisionsInput, optFns ...func(*GetFlowTemplateRevisionsPaginatorOptions)) *GetFlowTemplateRevisionsPaginator {
	if params == nil {
		params = &GetFlowTemplateRevisionsInput{}
	}

	options := GetFlowTemplateRevisionsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &GetFlowTemplateRevisionsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *GetFlowTemplateRevisionsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next GetFlowTemplateRevisions page.
func (p *GetFlowTemplateRevisionsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*GetFlowTemplateRevisionsOutput, error) {
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

	result, err := p.client.GetFlowTemplateRevisions(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opGetFlowTemplateRevisions(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotthingsgraph",
		OperationName: "GetFlowTemplateRevisions",
	}
}
