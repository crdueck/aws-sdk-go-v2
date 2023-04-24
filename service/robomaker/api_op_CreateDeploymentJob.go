// Code generated by smithy-go-codegen DO NOT EDIT.

package robomaker

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/robomaker/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
	"time"
)

// Deploys a specific version of a robot application to robots in a fleet. This
// API is no longer supported and will throw an error if used. The robot
// application must have a numbered applicationVersion for consistency reasons. To
// create a new version, use CreateRobotApplicationVersion or see Creating a Robot
// Application Version (https://docs.aws.amazon.com/robomaker/latest/dg/create-robot-application-version.html)
// . After 90 days, deployment jobs expire and will be deleted. They will no longer
// be accessible.
//
// Deprecated: AWS RoboMaker is unable to process this request as the support for
// the AWS RoboMaker application deployment feature has ended. For additional
// information, see https://docs.aws.amazon.com/robomaker/latest/dg/fleets.html.
func (c *Client) CreateDeploymentJob(ctx context.Context, params *CreateDeploymentJobInput, optFns ...func(*Options)) (*CreateDeploymentJobOutput, error) {
	if params == nil {
		params = &CreateDeploymentJobInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateDeploymentJob", params, optFns, c.addOperationCreateDeploymentJobMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateDeploymentJobOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateDeploymentJobInput struct {

	// Unique, case-sensitive identifier that you provide to ensure the idempotency of
	// the request.
	//
	// This member is required.
	ClientRequestToken *string

	// The deployment application configuration.
	//
	// This member is required.
	DeploymentApplicationConfigs []types.DeploymentApplicationConfig

	// The Amazon Resource Name (ARN) of the fleet to deploy.
	//
	// This member is required.
	Fleet *string

	// The requested deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// A map that contains tag keys and tag values that are attached to the deployment
	// job.
	Tags map[string]string

	noSmithyDocumentSerde
}

type CreateDeploymentJobOutput struct {

	// The Amazon Resource Name (ARN) of the deployment job.
	Arn *string

	// The time, in milliseconds since the epoch, when the fleet was created.
	CreatedAt *time.Time

	// The deployment application configuration.
	DeploymentApplicationConfigs []types.DeploymentApplicationConfig

	// The deployment configuration.
	DeploymentConfig *types.DeploymentConfig

	// The failure code of the simulation job if it failed: BadPermissionError AWS
	// Greengrass requires a service-level role permission to access other services.
	// The role must include the AWSGreengrassResourceAccessRolePolicy managed policy (https://console.aws.amazon.com/iam/home?#/policies/arn:aws:iam::aws:policy/service-role/AWSGreengrassResourceAccessRolePolicy$jsonEditor)
	// . ExtractingBundleFailure The robot application could not be extracted from the
	// bundle. FailureThresholdBreached The percentage of robots that could not be
	// updated exceeded the percentage set for the deployment.
	// GreengrassDeploymentFailed The robot application could not be deployed to the
	// robot. GreengrassGroupVersionDoesNotExist The AWS Greengrass group or version
	// associated with a robot is missing. InternalServerError An internal error has
	// occurred. Retry your request, but if the problem persists, contact us with
	// details. MissingRobotApplicationArchitecture The robot application does not have
	// a source that matches the architecture of the robot.
	// MissingRobotDeploymentResource One or more of the resources specified for the
	// robot application are missing. For example, does the robot application have the
	// correct launch package and launch file? PostLaunchFileFailure The post-launch
	// script failed. PreLaunchFileFailure The pre-launch script failed.
	// ResourceNotFound One or more deployment resources are missing. For example, do
	// robot application source bundles still exist? RobotDeploymentNoResponse There is
	// no response from the robot. It might not be powered on or connected to the
	// internet.
	FailureCode types.DeploymentJobErrorCode

	// The failure reason of the deployment job if it failed.
	FailureReason *string

	// The target fleet for the deployment job.
	Fleet *string

	// The status of the deployment job.
	Status types.DeploymentStatus

	// The list of all tags added to the deployment job.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateDeploymentJobMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateDeploymentJob{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateDeploymentJob{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateDeploymentJobMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateDeploymentJobValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateDeploymentJob(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateDeploymentJob struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateDeploymentJob) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateDeploymentJob) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateDeploymentJobInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateDeploymentJobInput ")
	}

	if input.ClientRequestToken == nil {
		t, err := m.tokenProvider.GetIdempotencyToken()
		if err != nil {
			return out, metadata, err
		}
		input.ClientRequestToken = &t
	}
	return next.HandleInitialize(ctx, in)
}
func addIdempotencyToken_opCreateDeploymentJobMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateDeploymentJob{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateDeploymentJob(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "robomaker",
		OperationName: "CreateDeploymentJob",
	}
}
