// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudtrail

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudtrail/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns all public keys whose private keys were used to sign the digest files
// within the specified time range. The public key is needed to validate digest
// files that were signed with its corresponding private key. CloudTrail uses
// different private and public key pairs per region. Each digest file is signed
// with a private key unique to its region. When you validate a digest file from a
// specific region, you must look in the same region for its corresponding public
// key.
func (c *Client) ListPublicKeys(ctx context.Context, params *ListPublicKeysInput, optFns ...func(*Options)) (*ListPublicKeysOutput, error) {
	if params == nil {
		params = &ListPublicKeysInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListPublicKeys", params, optFns, c.addOperationListPublicKeysMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListPublicKeysOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Requests the public keys for a specified time range.
type ListPublicKeysInput struct {

	// Optionally specifies, in UTC, the end of the time range to look up public keys
	// for CloudTrail digest files. If not specified, the current time is used.
	EndTime *time.Time

	// Reserved for future use.
	NextToken *string

	// Optionally specifies, in UTC, the start of the time range to look up public
	// keys for CloudTrail digest files. If not specified, the current time is used,
	// and the current public key is returned.
	StartTime *time.Time

	noSmithyDocumentSerde
}

// Returns the objects or data listed below if successful. Otherwise, returns an
// error.
type ListPublicKeysOutput struct {

	// Reserved for future use.
	NextToken *string

	// Contains an array of PublicKey objects. The returned public keys may have
	// validity time ranges that overlap.
	PublicKeyList []types.PublicKey

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListPublicKeysMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpListPublicKeys{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpListPublicKeys{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListPublicKeys(options.Region), middleware.Before); err != nil {
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

// ListPublicKeysAPIClient is a client that implements the ListPublicKeys
// operation.
type ListPublicKeysAPIClient interface {
	ListPublicKeys(context.Context, *ListPublicKeysInput, ...func(*Options)) (*ListPublicKeysOutput, error)
}

var _ ListPublicKeysAPIClient = (*Client)(nil)

// ListPublicKeysPaginatorOptions is the paginator options for ListPublicKeys
type ListPublicKeysPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListPublicKeysPaginator is a paginator for ListPublicKeys
type ListPublicKeysPaginator struct {
	options   ListPublicKeysPaginatorOptions
	client    ListPublicKeysAPIClient
	params    *ListPublicKeysInput
	nextToken *string
	firstPage bool
}

// NewListPublicKeysPaginator returns a new ListPublicKeysPaginator
func NewListPublicKeysPaginator(client ListPublicKeysAPIClient, params *ListPublicKeysInput, optFns ...func(*ListPublicKeysPaginatorOptions)) *ListPublicKeysPaginator {
	if params == nil {
		params = &ListPublicKeysInput{}
	}

	options := ListPublicKeysPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListPublicKeysPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListPublicKeysPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListPublicKeys page.
func (p *ListPublicKeysPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListPublicKeysOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListPublicKeys(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListPublicKeys(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudtrail",
		OperationName: "ListPublicKeys",
	}
}
