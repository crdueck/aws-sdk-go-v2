// Code generated by smithy-go-codegen DO NOT EDIT.

package efs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/efs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Returns the descriptions of all the current mount targets, or a specific mount
// target, for a file system. When requesting all of the current mount targets, the
// order of mount targets returned in the response is unspecified. This operation
// requires permissions for the elasticfilesystem:DescribeMountTargets action, on
// either the file system ID that you specify in FileSystemId , or on the file
// system of the mount target that you specify in MountTargetId .
func (c *Client) DescribeMountTargets(ctx context.Context, params *DescribeMountTargetsInput, optFns ...func(*Options)) (*DescribeMountTargetsOutput, error) {
	if params == nil {
		params = &DescribeMountTargetsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeMountTargets", params, optFns, c.addOperationDescribeMountTargetsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeMountTargetsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeMountTargetsInput struct {

	// (Optional) The ID of the access point whose mount targets that you want to
	// list. It must be included in your request if a FileSystemId or MountTargetId is
	// not included in your request. Accepts either an access point ID or ARN as input.
	AccessPointId *string

	// (Optional) ID of the file system whose mount targets you want to list (String).
	// It must be included in your request if an AccessPointId or MountTargetId is not
	// included. Accepts either a file system ID or ARN as input.
	FileSystemId *string

	// (Optional) Opaque pagination token returned from a previous DescribeMountTargets
	// operation (String). If present, it specifies to continue the list from where the
	// previous returning call left off.
	Marker *string

	// (Optional) Maximum number of mount targets to return in the response.
	// Currently, this number is automatically set to 10, and other values are ignored.
	// The response is paginated at 100 per page if you have more than 100 mount
	// targets.
	MaxItems *int32

	// (Optional) ID of the mount target that you want to have described (String). It
	// must be included in your request if FileSystemId is not included. Accepts
	// either a mount target ID or ARN as input.
	MountTargetId *string

	noSmithyDocumentSerde
}

type DescribeMountTargetsOutput struct {

	// If the request included the Marker , the response returns that value in this
	// field.
	Marker *string

	// Returns the file system's mount targets as an array of MountTargetDescription
	// objects.
	MountTargets []types.MountTargetDescription

	// If a value is present, there are more mount targets to return. In a subsequent
	// request, you can provide Marker in your request with this value to retrieve the
	// next set of mount targets.
	NextMarker *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeMountTargetsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeMountTargets{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeMountTargets{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeMountTargets(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeMountTargets(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticfilesystem",
		OperationName: "DescribeMountTargets",
	}
}
