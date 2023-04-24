// Code generated by smithy-go-codegen DO NOT EDIT.

package kinesis

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Splits a shard into two new shards in the Kinesis data stream, to increase the
// stream's capacity to ingest and transport data. SplitShard is called when there
// is a need to increase the overall capacity of a stream because of an expected
// increase in the volume of data records being ingested. This API is only
// supported for the data streams with the provisioned capacity mode. When invoking
// this API, it is recommended you use the StreamARN input parameter rather than
// the StreamName input parameter. You can also use SplitShard when a shard
// appears to be approaching its maximum utilization; for example, the producers
// sending data into the specific shard are suddenly sending more than previously
// anticipated. You can also call SplitShard to increase stream capacity, so that
// more Kinesis Data Streams applications can simultaneously read data from the
// stream for real-time processing. You must specify the shard to be split and the
// new hash key, which is the position in the shard where the shard gets split in
// two. In many cases, the new hash key might be the average of the beginning and
// ending hash key, but it can be any hash key value in the range being mapped into
// the shard. For more information, see Split a Shard (https://docs.aws.amazon.com/kinesis/latest/dev/kinesis-using-sdk-java-resharding-split.html)
// in the Amazon Kinesis Data Streams Developer Guide. You can use
// DescribeStreamSummary and the ListShards APIs to determine the shard ID and
// hash key values for the ShardToSplit and NewStartingHashKey parameters that are
// specified in the SplitShard request. SplitShard is an asynchronous operation.
// Upon receiving a SplitShard request, Kinesis Data Streams immediately returns a
// response and sets the stream status to UPDATING . After the operation is
// completed, Kinesis Data Streams sets the stream status to ACTIVE . Read and
// write operations continue to work while the stream is in the UPDATING state.
// You can use DescribeStreamSummary to check the status of the stream, which is
// returned in StreamStatus . If the stream is in the ACTIVE state, you can call
// SplitShard . If the specified stream does not exist, DescribeStreamSummary
// returns a ResourceNotFoundException . If you try to create more shards than are
// authorized for your account, you receive a LimitExceededException . For the
// default shard limit for an Amazon Web Services account, see Kinesis Data
// Streams Limits (https://docs.aws.amazon.com/kinesis/latest/dev/service-sizes-and-limits.html)
// in the Amazon Kinesis Data Streams Developer Guide. To increase this limit,
// contact Amazon Web Services Support (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html)
// . If you try to operate on too many streams simultaneously using CreateStream ,
// DeleteStream , MergeShards , and/or SplitShard , you receive a
// LimitExceededException . SplitShard has a limit of five transactions per second
// per account.
func (c *Client) SplitShard(ctx context.Context, params *SplitShardInput, optFns ...func(*Options)) (*SplitShardOutput, error) {
	if params == nil {
		params = &SplitShardInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SplitShard", params, optFns, c.addOperationSplitShardMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SplitShardOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input for SplitShard .
type SplitShardInput struct {

	// A hash key value for the starting hash key of one of the child shards created
	// by the split. The hash key range for a given shard constitutes a set of ordered
	// contiguous positive integers. The value for NewStartingHashKey must be in the
	// range of hash keys being mapped into the shard. The NewStartingHashKey hash key
	// value and all higher hash key values in hash key range are distributed to one of
	// the child shards. All the lower hash key values in the range are distributed to
	// the other child shard.
	//
	// This member is required.
	NewStartingHashKey *string

	// The shard ID of the shard to split.
	//
	// This member is required.
	ShardToSplit *string

	// The ARN of the stream.
	StreamARN *string

	// The name of the stream for the shard split.
	StreamName *string

	noSmithyDocumentSerde
}

type SplitShardOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSplitShardMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSplitShard{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSplitShard{}, middleware.After)
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
	if err = addOpSplitShardValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSplitShard(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSplitShard(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "kinesis",
		OperationName: "SplitShard",
	}
}
