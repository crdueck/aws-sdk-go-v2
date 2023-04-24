// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemaker

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This action batch describes a list of versioned model packages
func (c *Client) BatchDescribeModelPackage(ctx context.Context, params *BatchDescribeModelPackageInput, optFns ...func(*Options)) (*BatchDescribeModelPackageOutput, error) {
	if params == nil {
		params = &BatchDescribeModelPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "BatchDescribeModelPackage", params, optFns, c.addOperationBatchDescribeModelPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*BatchDescribeModelPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type BatchDescribeModelPackageInput struct {

	// The list of Amazon Resource Name (ARN) of the model package groups.
	//
	// This member is required.
	ModelPackageArnList []string

	noSmithyDocumentSerde
}

type BatchDescribeModelPackageOutput struct {

	// A map of the resource and BatchDescribeModelPackageError objects reporting the
	// error associated with describing the model package.
	BatchDescribeModelPackageErrorMap map[string]types.BatchDescribeModelPackageError

	// The summaries for the model package versions
	ModelPackageSummaries map[string]types.BatchDescribeModelPackageSummary

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationBatchDescribeModelPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpBatchDescribeModelPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpBatchDescribeModelPackage{}, middleware.After)
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
	if err = addOpBatchDescribeModelPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opBatchDescribeModelPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opBatchDescribeModelPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker",
		OperationName: "BatchDescribeModelPackage",
	}
}
