// Code generated by smithy-go-codegen DO NOT EDIT.

package lookoutmetrics

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/lookoutmetrics/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Describes a detector. Amazon Lookout for Metrics API actions are eventually
// consistent. If you do a read operation on a resource immediately after creating
// or modifying it, use retries to allow time for the write operation to complete.
func (c *Client) DescribeAnomalyDetector(ctx context.Context, params *DescribeAnomalyDetectorInput, optFns ...func(*Options)) (*DescribeAnomalyDetectorOutput, error) {
	if params == nil {
		params = &DescribeAnomalyDetectorInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DescribeAnomalyDetector", params, optFns, c.addOperationDescribeAnomalyDetectorMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DescribeAnomalyDetectorOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DescribeAnomalyDetectorInput struct {

	// The ARN of the detector to describe.
	//
	// This member is required.
	AnomalyDetectorArn *string

	noSmithyDocumentSerde
}

type DescribeAnomalyDetectorOutput struct {

	// The ARN of the detector.
	AnomalyDetectorArn *string

	// Contains information about the detector's configuration.
	AnomalyDetectorConfig *types.AnomalyDetectorConfigSummary

	// A description of the detector.
	AnomalyDetectorDescription *string

	// The name of the detector.
	AnomalyDetectorName *string

	// The time at which the detector was created.
	CreationTime *time.Time

	// The reason that the detector failed.
	FailureReason *string

	// The process that caused the detector to fail.
	FailureType types.AnomalyDetectorFailureType

	// The ARN of the KMS key to use to encrypt your data.
	KmsKeyArn *string

	// The time at which the detector was last modified.
	LastModificationTime *time.Time

	// The status of the detector.
	Status types.AnomalyDetectorStatus

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDescribeAnomalyDetectorMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDescribeAnomalyDetector{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDescribeAnomalyDetector{}, middleware.After)
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
	if err = addOpDescribeAnomalyDetectorValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDescribeAnomalyDetector(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDescribeAnomalyDetector(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "lookoutmetrics",
		OperationName: "DescribeAnomalyDetector",
	}
}
