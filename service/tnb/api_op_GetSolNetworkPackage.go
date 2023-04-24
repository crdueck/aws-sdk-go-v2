// Code generated by smithy-go-codegen DO NOT EDIT.

package tnb

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/tnb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Gets the details of a network package. A network package is a .zip file in CSAR
// (Cloud Service Archive) format defines the function packages you want to deploy
// and the Amazon Web Services infrastructure you want to deploy them on.
func (c *Client) GetSolNetworkPackage(ctx context.Context, params *GetSolNetworkPackageInput, optFns ...func(*Options)) (*GetSolNetworkPackageOutput, error) {
	if params == nil {
		params = &GetSolNetworkPackageInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "GetSolNetworkPackage", params, optFns, c.addOperationGetSolNetworkPackageMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*GetSolNetworkPackageOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type GetSolNetworkPackageInput struct {

	// ID of the network service descriptor in the network package.
	//
	// This member is required.
	NsdInfoId *string

	noSmithyDocumentSerde
}

type GetSolNetworkPackageOutput struct {

	// Network package ARN.
	//
	// This member is required.
	Arn *string

	// Network package ID.
	//
	// This member is required.
	Id *string

	// Metadata associated with a network package. A network package is a .zip file in
	// CSAR (Cloud Service Archive) format defines the function packages you want to
	// deploy and the Amazon Web Services infrastructure you want to deploy them on.
	//
	// This member is required.
	Metadata *types.GetSolNetworkPackageMetadata

	// Network service descriptor ID.
	//
	// This member is required.
	NsdId *string

	// Network service descriptor name.
	//
	// This member is required.
	NsdName *string

	// Network service descriptor onboarding state.
	//
	// This member is required.
	NsdOnboardingState types.NsdOnboardingState

	// Network service descriptor operational state.
	//
	// This member is required.
	NsdOperationalState types.NsdOperationalState

	// Network service descriptor usage state.
	//
	// This member is required.
	NsdUsageState types.NsdUsageState

	// Network service descriptor version.
	//
	// This member is required.
	NsdVersion *string

	// Identifies the function package for the function package descriptor referenced
	// by the onboarded network package.
	//
	// This member is required.
	VnfPkgIds []string

	// A tag is a label that you assign to an Amazon Web Services resource. Each tag
	// consists of a key and an optional value. You can use tags to search and filter
	// your resources or track your Amazon Web Services costs.
	Tags map[string]string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationGetSolNetworkPackageMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpGetSolNetworkPackage{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpGetSolNetworkPackage{}, middleware.After)
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
	if err = addOpGetSolNetworkPackageValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opGetSolNetworkPackage(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opGetSolNetworkPackage(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "tnb",
		OperationName: "GetSolNetworkPackage",
	}
}
