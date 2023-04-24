// Code generated by smithy-go-codegen DO NOT EDIT.

package dynamodbstreams

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns a shard iterator. A shard iterator provides information about how to
// retrieve the stream records from within a shard. Use the shard iterator in a
// subsequent GetRecords request to read the stream records from the shard. A
// shard iterator expires 15 minutes after it is returned to the requester.
func (c *Client) GetShardIterator(ctx context.Context, params *GetShardIteratorInput, optFns ...func(*Options)) (*GetShardIteratorOutput, error) {
	if params == nil {
		params = &GetShardIteratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetShardIterator", params, optFns, c.addOperationGetShardIteratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetShardIteratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input of a GetShardIterator operation.
type GetShardIteratorInput struct {

	// The identifier of the shard. The iterator will be returned for this shard ID.
	//
	// This member is required.
	ShardId *string

	// Determines how the shard iterator is used to start reading stream records from
	// the shard:
	//   - AT_SEQUENCE_NUMBER - Start reading exactly from the position denoted by a
	//   specific sequence number.
	//   - AFTER_SEQUENCE_NUMBER - Start reading right after the position denoted by a
	//   specific sequence number.
	//   - TRIM_HORIZON - Start reading at the last (untrimmed) stream record, which is
	//   the oldest record in the shard. In DynamoDB Streams, there is a 24 hour limit on
	//   data retention. Stream records whose age exceeds this limit are subject to
	//   removal (trimming) from the stream.
	//   - LATEST - Start reading just after the most recent stream record in the
	//   shard, so that you always read the most recent data in the shard.
	//
	// This member is required.
	ShardIteratorType types.ShardIteratorType

	// The Amazon Resource Name (ARN) for the stream.
	//
	// This member is required.
	StreamArn *string

	// The sequence number of a stream record in the shard from which to start reading.
	SequenceNumber *string

	noSmithyDocumentSerde
}

// Represents the output of a GetShardIterator operation.
type GetShardIteratorOutput struct {

	// The position in the shard from which to start reading stream records
	// sequentially. A shard iterator specifies this position using the sequence number
	// of a stream record in a shard.
	ShardIterator *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetShardIteratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpGetShardIterator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpGetShardIterator{}, middleware.After)
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
	if err = addOpGetShardIteratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetShardIterator(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetShardIterator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "dynamodb",
		OperationName: "GetShardIterator",
	}
}
