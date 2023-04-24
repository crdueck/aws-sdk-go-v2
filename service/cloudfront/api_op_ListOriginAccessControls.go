// Code generated by smithy-go-codegen DO NOT EDIT.

package cloudfront

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/cloudfront/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the list of CloudFront origin access controls in this Amazon Web Services
// account. You can optionally specify the maximum number of items to receive in
// the response. If the total number of items in the list exceeds the maximum that
// you specify, or the default maximum, the response is paginated. To get the next
// page of items, send another request that specifies the NextMarker value from
// the current response as the Marker value in the next request.
func (c *Client) ListOriginAccessControls(ctx context.Context, params *ListOriginAccessControlsInput, optFns ...func(*Options)) (*ListOriginAccessControlsOutput, error) {
	if params == nil {
		params = &ListOriginAccessControlsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ListOriginAccessControls", params, optFns, c.addOperationListOriginAccessControlsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ListOriginAccessControlsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ListOriginAccessControlsInput struct {

	// Use this field when paginating results to indicate where to begin in your list
	// of origin access controls. The response includes the items in the list that
	// occur after the marker. To get the next page of the list, set this field's value
	// to the value of NextMarker from the current page's response.
	Marker *string

	// The maximum number of origin access controls that you want in the response.
	MaxItems *int32

	noSmithyDocumentSerde
}

type ListOriginAccessControlsOutput struct {

	// A list of origin access controls.
	OriginAccessControlList *types.OriginAccessControlList

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationListOriginAccessControlsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpListOriginAccessControls{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpListOriginAccessControls{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opListOriginAccessControls(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opListOriginAccessControls(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "cloudfront",
		OperationName: "ListOriginAccessControls",
	}
}
