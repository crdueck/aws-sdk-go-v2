// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/gamelift/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves a fresh set of credentials for use when uploading a new set of game
// build files to Amazon GameLift's Amazon S3. This is done as part of the build
// creation process; see GameSession (https://docs.aws.amazon.com/gamelift/latest/apireference/API_CreateBuild.html)
// . To request new credentials, specify the build ID as returned with an initial
// CreateBuild request. If successful, a new set of credentials are returned, along
// with the S3 storage location associated with the build ID. Learn more Create a
// Build with Files in S3 (https://docs.aws.amazon.com/gamelift/latest/developerguide/gamelift-build-cli-uploading.html#gamelift-build-cli-uploading-create-build)
// All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) RequestUploadCredentials(ctx context.Context, params *RequestUploadCredentialsInput, optFns ...func(*Options)) (*RequestUploadCredentialsOutput, error) {
	if params == nil {
		params = &RequestUploadCredentialsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "RequestUploadCredentials", params, optFns, c.addOperationRequestUploadCredentialsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*RequestUploadCredentialsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type RequestUploadCredentialsInput struct {

	// A unique identifier for the build to get credentials for. You can use either
	// the build ID or ARN value.
	//
	// This member is required.
	BuildId *string

	noSmithyDocumentSerde
}

type RequestUploadCredentialsOutput struct {

	// Amazon S3 path and key, identifying where the game build files are stored.
	StorageLocation *types.S3Location

	// Amazon Web Services credentials required when uploading a game build to the
	// storage location. These credentials have a limited lifespan and are valid only
	// for the build they were issued for.
	UploadCredentials *types.AwsCredentials

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationRequestUploadCredentialsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpRequestUploadCredentials{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpRequestUploadCredentials{}, middleware.After)
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
	if err = addOpRequestUploadCredentialsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opRequestUploadCredentials(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opRequestUploadCredentials(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "RequestUploadCredentials",
	}
}
