// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoicev2

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new event destination in a configuration set. An event destination is
// a location where you send message events. The event options are Amazon
// CloudWatch, Amazon Kinesis Data Firehose, or Amazon SNS. For example, when a
// message is delivered successfully, you can send information about that event to
// an event destination, or send notifications to endpoints that are subscribed to
// an Amazon SNS topic. Each configuration set can contain between 0 and 5 event
// destinations. Each event destination can contain a reference to a single
// destination, such as a CloudWatch or Kinesis Data Firehose destination.
func (c *Client) CreateEventDestination(ctx context.Context, params *CreateEventDestinationInput, optFns ...func(*Options)) (*CreateEventDestinationOutput, error) {
	if params == nil {
		params = &CreateEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateEventDestination", params, optFns, c.addOperationCreateEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateEventDestinationInput struct {

	// Either the name of the configuration set or the configuration set ARN to apply
	// event logging to. The ConfigurateSetName and ConfigurationSetArn can be found
	// using the DescribeConfigurationSets action.
	//
	// This member is required.
	ConfigurationSetName *string

	// The name that identifies the event destination.
	//
	// This member is required.
	EventDestinationName *string

	// An array of event types that determine which events to log. If "ALL" is used,
	// then Amazon Pinpoint logs every event type.
	//
	// This member is required.
	MatchingEventTypes []types.EventType

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request. If you don't specify a client token, a randomly generated token is
	// used for the request to ensure idempotency.
	ClientToken *string

	// An object that contains information about an event destination for logging to
	// Amazon CloudWatch logs.
	CloudWatchLogsDestination *types.CloudWatchLogsDestination

	// An object that contains information about an event destination for logging to
	// Amazon Kinesis Data Firehose.
	KinesisFirehoseDestination *types.KinesisFirehoseDestination

	// An object that contains information about an event destination for logging to
	// Amazon SNS.
	SnsDestination *types.SnsDestination

	noSmithyDocumentSerde
}

type CreateEventDestinationOutput struct {

	// The ARN of the configuration set.
	ConfigurationSetArn *string

	// The name of the configuration set.
	ConfigurationSetName *string

	// The details of the destination where events are logged.
	EventDestination *types.EventDestination

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson10_serializeOpCreateEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson10_deserializeOpCreateEventDestination{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateEventDestinationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateEventDestination(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateEventDestination struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateEventDestination) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateEventDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateEventDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateEventDestinationInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateEventDestinationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateEventDestination{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "CreateEventDestination",
	}
}
