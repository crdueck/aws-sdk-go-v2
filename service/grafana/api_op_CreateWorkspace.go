// Code generated by smithy-go-codegen DO NOT EDIT.

package grafana

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/grafana/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a workspace. In a workspace, you can create Grafana dashboards and
// visualizations to analyze your metrics, logs, and traces. You don't have to
// build, package, or deploy any hardware to run the Grafana server. Don't use
// CreateWorkspace to modify an existing workspace. Instead, use UpdateWorkspace (https://docs.aws.amazon.com/grafana/latest/APIReference/API_UpdateWorkspace.html)
// .
func (c *Client) CreateWorkspace(ctx context.Context, params *CreateWorkspaceInput, optFns ...func(*Options)) (*CreateWorkspaceOutput, error) {
	if params == nil {
		params = &CreateWorkspaceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateWorkspace", params, optFns, c.addOperationCreateWorkspaceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateWorkspaceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateWorkspaceInput struct {

	// Specifies whether the workspace can access Amazon Web Services resources in
	// this Amazon Web Services account only, or whether it can also access Amazon Web
	// Services resources in other accounts in the same organization. If you specify
	// ORGANIZATION , you must specify which organizational units the workspace can
	// access in the workspaceOrganizationalUnits parameter.
	//
	// This member is required.
	AccountAccessType types.AccountAccessType

	// Specifies whether this workspace uses SAML 2.0, IAM Identity Center (successor
	// to Single Sign-On), or both to authenticate users for using the Grafana console
	// within a workspace. For more information, see User authentication in Amazon
	// Managed Grafana (https://docs.aws.amazon.com/grafana/latest/userguide/authentication-in-AMG.html)
	// .
	//
	// This member is required.
	AuthenticationProviders []types.AuthenticationProviderTypes

	// When creating a workspace through the Amazon Web Services API, CLI or Amazon
	// Web Services CloudFormation, you must manage IAM roles and provision the
	// permissions that the workspace needs to use Amazon Web Services data sources and
	// notification channels. You must also specify a workspaceRoleArn for a role that
	// you will manage for the workspace to use when accessing those datasources and
	// notification channels. The ability for Amazon Managed Grafana to create and
	// update IAM roles on behalf of the user is supported only in the Amazon Managed
	// Grafana console, where this value may be set to SERVICE_MANAGED . Use only the
	// CUSTOMER_MANAGED permission type when creating a workspace with the API, CLI or
	// Amazon Web Services CloudFormation. For more information, see Amazon Managed
	// Grafana permissions and policies for Amazon Web Services data sources and
	// notification channels (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-manage-permissions.html)
	// .
	//
	// This member is required.
	PermissionType types.PermissionType

	// A unique, case-sensitive, user-provided identifier to ensure the idempotency of
	// the request.
	ClientToken *string

	// The configuration string for the workspace that you create. For more
	// information about the format and configuration options available, see Working
	// in your Grafana workspace (https://docs.aws.amazon.com/grafana/latest/userguide/AMG-configure-workspace.html)
	// .
	//
	// This value conforms to the media type: application/json
	Configuration *string

	// Configuration for network access to your workspace. When this is configured,
	// only listed IP addresses and VPC endpoints will be able to access your
	// workspace. Standard Grafana authentication and authorization will still be
	// required. If this is not configured, or is removed, then all IP addresses and
	// VPC endpoints will be allowed. Standard Grafana authentication and authorization
	// will still be required.
	NetworkAccessControl *types.NetworkAccessConfiguration

	// The name of an IAM role that already exists to use with Organizations to access
	// Amazon Web Services data sources and notification channels in other accounts in
	// an organization.
	OrganizationRoleName *string

	// The name of the CloudFormation stack set to use to generate IAM roles to be
	// used for this workspace.
	StackSetName *string

	// The list of tags associated with the workspace.
	Tags map[string]string

	// The configuration settings for an Amazon VPC that contains data sources for
	// your Grafana workspace to connect to.
	VpcConfiguration *types.VpcConfiguration

	// This parameter is for internal use only, and should not be used.
	WorkspaceDataSources []types.DataSourceType

	// A description for the workspace. This is used only to help you identify this
	// workspace. Pattern: ^[\\p{L}\\p{Z}\\p{N}\\p{P}]{0,2048}$
	WorkspaceDescription *string

	// The name for the workspace. It does not have to be unique.
	WorkspaceName *string

	// Specify the Amazon Web Services notification channels that you plan to use in
	// this workspace. Specifying these data sources here enables Amazon Managed
	// Grafana to create IAM roles and permissions that allow Amazon Managed Grafana to
	// use these channels.
	WorkspaceNotificationDestinations []types.NotificationDestinationType

	// Specifies the organizational units that this workspace is allowed to use data
	// sources from, if this workspace is in an account that is part of an
	// organization.
	WorkspaceOrganizationalUnits []string

	// Specified the IAM role that grants permissions to the Amazon Web Services
	// resources that the workspace will view data from, including both data sources
	// and notification channels. You are responsible for managing the permissions for
	// this role as new data sources or notification channels are added.
	WorkspaceRoleArn *string

	noSmithyDocumentSerde
}

type CreateWorkspaceOutput struct {

	// A structure containing data about the workspace that was created.
	//
	// This member is required.
	Workspace *types.WorkspaceDescription

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateWorkspaceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateWorkspace{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateWorkspace{}, middleware.After)
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
	if err = addIdempotencyToken_opCreateWorkspaceMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpCreateWorkspaceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateWorkspace(options.Region), middleware.Before); err != nil {
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

type idempotencyToken_initializeOpCreateWorkspace struct {
	tokenProvider IdempotencyTokenProvider
}

func (*idempotencyToken_initializeOpCreateWorkspace) ID() string {
	return "OperationIdempotencyTokenAutoFill"
}

func (m *idempotencyToken_initializeOpCreateWorkspace) HandleInitialize(ctx context.Context, in middleware.InitializeInput, next middleware.InitializeHandler) (
	out middleware.InitializeOutput, metadata middleware.Metadata, err error,
) {
	if m.tokenProvider == nil {
		return next.HandleInitialize(ctx, in)
	}

	input, ok := in.Parameters.(*CreateWorkspaceInput)
	if !ok {
		return out, metadata, fmt.Errorf("expected middleware input to be of type *CreateWorkspaceInput ")
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
func addIdempotencyToken_opCreateWorkspaceMiddleware(stack *middleware.Stack, cfg Options) error {
	return stack.Initialize.Add(&idempotencyToken_initializeOpCreateWorkspace{tokenProvider: cfg.IdempotencyTokenProvider}, middleware.Before)
}

func newServiceMetadataMiddleware_opCreateWorkspace(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "grafana",
		OperationName: "CreateWorkspace",
	}
}
