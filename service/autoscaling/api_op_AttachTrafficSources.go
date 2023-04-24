// Code generated by smithy-go-codegen DO NOT EDIT.

package autoscaling

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/autoscaling/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Attaches one or more traffic sources to the specified Auto Scaling group. You
// can use any of the following as traffic sources for an Auto Scaling group:
//   - Application Load Balancer
//   - Classic Load Balancer
//   - Gateway Load Balancer
//   - Network Load Balancer
//   - VPC Lattice
//
// This operation is additive and does not detach existing traffic sources from
// the Auto Scaling group. After the operation completes, use the
// DescribeTrafficSources API to return details about the state of the attachments
// between traffic sources and your Auto Scaling group. To detach a traffic source
// from the Auto Scaling group, call the DetachTrafficSources API.
func (c *Client) AttachTrafficSources(ctx context.Context, params *AttachTrafficSourcesInput, optFns ...func(*Options)) (*AttachTrafficSourcesOutput, error) {
	if params == nil {
		params = &AttachTrafficSourcesInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "AttachTrafficSources", params, optFns, c.addOperationAttachTrafficSourcesMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*AttachTrafficSourcesOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type AttachTrafficSourcesInput struct {

	// The name of the Auto Scaling group.
	//
	// This member is required.
	AutoScalingGroupName *string

	// The unique identifiers of one or more traffic sources. You can specify up to 10
	// traffic sources.
	//
	// This member is required.
	TrafficSources []types.TrafficSourceIdentifier

	noSmithyDocumentSerde
}

type AttachTrafficSourcesOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationAttachTrafficSourcesMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpAttachTrafficSources{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpAttachTrafficSources{}, middleware.After)
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
	if err = addOpAttachTrafficSourcesValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opAttachTrafficSources(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opAttachTrafficSources(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "autoscaling",
		OperationName: "AttachTrafficSources",
	}
}
