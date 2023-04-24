// Code generated by smithy-go-codegen DO NOT EDIT.

package ram

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds or replaces the RAM permission for a resource type included in a resource
// share. You can have exactly one permission associated with each resource type in
// the resource share. You can add a new RAM permission only if there are currently
// no resources of that resource type currently in the resource share.
func (c *Client) AssociateResourceSharePermission(ctx context.Context, params *AssociateResourceSharePermissionInput, optFns ...func(*Options)) (*AssociateResourceSharePermissionOutput, error) {
	if params == nil {
		params = &AssociateResourceSharePermissionInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AssociateResourceSharePermission", params, optFns, c.addOperationAssociateResourceSharePermissionMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AssociateResourceSharePermissionOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AssociateResourceSharePermissionInput struct {

	// Specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the RAM permission to associate with the resource share. To find the ARN for
	// a permission, use either the ListPermissions operation or go to the Permissions
	// library (https://console.aws.amazon.com/ram/home#Permissions:) page in the RAM
	// console and then choose the name of the permission. The ARN is displayed on the
	// detail page.
	//
	// This member is required.
	PermissionArn *string

	// Specifies the Amazon Resource Name (ARN) (https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html)
	// of the resource share to which you want to add or replace permissions.
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

	// Specifies the version of the RAM permission to associate with the resource
	// share. You can specify only the version that is currently set as the default
	// version for the permission. If you also set the replace pararameter to true ,
	// then this operation updates an outdated version of the permission to the current
	// default version. You don't need to specify this parameter because the default
	// behavior is to use the version that is currently set as the default version for
	// the permission. This parameter is supported for backwards compatibility.
	PermissionVersion *int32

	// Specifies whether the specified permission should replace the existing
	// permission associated with the resource share. Use true to replace the current
	// permissions. Use false to add the permission to a resource share that currently
	// doesn't have a permission. The default value is false . A resource share can
	// have only one permission per resource type. If a resource share already has a
	// permission for the specified resource type and you don't set replace to true
	// then the operation returns an error. This helps prevent accidental overwriting
	// of a permission.
	Replace *bool

	noSmithyDocumentSerde
}

type AssociateResourceSharePermissionOutput struct {

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

func (c *Client) addOperationAssociateResourceSharePermissionMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpAssociateResourceSharePermission{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpAssociateResourceSharePermission{}, middleware.After)
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
	if err = addOpAssociateResourceSharePermissionValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAssociateResourceSharePermission(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAssociateResourceSharePermission(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ram",
		OperationName: "AssociateResourceSharePermission",
	}
}
