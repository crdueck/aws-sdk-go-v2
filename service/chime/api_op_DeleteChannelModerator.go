// Code generated by smithy-go-codegen DO NOT EDIT.

package chime

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deletes a channel moderator. The x-amz-chime-bearer request header is
// mandatory. Use the AppInstanceUserArn of the user that makes the API call as
// the value in the header.
func (c *Client) DeleteChannelModerator(ctx context.Context, params *DeleteChannelModeratorInput, optFns ...func(*Options)) (*DeleteChannelModeratorOutput, error) {
	if params == nil {
		params = &DeleteChannelModeratorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DeleteChannelModerator", params, optFns, c.addOperationDeleteChannelModeratorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DeleteChannelModeratorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DeleteChannelModeratorInput struct {

	// The ARN of the channel.
	//
	// This member is required.
	ChannelArn *string

	// The ARN of the moderator being deleted.
	//
	// This member is required.
	ChannelModeratorArn *string

	// The AppInstanceUserArn of the user that makes the API call.
	ChimeBearer *string

	noSmithyDocumentSerde
}

type DeleteChannelModeratorOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDeleteChannelModeratorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDeleteChannelModerator{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDeleteChannelModerator{}, middleware.After)
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
	if err = addEndpointPrefix_opDeleteChannelModeratorMiddleware(stack); err != nil {
		return err
	}
	if err = addOpDeleteChannelModeratorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDeleteChannelModerator(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opDeleteChannelModeratorMiddleware struct {
}

func (*endpointPrefix_opDeleteChannelModeratorMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opDeleteChannelModeratorMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "messaging-" + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opDeleteChannelModeratorMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opDeleteChannelModeratorMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opDeleteChannelModerator(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "chime",
		OperationName: "DeleteChannelModerator",
	}
}
