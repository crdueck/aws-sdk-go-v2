// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitoidentityprovider

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the user interface (UI) customization information for a user pool's
// built-in app UI. You can specify app UI customization settings for a single
// client (with a specific clientId ) or for all clients (by setting the clientId
// to ALL ). If you specify ALL , the default configuration is used for every
// client that has no previously set UI customization. If you specify UI
// customization settings for a particular client, it will no longer return to the
// ALL configuration. To use this API, your user pool must have a domain associated
// with it. Otherwise, there is no place to host the app's pages, and the service
// will throw an error.
func (c *Client) SetUICustomization(ctx context.Context, params *SetUICustomizationInput, optFns ...func(*Options)) (*SetUICustomizationOutput, error) {
	if params == nil {
		params = &SetUICustomizationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUICustomization", params, optFns, c.addOperationSetUICustomizationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUICustomizationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUICustomizationInput struct {

	// The user pool ID for the user pool.
	//
	// This member is required.
	UserPoolId *string

	// The CSS values in the UI customization.
	CSS *string

	// The client ID for the client app.
	ClientId *string

	// The uploaded logo image for the UI customization.
	ImageFile []byte

	noSmithyDocumentSerde
}

type SetUICustomizationOutput struct {

	// The UI customization information.
	//
	// This member is required.
	UICustomization *types.UICustomizationType

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetUICustomizationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUICustomization{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUICustomization{}, middleware.After)
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
	if err = addOpSetUICustomizationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUICustomization(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetUICustomization(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-idp",
		OperationName: "SetUICustomization",
	}
}
