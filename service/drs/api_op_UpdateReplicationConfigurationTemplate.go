// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/drs/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates a ReplicationConfigurationTemplate by ID.
func (c *Client) UpdateReplicationConfigurationTemplate(ctx context.Context, params *UpdateReplicationConfigurationTemplateInput, optFns ...func(*Options)) (*UpdateReplicationConfigurationTemplateOutput, error) {
	if params == nil {
		params = &UpdateReplicationConfigurationTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateReplicationConfigurationTemplate", params, optFns, c.addOperationUpdateReplicationConfigurationTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateReplicationConfigurationTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateReplicationConfigurationTemplateInput struct {

	// The Replication Configuration Template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// The Replication Configuration Template ARN.
	Arn *string

	// Whether to associate the default Elastic Disaster Recovery Security group with
	// the Replication Configuration Template.
	AssociateDefaultSecurityGroup *bool

	// Whether to allow the AWS replication agent to automatically replicate newly
	// added disks.
	AutoReplicateNewDisks *bool

	// Configure bandwidth throttling for the outbound data transfer rate of the
	// Source Server in Mbps.
	BandwidthThrottling int64

	// Whether to create a Public IP for the Recovery Instance by default.
	CreatePublicIP *bool

	// The data plane routing mechanism that will be used for replication.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// The Staging Disk EBS volume type to be used during replication.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// The type of EBS encryption to be used during replication.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// The ARN of the EBS encryption key to be used during replication.
	EbsEncryptionKeyArn *string

	// The Point in time (PIT) policy to manage snapshots taken during replication.
	PitPolicy []types.PITPolicyRule

	// The instance type to be used for the replication server.
	ReplicationServerInstanceType *string

	// The security group IDs that will be used by the replication server.
	ReplicationServersSecurityGroupsIDs []string

	// The subnet to be used by the replication staging area.
	StagingAreaSubnetId *string

	// A set of tags to be associated with all resources created in the replication
	// staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	StagingAreaTags map[string]string

	// Whether to use a dedicated Replication Server in the replication staging area.
	UseDedicatedReplicationServer *bool

	noSmithyDocumentSerde
}

type UpdateReplicationConfigurationTemplateOutput struct {

	// The Replication Configuration Template ID.
	//
	// This member is required.
	ReplicationConfigurationTemplateID *string

	// The Replication Configuration Template ARN.
	Arn *string

	// Whether to associate the default Elastic Disaster Recovery Security group with
	// the Replication Configuration Template.
	AssociateDefaultSecurityGroup *bool

	// Whether to allow the AWS replication agent to automatically replicate newly
	// added disks.
	AutoReplicateNewDisks *bool

	// Configure bandwidth throttling for the outbound data transfer rate of the
	// Source Server in Mbps.
	BandwidthThrottling int64

	// Whether to create a Public IP for the Recovery Instance by default.
	CreatePublicIP *bool

	// The data plane routing mechanism that will be used for replication.
	DataPlaneRouting types.ReplicationConfigurationDataPlaneRouting

	// The Staging Disk EBS volume type to be used during replication.
	DefaultLargeStagingDiskType types.ReplicationConfigurationDefaultLargeStagingDiskType

	// The type of EBS encryption to be used during replication.
	EbsEncryption types.ReplicationConfigurationEbsEncryption

	// The ARN of the EBS encryption key to be used during replication.
	EbsEncryptionKeyArn *string

	// The Point in time (PIT) policy to manage snapshots taken during replication.
	PitPolicy []types.PITPolicyRule

	// The instance type to be used for the replication server.
	ReplicationServerInstanceType *string

	// The security group IDs that will be used by the replication server.
	ReplicationServersSecurityGroupsIDs []string

	// The subnet to be used by the replication staging area.
	StagingAreaSubnetId *string

	// A set of tags to be associated with all resources created in the replication
	// staging area: EC2 replication server, EBS volumes, EBS snapshots, etc.
	StagingAreaTags map[string]string

	// A set of tags to be associated with the Replication Configuration Template
	// resource.
	Tags map[string]string

	// Whether to use a dedicated Replication Server in the replication staging area.
	UseDedicatedReplicationServer *bool

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateReplicationConfigurationTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdateReplicationConfigurationTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdateReplicationConfigurationTemplate{}, middleware.After)
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
	if err = addOpUpdateReplicationConfigurationTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateReplicationConfigurationTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateReplicationConfigurationTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "UpdateReplicationConfigurationTemplate",
	}
}
