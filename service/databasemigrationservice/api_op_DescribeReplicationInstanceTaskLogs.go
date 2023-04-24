// Code generated by smithy-go-codegen DO NOT EDIT.

package databasemigrationservice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databasemigrationservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns information about the task logs for the specified task.
func (c *Client) DescribeReplicationInstanceTaskLogs(ctx context.Context, params *DescribeReplicationInstanceTaskLogsInput, optFns ...func(*Options)) (*DescribeReplicationInstanceTaskLogsOutput, error) {
	if params == nil {
		params = &DescribeReplicationInstanceTaskLogsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReplicationInstanceTaskLogs", params, optFns, c.addOperationDescribeReplicationInstanceTaskLogsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReplicationInstanceTaskLogsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReplicationInstanceTaskLogsInput struct {

	// The Amazon Resource Name (ARN) of the replication instance.
	//
	// This member is required.
	ReplicationInstanceArn *string

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	MaxRecords *int32

	noSmithyDocumentSerde
}

type DescribeReplicationInstanceTaskLogsOutput struct {

	// An optional pagination token provided by a previous request. If this parameter
	// is specified, the response includes only records beyond the marker, up to the
	// value specified by MaxRecords .
	Marker *string

	// The Amazon Resource Name (ARN) of the replication instance.
	ReplicationInstanceArn *string

	// An array of replication task log metadata. Each member of the array contains
	// the replication task name, ARN, and task log size (in bytes).
	ReplicationInstanceTaskLogs []types.ReplicationInstanceTaskLog

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReplicationInstanceTaskLogsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReplicationInstanceTaskLogs{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReplicationInstanceTaskLogs{}, middleware.After)
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
	if err = addOpDescribeReplicationInstanceTaskLogsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReplicationInstanceTaskLogs(options.Region), middleware.Before); err != nil {
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

// DescribeReplicationInstanceTaskLogsAPIClient is a client that implements the
// DescribeReplicationInstanceTaskLogs operation.
type DescribeReplicationInstanceTaskLogsAPIClient interface {
	DescribeReplicationInstanceTaskLogs(context.Context, *DescribeReplicationInstanceTaskLogsInput, ...func(*Options)) (*DescribeReplicationInstanceTaskLogsOutput, error)
}

var _ DescribeReplicationInstanceTaskLogsAPIClient = (*Client)(nil)

// DescribeReplicationInstanceTaskLogsPaginatorOptions is the paginator options
// for DescribeReplicationInstanceTaskLogs
type DescribeReplicationInstanceTaskLogsPaginatorOptions struct {
	// The maximum number of records to include in the response. If more records exist
	// than the specified MaxRecords value, a pagination token called a marker is
	// included in the response so that the remaining results can be retrieved.
	// Default: 100 Constraints: Minimum 20, maximum 100.
	Limit int32

	// Set to true if pagination should stop if the service returns a pagination token
	// that matches the most recent token provided to the service.
	StopOnDuplicateToken bool
}

// DescribeReplicationInstanceTaskLogsPaginator is a paginator for
// DescribeReplicationInstanceTaskLogs
type DescribeReplicationInstanceTaskLogsPaginator struct {
	options   DescribeReplicationInstanceTaskLogsPaginatorOptions
	client    DescribeReplicationInstanceTaskLogsAPIClient
	params    *DescribeReplicationInstanceTaskLogsInput
	nextToken *string
	firstPage bool
}

// NewDescribeReplicationInstanceTaskLogsPaginator returns a new
// DescribeReplicationInstanceTaskLogsPaginator
func NewDescribeReplicationInstanceTaskLogsPaginator(client DescribeReplicationInstanceTaskLogsAPIClient, params *DescribeReplicationInstanceTaskLogsInput, optFns ...func(*DescribeReplicationInstanceTaskLogsPaginatorOptions)) *DescribeReplicationInstanceTaskLogsPaginator {
	if params == nil {
		params = &DescribeReplicationInstanceTaskLogsInput{}
	}

	options := DescribeReplicationInstanceTaskLogsPaginatorOptions{}
	if params.MaxRecords != nil {
		options.Limit = *params.MaxRecords
	}

	for _, fn := range optFns {
		fn(&options)
	}

	return &DescribeReplicationInstanceTaskLogsPaginator{
		options:   options,
		client:    client,
		params:    params,
		firstPage: true,
		nextToken: params.Marker,
	}
}

// HasMorePages returns a boolean indicating whether more pages are available
func (p *DescribeReplicationInstanceTaskLogsPaginator) HasMorePages() bool {
	return p.firstPage || (p.nextToken != nil && len(*p.nextToken) != 0)
}

// NextPage retrieves the next DescribeReplicationInstanceTaskLogs page.
func (p *DescribeReplicationInstanceTaskLogsPaginator) NextPage(ctx context.Context, optFns ...func(*Options)) (*DescribeReplicationInstanceTaskLogsOutput, error) {
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

	result, err := p.client.DescribeReplicationInstanceTaskLogs(ctx, &params, optFns...)
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

func newServiceMetadataMiddleware_opDescribeReplicationInstanceTaskLogs(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dms",
		OperationName: "DescribeReplicationInstanceTaskLogs",
	}
}
