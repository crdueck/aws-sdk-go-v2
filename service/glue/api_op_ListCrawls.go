// Code generated by smithy-go-codegen DO NOT EDIT.

package glue

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/glue/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns all the crawls of a specified crawler. Returns only the crawls that
// have occurred since the launch date of the crawler history feature, and only
// retains up to 12 months of crawls. Older crawls will not be returned. You may
// use this API to:
//   - Retrive all the crawls of a specified crawler.
//   - Retrieve all the crawls of a specified crawler within a limited count.
//   - Retrieve all the crawls of a specified crawler in a specific time range.
//   - Retrieve all the crawls of a specified crawler with a particular state,
//     crawl ID, or DPU hour value.
func (c *Client) ListCrawls(ctx context.Context, params *ListCrawlsInput, optFns ...func(*Options)) (*ListCrawlsOutput, error) {
	if params == nil {
		params = &ListCrawlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListCrawls", params, optFns, c.addOperationListCrawlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListCrawlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListCrawlsInput struct {

	// The name of the crawler whose runs you want to retrieve.
	//
	// This member is required.
	CrawlerName *string

	// Filters the crawls by the criteria you specify in a list of CrawlsFilter
	// objects.
	Filters []types.CrawlsFilter

	// The maximum number of results to return. The default is 20, and maximum is 100.
	MaxResults *int32

	// A continuation token, if this is a continuation call.
	NextToken *string

	noSmithyDocumentSerde
}

type ListCrawlsOutput struct {

	// A list of CrawlerHistory objects representing the crawl runs that meet your
	// criteria.
	Crawls []types.CrawlerHistory

	// A continuation token for paginating the returned list of tokens, returned if
	// the current segment of the list is not the last.
	NextToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListCrawlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListCrawls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListCrawls{}, middleware.After)
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
	if err = addOpListCrawlsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListCrawls(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListCrawls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "glue",
		OperationName: "ListCrawls",
	}
}
