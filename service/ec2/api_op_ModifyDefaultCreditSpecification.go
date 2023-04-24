// Code generated by smithy-go-codegen DO NOT EDIT.

package ec2

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the default credit option for CPU usage of burstable performance
// instances. The default credit option is set at the account level per Amazon Web
// Services Region, and is specified per instance family. All new burstable
// performance instances in the account launch using the default credit option.
// ModifyDefaultCreditSpecification is an asynchronous operation, which works at an
// Amazon Web Services Region level and modifies the credit option for each
// Availability Zone. All zones in a Region are updated within five minutes. But if
// instances are launched during this operation, they might not get the new credit
// option until the zone is updated. To verify whether the update has occurred, you
// can call GetDefaultCreditSpecification and check DefaultCreditSpecification for
// updates. For more information, see Burstable performance instances (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/burstable-performance-instances.html)
// in the Amazon EC2 User Guide.
func (c *Client) ModifyDefaultCreditSpecification(ctx context.Context, params *ModifyDefaultCreditSpecificationInput, optFns ...func(*Options)) (*ModifyDefaultCreditSpecificationOutput, error) {
	if params == nil {
		params = &ModifyDefaultCreditSpecificationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDefaultCreditSpecification", params, optFns, c.addOperationModifyDefaultCreditSpecificationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDefaultCreditSpecificationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type ModifyDefaultCreditSpecificationInput struct {

	// The credit option for CPU usage of the instance family. Valid Values: standard
	// | unlimited
	//
	// This member is required.
	CpuCredits *string

	// The instance family.
	//
	// This member is required.
	InstanceFamily types.UnlimitedSupportedInstanceFamily

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have the
	// required permissions, the error response is DryRunOperation . Otherwise, it is
	// UnauthorizedOperation .
	DryRun *bool

	noSmithyDocumentSerde
}

type ModifyDefaultCreditSpecificationOutput struct {

	// The default credit option for CPU usage of the instance family.
	InstanceFamilyCreditSpecification *types.InstanceFamilyCreditSpecification

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDefaultCreditSpecificationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsEc2query_serializeOpModifyDefaultCreditSpecification{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsEc2query_deserializeOpModifyDefaultCreditSpecification{}, middleware.After)
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
	if err = addOpModifyDefaultCreditSpecificationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDefaultCreditSpecification(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opModifyDefaultCreditSpecification(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "ec2",
		OperationName: "ModifyDefaultCreditSpecification",
	}
}
