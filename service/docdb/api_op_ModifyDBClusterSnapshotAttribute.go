// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Adds an attribute and values to, or removes an attribute and values from, a
// manual cluster snapshot. To share a manual cluster snapshot with other Amazon
// Web Services accounts, specify restore as the AttributeName , and use the
// ValuesToAdd parameter to add a list of IDs of the Amazon Web Services accounts
// that are authorized to restore the manual cluster snapshot. Use the value all
// to make the manual cluster snapshot public, which means that it can be copied or
// restored by all Amazon Web Services accounts. Do not add the all value for any
// manual cluster snapshots that contain private information that you don't want
// available to all Amazon Web Services accounts. If a manual cluster snapshot is
// encrypted, it can be shared, but only by specifying a list of authorized Amazon
// Web Services account IDs for the ValuesToAdd parameter. You can't use all as a
// value for that parameter in this case.
func (c *Client) ModifyDBClusterSnapshotAttribute(ctx context.Context, params *ModifyDBClusterSnapshotAttributeInput, optFns ...func(*Options)) (*ModifyDBClusterSnapshotAttributeOutput, error) {
	if params == nil {
		params = &ModifyDBClusterSnapshotAttributeInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBClusterSnapshotAttribute", params, optFns, c.addOperationModifyDBClusterSnapshotAttributeMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterSnapshotAttributeOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBClusterSnapshotAttribute .
type ModifyDBClusterSnapshotAttributeInput struct {

	// The name of the cluster snapshot attribute to modify. To manage authorization
	// for other Amazon Web Services accounts to copy or restore a manual cluster
	// snapshot, set this value to restore .
	//
	// This member is required.
	AttributeName *string

	// The identifier for the cluster snapshot to modify the attributes for.
	//
	// This member is required.
	DBClusterSnapshotIdentifier *string

	// A list of cluster snapshot attributes to add to the attribute specified by
	// AttributeName . To authorize other Amazon Web Services accounts to copy or
	// restore a manual cluster snapshot, set this list to include one or more Amazon
	// Web Services account IDs. To make the manual cluster snapshot restorable by any
	// Amazon Web Services account, set it to all . Do not add the all value for any
	// manual cluster snapshots that contain private information that you don't want to
	// be available to all Amazon Web Services accounts.
	ValuesToAdd []string

	// A list of cluster snapshot attributes to remove from the attribute specified by
	// AttributeName . To remove authorization for other Amazon Web Services accounts
	// to copy or restore a manual cluster snapshot, set this list to include one or
	// more Amazon Web Services account identifiers. To remove authorization for any
	// Amazon Web Services account to copy or restore the cluster snapshot, set it to
	// all . If you specify all , an Amazon Web Services account whose account ID is
	// explicitly added to the restore attribute can still copy or restore a manual
	// cluster snapshot.
	ValuesToRemove []string

	noSmithyDocumentSerde
}

type ModifyDBClusterSnapshotAttributeOutput struct {

	// Detailed information about the attributes that are associated with a cluster
	// snapshot.
	DBClusterSnapshotAttributesResult *types.DBClusterSnapshotAttributesResult

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterSnapshotAttributeMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBClusterSnapshotAttribute{}, middleware.After)
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
	if err = addOpModifyDBClusterSnapshotAttributeValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDBClusterSnapshotAttribute(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rds",
		OperationName: "ModifyDBClusterSnapshotAttribute",
	}
}
