// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cognitosync/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Sets the necessary configuration for push sync.This API can only be called with
// developer credentials. You cannot call this API with the temporary user
// credentials provided by Cognito Identity. SetIdentityPoolConfiguration The
// following examples have been edited for readability. POST / HTTP/1.1
// CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// a46db021-f5dd-45d6-af5b-7069fa4a211b X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.SetIdentityPoolConfiguration
// HOST: cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T200006Z
// AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation":
// "com.amazonaws.cognito.sync.model#SetIdentityPoolConfiguration", "Service":
// "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "PushSync": { "ApplicationArns":
// ["PLATFORMARN1", "PLATFORMARN2"], "RoleArn": "ROLEARN" } } } 1.1 200 OK
// x-amzn-requestid: a46db021-f5dd-45d6-af5b-7069fa4a211b date: Sat, 04 Oct 2014
// 20:00:06 GMT content-type: application/json content-length: 332 { "Output": {
// "__type":
// "com.amazonaws.cognito.sync.model#SetIdentityPoolConfigurationResponse",
// "IdentityPoolId": "ID_POOL_ID", "PushSync": { "ApplicationArns":
// ["PLATFORMARN1", "PLATFORMARN2"], "RoleArn": "ROLEARN" } }, "Version": "1.0" }
func (c *Client) SetIdentityPoolConfiguration(ctx context.Context, params *SetIdentityPoolConfigurationInput, optFns ...func(*Options)) (*SetIdentityPoolConfigurationOutput, error) {
	if params == nil {
		params = &SetIdentityPoolConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetIdentityPoolConfiguration", params, optFns, c.addOperationSetIdentityPoolConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetIdentityPoolConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// The input for the SetIdentityPoolConfiguration operation.
type SetIdentityPoolConfigurationInput struct {

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. This
	// is the ID of the pool to modify.
	//
	// This member is required.
	IdentityPoolId *string

	// Options to apply to this identity pool for Amazon Cognito streams.
	CognitoStreams *types.CognitoStreams

	// Options to apply to this identity pool for push synchronization.
	PushSync *types.PushSync

	noSmithyDocumentSerde
}

// The output for the SetIdentityPoolConfiguration operation
type SetIdentityPoolConfigurationOutput struct {

	// Options to apply to this identity pool for Amazon Cognito streams.
	CognitoStreams *types.CognitoStreams

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito.
	IdentityPoolId *string

	// Options to apply to this identity pool for push synchronization.
	PushSync *types.PushSync

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetIdentityPoolConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSetIdentityPoolConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSetIdentityPoolConfiguration{}, middleware.After)
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
	if err = addOpSetIdentityPoolConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetIdentityPoolConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSetIdentityPoolConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "SetIdentityPoolConfiguration",
	}
}
