// Code generated by smithy-go-codegen DO NOT EDIT.

package rolesanywhere

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/rolesanywhere/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Lists the subjects in the authenticated account and Amazon Web Services Region.
// Required permissions: rolesanywhere:ListSubjects .
func (c *Client) ListSubjects(ctx context.Context, params *ListSubjectsInput, optFns ...func(*Options)) (*ListSubjectsOutput, error) {
	if params == nil {
		params = &ListSubjectsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListSubjects", params, optFns, c.addOperationListSubjectsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListSubjectsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListSubjectsInput struct {

	// A token that indicates where the output should continue from, if a previous
	// operation did not show all results. To get the next results, call the operation
	// again with this value.
	NextToken *string

	// The number of resources in the paginated list.
	PageSize *int32

	noSmithyDocumentSerde
}

type ListSubjectsOutput struct {

	// A token that indicates where the output should continue from, if a previous
	// operation did not show all results. To get the next results, call the operation
	// again with this value.
	NextToken *string

	// A list of subjects.
	Subjects []types.SubjectSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListSubjectsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpListSubjects{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpListSubjects{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListSubjects(options.Region), middleware.Before); err != nil {
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

// ListSubjectsAPIClient is a client that implements the ListSubjects operation.
type ListSubjectsAPIClient interface {
	ListSubjects(context.Context, *ListSubjectsInput, ...func(*Options)) (*ListSubjectsOutput, error)
}

var _ ListSubjectsAPIClient = (*Client)(nil)

// ListSubjectsPaginatorOptions is the paginator options for ListSubjects
type ListSubjectsPaginatorOptions struct {
	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// ListSubjectsPaginator is a paginator for ListSubjects
type ListSubjectsPaginator struct {
	options   ListSubjectsPaginatorOptions
	client    ListSubjectsAPIClient
	params    *ListSubjectsInput
	nextToken *string
	firstPage bool
}

// NewListSubjectsPaginator returns a new ListSubjectsPaginator
func NewListSubjectsPaginator(client ListSubjectsAPIClient, params *ListSubjectsInput, optFns ...func(*ListSubjectsPaginatorOptions)) *ListSubjectsPaginator {
	if params == nil {
		params = &ListSubjectsInput{}
	}

	options := ListSubjectsPaginatorOptions{}

	for _, fn := range optFns {
		fn(&options)
	}

	return &ListSubjectsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.NextToken,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *ListSubjectsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next ListSubjects page.
func (p *ListSubjectsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*ListSubjectsOutput, error) {
	if !p.HasMorePages() {
		return nil, fmt.Errorf("no more pages available")
	}

	params := *p.params
	params.NextToken = p.nextToken

	result, err := p.client.ListSubjects(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opListSubjects(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rolesanywhere",
		OperationName: "ListSubjects",
	}
}
