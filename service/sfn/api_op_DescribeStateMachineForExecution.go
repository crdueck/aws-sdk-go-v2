// Code generated by smithy-go-codegen DO NOT EDIT.

package sfn

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sfn/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Provides information about a state machine's definition, its execution role
// ARN, and configuration. If an execution was dispatched by a Map Run, the Map Run
// is returned in the response. Additionally, the state machine returned will be
// the state machine associated with the Map Run. This operation is eventually
// consistent. The results are best effort and may not reflect very recent updates
// and changes. This API action is not supported by EXPRESS state machines.
func (c *Client) DescribeStateMachineForExecution(ctx context.Context, params *DescribeStateMachineForExecutionInput, optFns ...func(*Options)) (*DescribeStateMachineForExecutionOutput, error) {
	if params == nil {
		params = &DescribeStateMachineForExecutionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeStateMachineForExecution", params, optFns, c.addOperationDescribeStateMachineForExecutionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeStateMachineForExecutionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeStateMachineForExecutionInput struct {

	// The Amazon Resource Name (ARN) of the execution you want state machine
	// information for.
	//
	// This member is required.
	ExecutionArn *string

	noSmithyDocumentSerde
}

type DescribeStateMachineForExecutionOutput struct {

	// The Amazon States Language definition of the state machine. See Amazon States
	// Language (https://docs.aws.amazon.com/step-functions/latest/dg/concepts-amazon-states-language.html)
	// .
	//
	// This member is required.
	Definition *string

	// The name of the state machine associated with the execution.
	//
	// This member is required.
	Name *string

	// The Amazon Resource Name (ARN) of the IAM role of the State Machine for the
	// execution.
	//
	// This member is required.
	RoleArn *string

	// The Amazon Resource Name (ARN) of the state machine associated with the
	// execution.
	//
	// This member is required.
	StateMachineArn *string

	// The date and time the state machine associated with an execution was updated.
	// For a newly created state machine, this is the creation date.
	//
	// This member is required.
	UpdateDate *time.Time

	// A user-defined or an auto-generated string that identifies a Map state. This
	// ﬁeld is returned only if the executionArn is a child workflow execution that
	// was started by a Distributed Map state.
	Label *string

	// The LoggingConfiguration data type is used to set CloudWatch Logs options.
	LoggingConfiguration *types.LoggingConfiguration

	// The Amazon Resource Name (ARN) of the Map Run that started the child workflow
	// execution. This field is returned only if the executionArn is a child workflow
	// execution that was started by a Distributed Map state.
	MapRunArn *string

	// Selects whether X-Ray tracing is enabled.
	TracingConfiguration *types.TracingConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeStateMachineForExecutionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpDescribeStateMachineForExecution{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpDescribeStateMachineForExecution{}, middleware.After)
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
	if err = addOpDescribeStateMachineForExecutionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeStateMachineForExecution(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeStateMachineForExecution(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "states",
		OperationName: "DescribeStateMachineForExecution",
	}
}
