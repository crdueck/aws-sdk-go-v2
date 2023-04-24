// Code generated by smithy-go-codegen DO NOT EDIT.

package gamelift

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the location of stored game session logs for a specified game
// session. When a game session is terminated, Amazon GameLift automatically stores
// the logs in Amazon S3 and retains them for 14 days. Use this URL to download the
// logs. See the Amazon Web Services Service Limits (https://docs.aws.amazon.com/general/latest/gr/aws_service_limits.html#limits_gamelift)
// page for maximum log file sizes. Log files that exceed this limit are not saved.
// All APIs by task (https://docs.aws.amazon.com/gamelift/latest/developerguide/reference-awssdk.html#reference-awssdk-resources-fleets)
func (c *Client) GetGameSessionLogUrl(ctx context.Context, params *GetGameSessionLogUrlInput, optFns ...func(*Options)) (*GetGameSessionLogUrlOutput, error) {
	if params == nil {
		params = &GetGameSessionLogUrlInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetGameSessionLogUrl", params, optFns, c.addOperationGetGameSessionLogUrlMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetGameSessionLogUrlOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetGameSessionLogUrlInput struct {

	// A unique identifier for the game session to get logs for.
	//
	// This member is required.
	GameSessionId *string

	noSmithyDocumentSerde
}

type GetGameSessionLogUrlOutput struct {

	// Location of the requested game session logs, available for download. This URL
	// is valid for 15 minutes, after which S3 will reject any download request using
	// this URL. You can request a new URL any time within the 14-day period that the
	// logs are retained.
	PreSignedUrl *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetGameSessionLogUrlMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpGetGameSessionLogUrl{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpGetGameSessionLogUrl{}, middleware.After)
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
	if err = addOpGetGameSessionLogUrlValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetGameSessionLogUrl(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetGameSessionLogUrl(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "gamelift",
		OperationName: "GetGameSessionLogUrl",
	}
}
