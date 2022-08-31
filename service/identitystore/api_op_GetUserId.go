// Code generated by smithy-go-codegen DO NOT EDIT.

package identitystore

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/identitystore/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the UserId in an identity store.
func (c *Client) GetUserId(ctx context.Context, params *GetUserIdInput, optFns ...func(*Options)) (*GetUserIdOutput, error) {
	if params == nil {
		params = &GetUserIdInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetUserId", params, optFns, c.addOperationGetUserIdMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetUserIdOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetUserIdInput struct {

	// Any unique attribute associated with a user that is not the UserId.
	//
	// This member is required.
	AlternateIdentifier types.AlternateIdentifier

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	noSmithyDocumentSerde
}

type GetUserIdOutput struct {

	// The globally unique identifier for the identity store.
	//
	// This member is required.
	IdentityStoreId *string

	// The identifier for a user in the identity store.
	//
	// This member is required.
	UserId *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetUserIdMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetUserId{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetUserId{}, middleware.After)
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
	if err = addOpGetUserIdValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetUserId(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetUserId(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "identitystore",
		OperationName: "GetUserId",
	}
}
