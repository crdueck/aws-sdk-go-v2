// Code generated by smithy-go-codegen DO NOT EDIT.

package lexmodelsv2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lexmodelsv2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Get a list of bot recommendations that meet the specified criteria.
func (c *Client) ListBotRecommendations(ctx context.Context, params *ListBotRecommendationsInput, optFns ...func(*Options)) (*ListBotRecommendationsOutput, error) {
	if params == nil {
		params = &ListBotRecommendationsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListBotRecommendations", params, optFns, c.addOperationListBotRecommendationsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListBotRecommendationsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListBotRecommendationsInput struct {

	// The unique identifier of the bot that contains the bot recommendation list.
	//
	// This member is required.
	BotId *string

	// The version of the bot that contains the bot recommendation list.
	//
	// This member is required.
	BotVersion *string

	// The identifier of the language and locale of the bot recommendation list.
	//
	// This member is required.
	LocaleId *string

	// The maximum number of bot recommendations to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	MaxResults *int32

	// If the response from the ListBotRecommendation operation contains more results
	// than specified in the maxResults parameter, a token is returned in the response.
	// Use that token in the nextToken parameter to return the next page of results.
	NextToken *string

	noSmithyDocumentSerde
}

type ListBotRecommendationsOutput struct {

	// The unique identifier of the bot that contains the bot recommendation list.
	BotId *string

	// Summary information for the bot recommendations that meet the filter specified
	// in this request. The length of the list is specified in the maxResults parameter
	// of the request. If there are more bot recommendations available, the nextToken
	// field contains a token to get the next page of results.
	BotRecommendationSummaries []types.BotRecommendationSummary

	// The version of the bot that contains the bot recommendation list.
	BotVersion *string

	// The identifier of the language and locale of the bot recommendation list.
	LocaleId *string

	// A token that indicates whether there are more results to return in a response
	// to the ListBotRecommendations operation. If the nextToken field is present, you
	// send the contents as the nextToken parameter of a ListBotRecommendations
	// operation request to get the next page of results.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListBotRecommendationsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListBotRecommendations{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListBotRecommendations{}, middleware.After)
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
	if err = addOpListBotRecommendationsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListBotRecommendations(options.Region), middleware.Before); err != nil {
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

// ListBotRecommendationsAPIClient is a client that implements the
// ListBotRecommendations operation.
type ListBotRecommendationsAPIClient interface {
	ListBotRecommendations(context.Context, *ListBotRecommendationsInput, ...func(*Options)) (*ListBotRecommendationsOutput, error)
}

var _ ListBotRecommendationsAPIClient = (*Client)(nil)

// ListBotRecommendationsPaginatorOptions is the paginator options for
// ListBotRecommendations
type ListBotRecommendationsPaginatorOptions struct {
	// The maximum number of bot recommendations to return in each page of results. If
	// there are fewer results than the max page size, only the actual number of
	// results are returned.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListBotRecommendationsPaginator is a paginator for ListBotRecommendations
type ListBotRecommendationsPaginator struct {
	options   ListBotRecommendationsPaginatorOptions
	client    ListBotRecommendationsAPIClient
	params    *ListBotRecommendationsInput
	nextToken *string
	firstPage bool
}

// NewListBotRecommendationsPaginator returns a new ListBotRecommendationsPaginator
func NewListBotRecommendationsPaginator(client ListBotRecommendationsAPIClient, params *ListBotRecommendationsInput, optFns ...func(*ListBotRecommendationsPaginatorOptions)) *ListBotRecommendationsPaginator {
	if params == nil {
		params = &ListBotRecommendationsInput{}
	}

	options := ListBotRecommendationsPaginatorOptions{}
	if params.MaxResults != nil {
		options.Limit = *params.MaxResults
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListBotRecommendationsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListBotRecommendationsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListBotRecommendations page.
func (p *ListBotRecommendationsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListBotRecommendationsOutput, error) {
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

	result, err := p.client.ListBotRecommendations(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListBotRecommendations(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lex",
		OperationName: "ListBotRecommendations",
	}
}
