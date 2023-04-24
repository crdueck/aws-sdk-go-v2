// Code generated by smithy-go-codegen DO NOT EDIT.

package configservice

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Deploys conformance packs across member accounts in an Amazon Web Services
// Organization. For information on how many organization conformance packs and how
// many Config rules you can have per account, see Service Limits  (https://docs.aws.amazon.com/config/latest/developerguide/configlimits.html)
// in the Config Developer Guide. Only a management account and a delegated
// administrator can call this API. When calling this API with a delegated
// administrator, you must ensure Organizations ListDelegatedAdministrator
// permissions are added. An organization can have up to 3 delegated
// administrators. This API enables organization service access for
// config-multiaccountsetup.amazonaws.com through the EnableAWSServiceAccess
// action and creates a service-linked role
// AWSServiceRoleForConfigMultiAccountSetup in the management or delegated
// administrator account of your organization. The service-linked role is created
// only when the role does not exist in the caller account. To use this API with
// delegated administrator, register a delegated administrator by calling Amazon
// Web Services Organization register-delegate-admin for
// config-multiaccountsetup.amazonaws.com . Prerequisite: Ensure you call
// EnableAllFeatures API to enable all features in an organization. You must
// specify either the TemplateS3Uri or the TemplateBody parameter, but not both.
// If you provide both Config uses the TemplateS3Uri parameter and ignores the
// TemplateBody parameter. Config sets the state of a conformance pack to
// CREATE_IN_PROGRESS and UPDATE_IN_PROGRESS until the conformance pack is created
// or updated. You cannot update a conformance pack while it is in this state.
func (c *Client) PutOrganizationConformancePack(ctx context.Context, params *PutOrganizationConformancePackInput, optFns ...func(*Options)) (*PutOrganizationConformancePackOutput, error) {
	if params == nil {
		params = &PutOrganizationConformancePackInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "PutOrganizationConformancePack", params, optFns, c.addOperationPutOrganizationConformancePackMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*PutOrganizationConformancePackOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type PutOrganizationConformancePackInput struct {

	// Name of the organization conformance pack you want to create.
	//
	// This member is required.
	OrganizationConformancePackName *string

	// A list of ConformancePackInputParameter objects.
	ConformancePackInputParameters []types.ConformancePackInputParameter

	// The name of the Amazon S3 bucket where Config stores conformance pack
	// templates. This field is optional. If used, it must be prefixed with
	// awsconfigconforms .
	DeliveryS3Bucket *string

	// The prefix for the Amazon S3 bucket. This field is optional.
	DeliveryS3KeyPrefix *string

	// A list of Amazon Web Services accounts to be excluded from an organization
	// conformance pack while deploying a conformance pack.
	ExcludedAccounts []string

	// A string containing full conformance pack template body. Structure containing
	// the template body with a minimum length of 1 byte and a maximum length of 51,200
	// bytes.
	TemplateBody *string

	// Location of file containing the template body. The uri must point to the
	// conformance pack template (max size: 300 KB). You must have access to read
	// Amazon S3 bucket.
	TemplateS3Uri *string

	noSmithyDocumentSerde
}

type PutOrganizationConformancePackOutput struct {

	// ARN of the organization conformance pack.
	OrganizationConformancePackArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationPutOrganizationConformancePackMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpPutOrganizationConformancePack{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpPutOrganizationConformancePack{}, middleware.After)
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
	if err = addOpPutOrganizationConformancePackValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opPutOrganizationConformancePack(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opPutOrganizationConformancePack(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "config",
		OperationName: "PutOrganizationConformancePack",
	}
}
