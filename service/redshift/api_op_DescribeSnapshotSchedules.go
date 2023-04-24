// Code generated by smithy-go-codegen DO NOT EDIT.

package redshift

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/redshift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a list of snapshot schedules.
func (c *Client) DescribeSnapshotSchedules(ctx context.Context, params *DescribeSnapshotSchedulesInput, optFns ...func(*Options)) (*DescribeSnapshotSchedulesOutput, error) {
	if params == nil {
		params = &DescribeSnapshotSchedulesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeSnapshotSchedules", params, optFns, c.addOperationDescribeSnapshotSchedulesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeSnapshotSchedulesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeSnapshotSchedulesInput struct {

	// The unique identifier for the cluster whose snapshot schedules you want to view.
	ClusterIdentifier *string

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// The maximum number or response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	MaxRecords *int32

	// A unique identifier for a snapshot schedule.
	ScheduleIdentifier *string

	// The key value for a snapshot schedule tag.
	TagKeys []string

	// The value corresponding to the key of the snapshot schedule tag.
	TagValues []string

	noSmithyDocumentSerde
}

type DescribeSnapshotSchedulesOutput struct {

	// A value that indicates the starting point for the next set of response records
	// in a subsequent request. If a value is returned in a response, you can retrieve
	// the next set of records by providing this returned marker value in the marker
	// parameter and retrying the command. If the marker field is empty, all response
	// records have been retrieved for the request.
	Marker *string

	// A list of SnapshotSchedules.
	SnapshotSchedules []types.SnapshotSchedule

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeSnapshotSchedulesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpDescribeSnapshotSchedules{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpDescribeSnapshotSchedules{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeSnapshotSchedules(options.Region), middleware.Before); err != nil {
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

// DescribeSnapshotSchedulesAPIClient is a client that implements the
// DescribeSnapshotSchedules operation.
type DescribeSnapshotSchedulesAPIClient interface {
	DescribeSnapshotSchedules(context.Context, *DescribeSnapshotSchedulesInput, ...func(*Options)) (*DescribeSnapshotSchedulesOutput, error)
}

var _ DescribeSnapshotSchedulesAPIClient = (*Client)(nil)

// DescribeSnapshotSchedulesPaginatorOptions is the paginator options for
// DescribeSnapshotSchedules
type DescribeSnapshotSchedulesPaginatorOptions struct {
	// The maximum number or response records to return in each call. If the number of
	// remaining response records exceeds the specified MaxRecords value, a value is
	// returned in a marker field of the response. You can retrieve the next set of
	// records by retrying the command with the returned marker value.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeSnapshotSchedulesPaginator is a paginator for DescribeSnapshotSchedules
type DescribeSnapshotSchedulesPaginator struct {
	options   DescribeSnapshotSchedulesPaginatorOptions
	client    DescribeSnapshotSchedulesAPIClient
	params    *DescribeSnapshotSchedulesInput
	nextToken *string
	firstPage bool
}

// NewDescribeSnapshotSchedulesPaginator returns a new
// DescribeSnapshotSchedulesPaginator
func NewDescribeSnapshotSchedulesPaginator(client DescribeSnapshotSchedulesAPIClient, params *DescribeSnapshotSchedulesInput, optFns ...func(*DescribeSnapshotSchedulesPaginatorOptions)) *DescribeSnapshotSchedulesPaginator {
	if params == nil {
		params = &DescribeSnapshotSchedulesInput{}
	}

	options := DescribeSnapshotSchedulesPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeSnapshotSchedulesPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeSnapshotSchedulesPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeSnapshotSchedules page.
func (p *DescribeSnapshotSchedulesPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeSnapshotSchedulesOutput, error) {
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

	result, err := p.client.DescribeSnapshotSchedules(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeSnapshotSchedules(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "redshift",
		OperationName: "DescribeSnapshotSchedules",
	}
}
