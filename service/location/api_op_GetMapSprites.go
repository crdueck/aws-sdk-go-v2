// Code generated by smithy-go-codegen DO NOT EDIT.

package location

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Retrieves the sprite sheet corresponding to a map resource. The sprite sheet is
// a PNG image paired with a JSON document describing the offsets of individual
// icons that will be displayed on a rendered map.
func (c *Client) GetMapSprites(ctx context.Context, params *GetMapSpritesInput, optFns ...func(*Options)) (*GetMapSpritesOutput, error) {
	if params == nil {
		params = &GetMapSpritesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetMapSprites", params, optFns, c.addOperationGetMapSpritesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetMapSpritesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetMapSpritesInput struct {

	// The name of the sprite ﬁle. Use the following ﬁle names for the sprite sheet:
	//   - sprites.png
	//   - sprites@2x.png for high pixel density displays
	// For the JSON document containing image offsets. Use the following ﬁle names:
	//   - sprites.json
	//   - sprites@2x.json for high pixel density displays
	//
	// This member is required.
	FileName *string

	// The map resource associated with the sprite ﬁle.
	//
	// This member is required.
	MapName *string

	// The optional API key (https://docs.aws.amazon.com/location/latest/developerguide/using-apikeys.html)
	// to authorize the request.
	Key *string

	noSmithyDocumentSerde
}

type GetMapSpritesOutput struct {

	// Contains the body of the sprite sheet or JSON offset ﬁle.
	Blob []byte

	// The HTTP Cache-Control directive for the value.
	CacheControl *string

	// The content type of the sprite sheet and offsets. For example, the sprite sheet
	// content type is image/png , and the sprite offset JSON document is
	// application/json .
	ContentType *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetMapSpritesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetMapSprites{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetMapSprites{}, middleware.After)
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
	if err = addEndpointPrefix_opGetMapSpritesMiddleware(stack); err != nil {
		return err
	}
	if err = addOpGetMapSpritesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetMapSprites(options.Region), middleware.Before); err != nil {
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

type endpointPrefix_opGetMapSpritesMiddleware struct {
}

func (*endpointPrefix_opGetMapSpritesMiddleware) ID() string {
	return "EndpointHostPrefix"
}

func (m *endpointPrefix_opGetMapSpritesMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if smithyhttp.GetHostnameImmutable(ctx) || smithyhttp.IsEndpointHostPrefixDisabled(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	req.URL.Host = "maps." + req.URL.Host

	return next.HandleSerialize(ctx, in)
}
func addEndpointPrefix_opGetMapSpritesMiddleware(stack *middleware.Stack) error {
	return stack.Serialize.Insert(&endpointPrefix_opGetMapSpritesMiddleware{}, `OperationSerializer`, middleware.After)
}

func newServiceMetadataMiddleware_opGetMapSprites(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "geo",
		OperationName: "GetMapSprites",
	}
}
