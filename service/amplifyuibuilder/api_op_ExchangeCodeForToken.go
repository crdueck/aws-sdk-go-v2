// Code generated by smithy-go-codegen DO NOT EDIT.

package amplifyuibuilder

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/amplifyuibuilder/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Exchanges an access code for a token.
func (c *Client) ExchangeCodeForToken(ctx context.Context, params *ExchangeCodeForTokenInput, optFns ...func(*Options)) (*ExchangeCodeForTokenOutput, error) {
	if params == nil {
		params = &ExchangeCodeForTokenInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExchangeCodeForToken", params, optFns, c.addOperationExchangeCodeForTokenMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExchangeCodeForTokenOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExchangeCodeForTokenInput struct {

	// The third-party provider for the token. The only valid value is figma .
	//
	// This member is required.
	Provider types.TokenProviders

	// Describes the configuration of the request.
	//
	// This member is required.
	Request *types.ExchangeCodeForTokenRequestBody

	noSmithyDocumentSerde
}

type ExchangeCodeForTokenOutput struct {

	// The access token.
	//
	// This member is required.
	AccessToken *string

	// The date and time when the new access token expires.
	//
	// This member is required.
	ExpiresIn *int32

	// The token to use to refresh a previously issued access token that might have
	// expired.
	//
	// This member is required.
	RefreshToken *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExchangeCodeForTokenMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExchangeCodeForToken{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExchangeCodeForToken{}, middleware.After)
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
	if err = addOpExchangeCodeForTokenValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExchangeCodeForToken(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opExchangeCodeForToken(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "amplifyuibuilder",
		OperationName: "ExchangeCodeForToken",
	}
}
