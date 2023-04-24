// Code generated by smithy-go-codegen DO NOT EDIT.

package servicecatalog

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/servicecatalog/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Requests updates to the properties of the specified provisioned product.
func (c *Client) UpdateProvisionedProductProperties(ctx context.Context, params *UpdateProvisionedProductPropertiesInput, optFns ...func(*Options)) (*UpdateProvisionedProductPropertiesOutput, error) {
	if params == nil {
		params = &UpdateProvisionedProductPropertiesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateProvisionedProductProperties", params, optFns, c.addOperationUpdateProvisionedProductPropertiesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateProvisionedProductPropertiesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateProvisionedProductPropertiesInput struct {

	// The idempotency token that uniquely identifies the provisioning product update
	// request.
	//
	// This member is required.
	IdempotencyToken *string

	// The identifier of the provisioned product.
	//
	// This member is required.
	ProvisionedProductId *string

	// A map that contains the provisioned product properties to be updated. The
	// LAUNCH_ROLE key accepts role ARNs. This key allows an administrator to call
	// UpdateProvisionedProductProperties to update the launch role that is associated
	// with a provisioned product. This role is used when an end user calls a
	// provisioning operation such as UpdateProvisionedProduct ,
	// TerminateProvisionedProduct , or ExecuteProvisionedProductServiceAction . Only a
	// role ARN is valid. A user ARN is invalid. The OWNER key accepts user ARNs, IAM
	// role ARNs, and STS assumed-role ARNs. The owner is the user that has permission
	// to see, update, terminate, and execute service actions in the provisioned
	// product. The administrator can change the owner of a provisioned product to
	// another IAM or STS entity within the same account. Both end user owners and
	// administrators can see ownership history of the provisioned product using the
	// ListRecordHistory API. The new owner can describe all past records for the
	// provisioned product using the DescribeRecord API. The previous owner can no
	// longer use DescribeRecord , but can still see the product's history from when he
	// was an owner using ListRecordHistory . If a provisioned product ownership is
	// assigned to an end user, they can see and perform any action through the API or
	// Service Catalog console such as update, terminate, and execute service actions.
	// If an end user provisions a product and the owner is updated to someone else,
	// they will no longer be able to see or perform any actions through API or the
	// Service Catalog console on that provisioned product.
	//
	// This member is required.
	ProvisionedProductProperties map[string]string

	// The language code.
	//   - jp - Japanese
	//   - zh - Chinese
	AcceptLanguage *string

	noSmithyDocumentSerde
}

type UpdateProvisionedProductPropertiesOutput struct {

	// The provisioned product identifier.
	ProvisionedProductId *string

	// A map that contains the properties updated.
	ProvisionedProductProperties map[string]string

	// The identifier of the record.
	RecordId *string

	// The status of the request.
	Status types.RecordStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateProvisionedProductPropertiesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateProvisionedProductProperties{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateProvisionedProductProperties{}, middleware.After)
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
	if err = addIdempotencyToken_opUpdateProvisionedProductPropertiesMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpUpdateProvisionedProductPropertiesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateProvisionedProductProperties(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpUpdateProvisionedProductProperties struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpUpdateProvisionedProductProperties) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpUpdateProvisionedProductProperties) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*UpdateProvisionedProductPropertiesInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *UpdateProvisionedProductPropertiesInput ")
	}

	if input.IdempotencyToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.IdempotencyToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opUpdateProvisionedProductPropertiesMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpUpdateProvisionedProductProperties{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opUpdateProvisionedProductProperties(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "servicecatalog",
		OperationName: "UpdateProvisionedProductProperties",
	}
}
