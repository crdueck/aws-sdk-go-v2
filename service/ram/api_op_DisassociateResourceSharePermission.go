// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Removes a managed permission from a resource share. Permission changes take
// effect immediately. You can remove a managed permission from a resource share
// only if there are currently no resources of the relevant resource type currently
// attached to the resource share.
func (c *Client) DisassociateResourceSharePermission(ctx context.Context, params *DisassociateResourceSharePermissionInput, optFns ...func(*Options)) (*DisassociateResourceSharePermissionOutput, error) {
	if params == nil {
		params = &DisassociateResourceSharePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisassociateResourceSharePermission", params, optFns, c.addOperationDisassociateResourceSharePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisassociateResourceSharePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisassociateResourceSharePermissionInput struct {

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the managed permission to disassociate from the resource share. Changes to
	// permissions take effect immediately.
	//
	// This member is required.
	PermissionArn *string

	// The Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share that you want to remove the managed permission from.
	//
	// This member is required.
	ResourceShareArn *string

	// Specifies a unique, case-sensitive identifier that you provide to ensure the
	// idempotency of the request. This lets you safely retry the request without
	// accidentally performing the same operation a second time. Passing the same value
	// to a later call to an operation requires that you also pass the same value for
	// all other parameters. We recommend that you use a UUID type of value. (https://wikipedia.org/wiki/Universally_unique_identifier)
	// . If you don't provide this value, then Amazon Web Services generates a random
	// one for you. If you retry the operation with the same ClientToken , but with
	// different parameters, the retry fails with an IdempotentParameterMismatch error.
	ClientToken *string

	noSmithyDocumentSerde
}

type DisassociateResourceSharePermissionOutput struct {

	// The idempotency identifier associated with this request. If you want to repeat
	// the same operation in an idempotent manner then you must include this value in
	// the clientToken request parameter of that later call. All other parameters must
	// also have the same values that you used in the first call.
	ClientToken *string

	// A return value of true indicates that the request succeeded. A value of false
	// indicates that the request failed.
	ReturnValue *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisassociateResourceSharePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisassociateResourceSharePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisassociateResourceSharePermission{}, middleware.After)
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
	if err = addOpDisassociateResourceSharePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisassociateResourceSharePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisassociateResourceSharePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "DisassociateResourceSharePermission",
	}
}
