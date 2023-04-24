// Code generated by smithy-go-codegen DO NOT EDIT.

package datasync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/datasync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Returns metadata about a task.
func (c *Client) DescribeTask(ctx context.Context, params *DescribeTaskInput, optFns ...func(*Options)) (*DescribeTaskOutput, error) {
	if params == nil {
		params = &DescribeTaskInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeTask", params, optFns, c.addOperationDescribeTaskMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeTaskOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// DescribeTaskRequest
type DescribeTaskInput struct {

	// The Amazon Resource Name (ARN) of the task to describe.
	//
	// This member is required.
	TaskArn *string

	noSmithyDocumentSerde
}

// DescribeTaskResponse
type DescribeTaskOutput struct {

	// The Amazon Resource Name (ARN) of the Amazon CloudWatch log group that was used
	// to monitor and log events in the task. For more information on these groups, see
	// Working with Log Groups and Log Streams in the Amazon CloudWatch User Guide.
	CloudWatchLogGroupArn *string

	// The time that the task was created.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the task execution that is transferring files.
	CurrentTaskExecutionArn *string

	// The Amazon Resource Name (ARN) of the Amazon Web Services storage resource's
	// location.
	DestinationLocationArn *string

	// The Amazon Resource Names (ARNs) of the network interfaces created for your
	// destination location. For more information, see Network interface requirements (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// .
	DestinationNetworkInterfaceArns []string

	// Errors that DataSync encountered during execution of the task. You can use this
	// error code to help troubleshoot issues.
	ErrorCode *string

	// Detailed description of an error that was encountered during the task
	// execution. You can use this information to help troubleshoot issues.
	ErrorDetail *string

	// A list of filter rules that exclude specific data during your transfer. For
	// more information and examples, see Filtering data transferred by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Excludes []types.FilterRule

	// A list of filter rules that include specific data during your transfer. For
	// more information and examples, see Filtering data transferred by DataSync (https://docs.aws.amazon.com/datasync/latest/userguide/filtering.html)
	// .
	Includes []types.FilterRule

	// The name of the task that was described.
	Name *string

	// The configuration options that control the behavior of the StartTaskExecution
	// operation. Some options include preserving file or object metadata and verifying
	// data integrity. You can override these options for each task execution. For more
	// information, see StartTaskExecution (https://docs.aws.amazon.com/datasync/latest/userguide/API_StartTaskExecution.html)
	// .
	Options *types.Options

	// The schedule used to periodically transfer files from a source to a destination
	// location.
	Schedule *types.TaskSchedule

	// The Amazon Resource Name (ARN) of the source file system's location.
	SourceLocationArn *string

	// The Amazon Resource Names (ARNs) of the network interfaces created for your
	// source location. For more information, see Network interface requirements (https://docs.aws.amazon.com/datasync/latest/userguide/datasync-network.html#required-network-interfaces)
	// .
	SourceNetworkInterfaceArns []string

	// The status of the task that was described. For detailed information about task
	// execution statuses, see Understanding Task Statuses in the DataSync User Guide.
	Status types.TaskStatus

	// The Amazon Resource Name (ARN) of the task that was described.
	TaskArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeTaskMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeTask{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeTask{}, middleware.After)
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
	if err = addOpDescribeTaskValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeTask(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeTask(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "datasync",
		OperationName: "DescribeTask",
	}
}
