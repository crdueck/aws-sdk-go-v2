// Code generated by smithy-go-codegen DO NOT EDIT.

package drs

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Disconnect a Recovery Instance from Elastic Disaster Recovery. Data replication
// is stopped immediately. All AWS resources created by Elastic Disaster Recovery
// for enabling the replication of the Recovery Instance will be terminated /
// deleted within 90 minutes. If the agent on the Recovery Instance has not been
// prevented from communicating with the Elastic Disaster Recovery service, then it
// will receive a command to uninstall itself (within approximately 10 minutes).
// The following properties of the Recovery Instance will be changed immediately:
// dataReplicationInfo.dataReplicationState will be set to DISCONNECTED; The
// totalStorageBytes property for each of dataReplicationInfo.replicatedDisks will
// be set to zero; dataReplicationInfo.lagDuration and
// dataReplicationInfo.lagDuration will be nullified.
func (c *Client) DisconnectRecoveryInstance(ctx context.Context, params *DisconnectRecoveryInstanceInput, optFns ...func(*Options)) (*DisconnectRecoveryInstanceOutput, error) {
	if params == nil {
		params = &DisconnectRecoveryInstanceInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DisconnectRecoveryInstance", params, optFns, c.addOperationDisconnectRecoveryInstanceMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DisconnectRecoveryInstanceOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DisconnectRecoveryInstanceInput struct {

	// The ID of the Recovery Instance to disconnect.
	//
	// This member is required.
	RecoveryInstanceID *string

	noSmithyDocumentSerde
}

type DisconnectRecoveryInstanceOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDisconnectRecoveryInstanceMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpDisconnectRecoveryInstance{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpDisconnectRecoveryInstance{}, middleware.After)
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
	if err = addOpDisconnectRecoveryInstanceValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDisconnectRecoveryInstance(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opDisconnectRecoveryInstance(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "drs",
		OperationName: "DisconnectRecoveryInstance",
	}
}
