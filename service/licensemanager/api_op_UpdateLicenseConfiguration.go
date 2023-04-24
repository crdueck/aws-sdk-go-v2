// Code generated by smithy-go-codegen DO NOT EDIT.

package licensemanager

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/licensemanager/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies the attributes of an existing license configuration.
func (c *Client) UpdateLicenseConfiguration(ctx context.Context, params *UpdateLicenseConfigurationInput, optFns ...func(*Options)) (*UpdateLicenseConfigurationOutput, error) {
	if params == nil {
		params = &UpdateLicenseConfigurationInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdateLicenseConfiguration", params, optFns, c.addOperationUpdateLicenseConfigurationMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdateLicenseConfigurationOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdateLicenseConfigurationInput struct {

	// Amazon Resource Name (ARN) of the license configuration.
	//
	// This member is required.
	LicenseConfigurationArn *string

	// New description of the license configuration.
	Description *string

	// When true, disassociates a resource when software is uninstalled.
	DisassociateWhenNotFound *bool

	// New status of the license configuration.
	LicenseConfigurationStatus types.LicenseConfigurationStatus

	// New number of licenses managed by the license configuration.
	LicenseCount *int64

	// New hard limit of the number of available licenses.
	LicenseCountHardLimit *bool

	// New license rule. The only rule that you can add after you create a license
	// configuration is licenseAffinityToHost.
	LicenseRules []string

	// New name of the license configuration.
	Name *string

	// New product information.
	ProductInformationList []types.ProductInformation

	noSmithyDocumentSerde
}

type UpdateLicenseConfigurationOutput struct {
	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdateLicenseConfigurationMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpUpdateLicenseConfiguration{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpUpdateLicenseConfiguration{}, middleware.After)
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
	if err = addOpUpdateLicenseConfigurationValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdateLicenseConfiguration(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdateLicenseConfiguration(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "license-manager",
		OperationName: "UpdateLicenseConfiguration",
	}
}
