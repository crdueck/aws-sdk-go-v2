// Code generated by smithy-go-codegen DO NOT EDIT.

package vpclattice

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Enables access logs to be sent to Amazon CloudWatch, Amazon S3, and Amazon
// Kinesis Data Firehose. The service network owner can use the access logs to
// audit the services in the network. The service network owner will only see
// access logs from clients and services that are associated with their service
// network. Access log entries represent traffic originated from VPCs associated
// with that network. For more information, see Access logs (https://docs.aws.amazon.com/vpc-lattice/latest/ug/monitoring-access-logs.html)
// in the Amazon VPC Lattice User Guide.
func (c *Client) CreateAccessLogSubscription(ctx context.Context, params *CreateAccessLogSubscriptionInput, optFns ...func(*Options)) (*CreateAccessLogSubscriptionOutput, error) {
	if params == nil {
		params = &CreateAccessLogSubscriptionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateAccessLogSubscription", params, optFns, c.addOperationCreateAccessLogSubscriptionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateAccessLogSubscriptionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateAccessLogSubscriptionInput struct {

	// The Amazon Resource Name (ARN) of the destination. The supported destination
	// types are CloudWatch Log groups, Kinesis Data Firehose delivery streams, and
	// Amazon S3 buckets.
	//
	// This member is required.
	DestinationArn *string

	// The ID or Amazon Resource Name (ARN) of the service network or service.
	//
	// This member is required.
	ResourceIdentifier *string

	// A unique, case-sensitive identifier that you provide to ensure the idempotency
	// of the request. If you retry a request that completed successfully using the
	// same client token and parameters, the retry succeeds without performing any
	// actions. If the parameters aren't identical, the retry fails.
	ClientToken *string

	// The tags for the access log subscription.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateAccessLogSubscriptionOutput struct {

	// The Amazon Resource Name (ARN) of the access log subscription.
	//
	// This member is required.
	Arn *string

	// The Amazon Resource Name (ARN) of the log destination.
	//
	// This member is required.
	DestinationArn *string

	// The ID of the access log subscription.
	//
	// This member is required.
	Id *string

	// The Amazon Resource Name (ARN) of the service network or service.
	//
	// This member is required.
	ResourceArn *string

	// The ID of the service network or service.
	//
	// This member is required.
	ResourceId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateAccessLogSubscriptionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateAccessLogSubscription{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateAccessLogSubscription{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateAccessLogSubscriptionMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateAccessLogSubscriptionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateAccessLogSubscription(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateAccessLogSubscription struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateAccessLogSubscription) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateAccessLogSubscription) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateAccessLogSubscriptionInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateAccessLogSubscriptionInput ")
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
func addIdempotencyToken_opCreateAccessLogSubscriptionMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateAccessLogSubscription{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateAccessLogSubscription(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "vpc-lattice",
		OperationName: "CreateAccessLogSubscription",
	}
}
