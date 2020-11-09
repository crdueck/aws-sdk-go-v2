// Code generated by smithy-go-codegen DO NOT EDIT.

package s3

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	s3cust "github.com/aws/aws-sdk-go-v2/service/s3/internal/customizations"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/awslabs/smithy-go/middleware"
	smithyhttp "github.com/awslabs/smithy-go/transport/http"
)

// Returns an inventory configuration (identified by the inventory configuration
// ID) from the bucket. To use this operation, you must have permissions to perform
// the s3:GetInventoryConfiguration action. The bucket owner has this permission by
// default and can grant this permission to others. For more information about
// permissions, see Permissions Related to Bucket Subresource Operations
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/using-with-s3-actions.html#using-with-s3-actions-related-to-bucket-subresources)
// and Managing Access Permissions to Your Amazon S3 Resources
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/s3-access-control.html). For
// information about the Amazon S3 inventory feature, see Amazon S3 Inventory
// (https://docs.aws.amazon.com/AmazonS3/latest/dev/storage-inventory.html). The
// following operations are related to GetBucketInventoryConfiguration:
//
// *
// DeleteBucketInventoryConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucketInventoryConfiguration.html)
//
// *
// ListBucketInventoryConfigurations
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_ListBucketInventoryConfigurations.html)
//
// *
// PutBucketInventoryConfiguration
// (https://docs.aws.amazon.com/AmazonS3/latest/API/API_PutBucketInventoryConfiguration.html)
func (c *Client) GetBucketInventoryConfiguration(ctx context.Context, params *GetBucketInventoryConfigurationInput, optFns ...func(*Options)) (*GetBucketInventoryConfigurationOutput, error) {
	if params == nil {
		params = &GetBucketInventoryConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetBucketInventoryConfiguration", params, optFns, addOperationGetBucketInventoryConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetBucketInventoryConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetBucketInventoryConfigurationInput struct {

	// The name of the bucket containing the inventory configuration to retrieve.
	//
	// This member is required.
	Bucket *string

	// The ID used to identify the inventory configuration.
	//
	// This member is required.
	Id *string

	// The account id of the expected bucket owner. If the bucket is owned by a
	// different account, the request will fail with an HTTP 403 (Access Denied) error.
	ExpectedBucketOwner *string
}

type GetBucketInventoryConfigurationOutput struct {

	// Specifies the inventory configuration.
	InventoryConfiguration *types.InventoryConfiguration

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata
}

func addOperationGetBucketInventoryConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestxml_serializeOpGetBucketInventoryConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestxml_deserializeOpGetBucketInventoryConfiguration{}, middleware.After)
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
	if err = awsmiddleware.AddAttemptClockSkewMiddleware(stack); err != nil {
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
	if err = addOpGetBucketInventoryConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetBucketInventoryConfiguration(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = addMetadataRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addGetBucketInventoryConfigurationUpdateEndpoint(stack, options); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = v4.AddContentSHA256HeaderMiddleware(stack); err != nil {
		return err
	}
	if err = disableAcceptEncodingGzip(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opGetBucketInventoryConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "s3",
		OperationName: "GetBucketInventoryConfiguration",
	}
}

func addGetBucketInventoryConfigurationUpdateEndpoint(stack *middleware.Stack, options Options) error {
	return s3cust.UpdateEndpoint(stack, s3cust.UpdateEndpointOptions{
		Accessor: s3cust.UpdateEndpointParameterAccessor{GetBucketFromInput: getBucketFromInput,
			SupportsAccelerate: supportAccelerate,
		},
		UsePathStyle:            options.UsePathStyle,
		UseAccelerate:           options.UseAccelerate,
		EndpointResolver:        options.EndpointResolver,
		EndpointResolverOptions: options.EndpointOptions,
		UseDualstack:            options.UseDualstack,
		UseARNRegion:            options.UseARNRegion,
	})
}
