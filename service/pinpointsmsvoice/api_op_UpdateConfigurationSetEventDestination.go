// Code generated by smithy-go-codegen DO NOT EDIT.

package pinpointsmsvoice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Update an event destination in a configuration set. An event destination is a
// location that you publish information about your voice calls to. For example,
// you can log an event to an Amazon CloudWatch destination when a call fails.
func (c *Client) UpdateConfigurationSetEventDestination(ctx context.Context, params *UpdateConfigurationSetEventDestinationInput, optFns ...func(*Options)) (*UpdateConfigurationSetEventDestinationOutput, error) {
	if params == nil {
		params = &UpdateConfigurationSetEventDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateConfigurationSetEventDestination", params, optFns, c.addOperationUpdateConfigurationSetEventDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateConfigurationSetEventDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// UpdateConfigurationSetEventDestinationRequest
type UpdateConfigurationSetEventDestinationInput struct {

	// ConfigurationSetName
	//
	// This member is required.
	ConfigurationSetName *string

	// EventDestinationName
	//
	// This member is required.
	EventDestinationName *string

	// An object that defines a single event destination.
	EventDestination *types.EventDestinationDefinition

	noSmithyDocumentSerde
}

// An empty object that indicates that the event destination was updated
// successfully.
type UpdateConfigurationSetEventDestinationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateConfigurationSetEventDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateConfigurationSetEventDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateConfigurationSetEventDestination{}, middleware.After)
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
	if err = addOpUpdateConfigurationSetEventDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateConfigurationSetEventDestination(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateConfigurationSetEventDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sms-voice",
		OperationName: "UpdateConfigurationSetEventDestination",
	}
}
