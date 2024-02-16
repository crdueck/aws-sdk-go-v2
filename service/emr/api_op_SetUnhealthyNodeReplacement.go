// Code generated by smithy-go-codegen DO NOT EDIT.

package emr

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Specify whether to enable unhealthy node replacement, which lets Amazon EMR
// gracefully replace core nodes on a cluster if any nodes become unhealthy. For
// example, a node becomes unhealthy if disk usage is above 90%. If unhealthy node
// replacement is on and TerminationProtected are off, Amazon EMR immediately
// terminates the unhealthy core nodes. To use unhealthy node replacement and
// retain unhealthy core nodes, use to turn on termination protection. In such
// cases, Amazon EMR adds the unhealthy nodes to a denylist, reducing job
// interruptions and failures. If unhealthy node replacement is on, Amazon EMR
// notifies YARN and other applications on the cluster to stop scheduling tasks
// with these nodes, moves the data, and then terminates the nodes. For more
// information, see graceful node replacement (https://docs.aws.amazon.com/emr/latest/ManagementGuide/UsingEMR_UnhealthyNodeReplacement.html)
// in the Amazon EMR Management Guide.
func (c *Client) SetUnhealthyNodeReplacement(ctx context.Context, params *SetUnhealthyNodeReplacementInput, optFns ...func(*Options)) (*SetUnhealthyNodeReplacementOutput, error) {
	if params == nil {
		params = &SetUnhealthyNodeReplacementInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "SetUnhealthyNodeReplacement", params, optFns, c.addOperationSetUnhealthyNodeReplacementMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*SetUnhealthyNodeReplacementOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type SetUnhealthyNodeReplacementInput struct {

	// The list of strings that uniquely identify the clusters for which to turn on
	// unhealthy node replacement. You can get these identifiers by running the
	// RunJobFlow or the DescribeJobFlows operations.
	//
	// This member is required.
	JobFlowIds []string

	// Indicates whether to turn on or turn off graceful unhealthy node replacement.
	//
	// This member is required.
	UnhealthyNodeReplacement *bool

	noSmithyDocumentSerde
}

type SetUnhealthyNodeReplacementOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationSetUnhealthyNodeReplacementMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpSetUnhealthyNodeReplacement{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpSetUnhealthyNodeReplacement{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "SetUnhealthyNodeReplacement"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
	}

	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
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
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpSetUnhealthyNodeReplacementValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opSetUnhealthyNodeReplacement(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opSetUnhealthyNodeReplacement(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "SetUnhealthyNodeReplacement",
	}
}
