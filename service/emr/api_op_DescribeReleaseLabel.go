// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/emr/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Provides EMR release label details, such as releases available the region where
// the API request is run, and the available applications for a specific EMR
// release label. Can also list EMR release versions that support a specified
// version of Spark.
func (c *Client) DescribeReleaseLabel(ctx context.Context, params *DescribeReleaseLabelInput, optFns ...func(*Options)) (*DescribeReleaseLabelOutput, error) {
	if params == nil {
		params = &DescribeReleaseLabelInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeReleaseLabel", params, optFns, c.addOperationDescribeReleaseLabelMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeReleaseLabelOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeReleaseLabelInput struct {

	// Reserved for future use. Currently set to null.
	MaxResults *int32

	// The pagination token. Reserved for future use. Currently set to null.
	NextToken *string

	// The target release label to be described.
	ReleaseLabel *string

	noSmithyDocumentSerde
}

type DescribeReleaseLabelOutput struct {

	// The list of applications available for the target release label. Name is the
	// name of the application. Version is the concise version of the application.
	Applications []types.SimplifiedApplication

	// The list of available Amazon Linux release versions for an Amazon EMR release.
	// Contains a Label field that is formatted as shown in Amazon Linux 2 Release
	// Notes  (https://docs.aws.amazon.com/AL2/latest/relnotes/relnotes-al2.html) . For
	// example, 2.0.20220218.1 (https://docs.aws.amazon.com/AL2/latest/relnotes/relnotes-20220218.html)
	// .
	AvailableOSReleases []types.OSRelease

	// The pagination token. Reserved for future use. Currently set to null.
	NextToken *string

	// The target release label described in the response.
	ReleaseLabel *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeReleaseLabelMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDescribeReleaseLabel{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDescribeReleaseLabel{}, middleware.After)
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
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeReleaseLabel(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeReleaseLabel(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "elasticmapreduce",
		OperationName: "DescribeReleaseLabel",
	}
}
