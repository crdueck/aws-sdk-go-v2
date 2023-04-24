// Code generated by smithy-go-codegen DO NOT EDIT.

package iotroborunner

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/iotroborunner/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Grants permission to create a destination
func (c *Client) CreateDestination(ctx context.Context, params *CreateDestinationInput, optFns ...func(*Options)) (*CreateDestinationOutput, error) {
	if params == nil {
		params = &CreateDestinationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDestination", params, optFns, c.addOperationCreateDestinationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDestinationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDestinationInput struct {

	// Human friendly name of the resource.
	//
	// This member is required.
	Name *string

	// Site ARN.
	//
	// This member is required.
	Site *string

	// JSON document containing additional fixed properties regarding the destination
	AdditionalFixedProperties *string

	// Token used for detecting replayed requests. Replayed requests will not be
	// performed multiple times.
	ClientToken *string

	// The state of the destination. Default used if not specified.
	State types.DestinationState

	noSmithyDocumentSerde
}

type CreateDestinationOutput struct {

	// Destination ARN.
	//
	// This member is required.
	Arn *string

	// Timestamp at which the resource was created.
	//
	// This member is required.
	CreatedAt *time.Time

	// Filters access by the destination's identifier
	//
	// This member is required.
	Id *string

	// State of the destination.
	//
	// This member is required.
	State types.DestinationState

	// Timestamp at which the resource was last updated.
	//
	// This member is required.
	UpdatedAt *time.Time

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDestinationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDestination{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDestination{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDestinationMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDestinationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDestination(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDestination struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDestination) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDestination) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDestinationInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDestinationInput ")
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
func addIdempotencyToken_opCreateDestinationMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDestination{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDestination(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "iotroborunner",
		OperationName: "CreateDestination",
	}
}
