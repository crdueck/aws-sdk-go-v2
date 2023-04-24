// Code generated by smithy-go-codegen DO NOT EDIT.

package cognitosync

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Subscribes to receive notifications when a dataset is modified by another
// device.This API can only be called with temporary credentials provided by
// Cognito Identity. You cannot call this API with developer credentials.
// SubscribeToDataset The following examples have been edited for readability. POST
// / HTTP/1.1 CONTENT-TYPE: application/json X-AMZN-REQUESTID:
// 8b9932b7-201d-4418-a960-0a470e11de9f X-AMZ-TARGET:
// com.amazonaws.cognito.sync.model.AWSCognitoSyncService.SubscribeToDataset HOST:
// cognito-sync.us-east-1.amazonaws.com X-AMZ-DATE: 20141004T195350Z
// X-AMZ-SECURITY-TOKEN: AUTHORIZATION: AWS4-HMAC-SHA256 Credential=,
// SignedHeaders=content-type;content-length;host;x-amz-date;x-amz-target,
// Signature= { "Operation": "com.amazonaws.cognito.sync.model#SubscribeToDataset",
// "Service": "com.amazonaws.cognito.sync.model#AWSCognitoSyncService", "Input": {
// "IdentityPoolId": "ID_POOL_ID", "IdentityId": "IDENTITY_ID", "DatasetName":
// "Rufus", "DeviceId": "5cd28fbe-dd83-47ab-9f83-19093a5fb014" } } 1.1 200 OK
// x-amzn-requestid: 8b9932b7-201d-4418-a960-0a470e11de9f date: Sat, 04 Oct 2014
// 19:53:50 GMT content-type: application/json content-length: 99 { "Output": {
// "__type": "com.amazonaws.cognito.sync.model#SubscribeToDatasetResponse" },
// "Version": "1.0" }
func (c *Client) SubscribeToDataset(ctx context.Context, params *SubscribeToDatasetInput, optFns ...func(*Options)) (*SubscribeToDatasetOutput, error) {
	if params == nil {
		params = &SubscribeToDatasetInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SubscribeToDataset", params, optFns, c.addOperationSubscribeToDatasetMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SubscribeToDatasetOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// A request to SubscribeToDatasetRequest.
type SubscribeToDatasetInput struct {

	// The name of the dataset to subcribe to.
	//
	// This member is required.
	DatasetName *string

	// The unique ID generated for this device by Cognito.
	//
	// This member is required.
	DeviceId *string

	// Unique ID for this identity.
	//
	// This member is required.
	IdentityId *string

	// A name-spaced GUID (for example,
	// us-east-1:23EC4050-6AEA-7089-A2DD-08002EXAMPLE) created by Amazon Cognito. The
	// ID of the pool to which the identity belongs.
	//
	// This member is required.
	IdentityPoolId *string

	noSmithyDocumentSerde
}

// Response to a SubscribeToDataset request.
type SubscribeToDatasetOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSubscribeToDatasetMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpSubscribeToDataset{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpSubscribeToDataset{}, middleware.After)
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
	if err = addOpSubscribeToDatasetValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSubscribeToDataset(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opSubscribeToDataset(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cognito-sync",
		OperationName: "SubscribeToDataset",
	}
}
