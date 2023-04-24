// Code generated by smithy-go-codegen DO NOT EDIT.

package databrew

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/databrew/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a new job to analyze a dataset and create its data profile.
func (c *Client) CreateProfileJob(ctx context.Context, params *CreateProfileJobInput, optFns ...func(*Options)) (*CreateProfileJobOutput, error) {
	if params == nil {
		params = &CreateProfileJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateProfileJob", params, optFns, c.addOperationCreateProfileJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateProfileJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateProfileJobInput struct {

	// The name of the dataset that this job is to act upon.
	//
	// This member is required.
	DatasetName *string

	// The name of the job to be created. Valid characters are alphanumeric (A-Z, a-z,
	// 0-9), hyphen (-), period (.), and space.
	//
	// This member is required.
	Name *string

	// Represents an Amazon S3 location (bucket name, bucket owner, and object key)
	// where DataBrew can read input data, or write output from a job.
	//
	// This member is required.
	OutputLocation *types.S3Location

	// The Amazon Resource Name (ARN) of the Identity and Access Management (IAM) role
	// to be assumed when DataBrew runs the job.
	//
	// This member is required.
	RoleArn *string

	// Configuration for profile jobs. Used to select columns, do evaluations, and
	// override default parameters of evaluations. When configuration is null, the
	// profile job will run with default settings.
	Configuration *types.ProfileConfiguration

	// The Amazon Resource Name (ARN) of an encryption key that is used to protect the
	// job.
	EncryptionKeyArn *string

	// The encryption mode for the job, which can be one of the following:
	//   - SSE-KMS - SSE-KMS - Server-side encryption with KMS-managed keys.
	//   - SSE-S3 - Server-side encryption with keys managed by Amazon S3.
	EncryptionMode types.EncryptionMode

	// Sample configuration for profile jobs only. Determines the number of rows on
	// which the profile job will be executed. If a JobSample value is not provided,
	// the default value will be used. The default value is CUSTOM_ROWS for the mode
	// parameter and 20000 for the size parameter.
	JobSample *types.JobSample

	// Enables or disables Amazon CloudWatch logging for the job. If logging is
	// enabled, CloudWatch writes one log stream for each job run.
	LogSubscription types.LogSubscription

	// The maximum number of nodes that DataBrew can use when the job processes data.
	MaxCapacity int32

	// The maximum number of times to retry the job after a job run fails.
	MaxRetries int32

	// Metadata tags to apply to this job.
	Tags map[string]string

	// The job's timeout in minutes. A job that attempts to run longer than this
	// timeout period ends with a status of TIMEOUT .
	Timeout int32

	// List of validation configurations that are applied to the profile job.
	ValidationConfigurations []types.ValidationConfiguration

	noSmithyDocumentSerde
}

type CreateProfileJobOutput struct {

	// The name of the job that was created.
	//
	// This member is required.
	Name *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateProfileJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateProfileJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateProfileJob{}, middleware.After)
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
	if err = addOpCreateProfileJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateProfileJob(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateProfileJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "databrew",
		OperationName: "CreateProfileJob",
	}
}
