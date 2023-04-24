// Code generated by smithy-go-codegen DO NOT EDIT.

package ivschat

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ivschat/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Updates a room’s configuration.
func (c *Client) UpdateRoom(ctx context.Context, params *UpdateRoomInput, optFns ...func(*Options)) (*UpdateRoomOutput, error) {
	if params == nil {
		params = &UpdateRoomInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateRoom", params, optFns, c.addOperationUpdateRoomMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateRoomOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateRoomInput struct {

	// Identifier of the room to be updated. Currently this must be an ARN.
	//
	// This member is required.
	Identifier *string

	// Array of logging-configuration identifiers attached to the room.
	LoggingConfigurationIdentifiers []string

	// The maximum number of characters in a single message. Messages are expected to
	// be UTF-8 encoded and this limit applies specifically to rune/code-point count,
	// not number of bytes. Default: 500.
	MaximumMessageLength int32

	// Maximum number of messages per second that can be sent to the room (by all
	// clients). Default: 10.
	MaximumMessageRatePerSecond int32

	// Configuration information for optional review of messages. Specify an empty uri
	// string to disassociate a message review handler from the specified room.
	MessageReviewHandler *types.MessageReviewHandler

	// Room name. The value does not need to be unique.
	Name *string

	noSmithyDocumentSerde
}

type UpdateRoomOutput struct {

	// Room ARN, from the request (if identifier was an ARN).
	Arn *string

	// Time when the room was created. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	CreateTime *time.Time

	// Room ID, generated by the system. This is a relative identifier, the part of
	// the ARN that uniquely identifies the room.
	Id *string

	// Array of logging configurations attached to the room, from the request (if
	// specified).
	LoggingConfigurationIdentifiers []string

	// Maximum number of characters in a single message, from the request (if
	// specified).
	MaximumMessageLength int32

	// Maximum number of messages per second that can be sent to the room (by all
	// clients), from the request (if specified).
	MaximumMessageRatePerSecond int32

	// Configuration information for optional review of messages.
	MessageReviewHandler *types.MessageReviewHandler

	// Room name, from the request (if specified).
	Name *string

	// Tags attached to the resource. Array of maps, each of the form string:string
	// (key:value) .
	Tags map[string]string

	// Time of the room’s last update. This is an ISO 8601 timestamp; note that this
	// is returned as a string.
	UpdateTime *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateRoomMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateRoom{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateRoom{}, middleware.After)
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
	if err = addOpUpdateRoomValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateRoom(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateRoom(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ivschat",
		OperationName: "UpdateRoom",
	}
}
