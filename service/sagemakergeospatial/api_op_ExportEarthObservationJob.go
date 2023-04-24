// Code generated by smithy-go-codegen DO NOT EDIT.

package sagemakergeospatial

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/sagemakergeospatial/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Use this operation to export results of an Earth Observation job and optionally
// source images used as input to the EOJ to an Amazon S3 location.
func (c *Client) ExportEarthObservationJob(ctx context.Context, params *ExportEarthObservationJobInput, optFns ...func(*Options)) (*ExportEarthObservationJobOutput, error) {
	if params == nil {
		params = &ExportEarthObservationJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ExportEarthObservationJob", params, optFns, c.addOperationExportEarthObservationJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ExportEarthObservationJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ExportEarthObservationJobInput struct {

	// The input Amazon Resource Name (ARN) of the Earth Observation job being
	// exported.
	//
	// This member is required.
	Arn *string

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the job.
	//
	// This member is required.
	ExecutionRoleArn *string

	// An object containing information about the output file.
	//
	// This member is required.
	OutputConfig *types.OutputConfigInput

	// A unique token that guarantees that the call to this API is idempotent.
	ClientToken *string

	// The source images provided to the Earth Observation job being exported.
	ExportSourceImages *bool

	noSmithyDocumentSerde
}

type ExportEarthObservationJobOutput struct {

	// The output Amazon Resource Name (ARN) of the Earth Observation job being
	// exported.
	//
	// This member is required.
	Arn *string

	// The creation time.
	//
	// This member is required.
	CreationTime *time.Time

	// The Amazon Resource Name (ARN) of the IAM role that you specified for the job.
	//
	// This member is required.
	ExecutionRoleArn *string

	// The status of the results of the Earth Observation job being exported.
	//
	// This member is required.
	ExportStatus types.EarthObservationJobExportStatus

	// An object containing information about the output file.
	//
	// This member is required.
	OutputConfig *types.OutputConfigInput

	// The source images provided to the Earth Observation job being exported.
	ExportSourceImages *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationExportEarthObservationJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpExportEarthObservationJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpExportEarthObservationJob{}, middleware.After)
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
	if err = addIdempotencyToken_opExportEarthObservationJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpExportEarthObservationJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opExportEarthObservationJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpExportEarthObservationJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpExportEarthObservationJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpExportEarthObservationJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*ExportEarthObservationJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *ExportEarthObservationJobInput ")
	}

	if input.ClientToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opExportEarthObservationJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpExportEarthObservationJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opExportEarthObservationJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "sagemaker-geospatial",
		OperationName: "ExportEarthObservationJob",
	}
}
