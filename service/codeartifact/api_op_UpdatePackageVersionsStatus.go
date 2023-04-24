// Code generated by smithy-go-codegen DO NOT EDIT.

package codeartifact

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/codeartifact/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Updates the status of one or more versions of a package. Using
// UpdatePackageVersionsStatus , you can update the status of package versions to
// Archived , Published , or Unlisted . To set the status of a package version to
// Disposed , use DisposePackageVersions (https://docs.aws.amazon.com/codeartifact/latest/APIReference/API_DisposePackageVersions.html)
// .
func (c *Client) UpdatePackageVersionsStatus(ctx context.Context, params *UpdatePackageVersionsStatusInput, optFns ...func(*Options)) (*UpdatePackageVersionsStatusOutput, error) {
	if params == nil {
		params = &UpdatePackageVersionsStatusInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "UpdatePackageVersionsStatus", params, optFns, c.addOperationUpdatePackageVersionsStatusMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*UpdatePackageVersionsStatusOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type UpdatePackageVersionsStatusInput struct {

	// The name of the domain that contains the repository that contains the package
	// versions with a status to be updated.
	//
	// This member is required.
	Domain *string

	// A format that specifies the type of the package with the statuses to update.
	//
	// This member is required.
	Format types.PackageFormat

	// The name of the package with the version statuses to update.
	//
	// This member is required.
	Package *string

	// The repository that contains the package versions with the status you want to
	// update.
	//
	// This member is required.
	Repository *string

	// The status you want to change the package version status to.
	//
	// This member is required.
	TargetStatus types.PackageVersionStatus

	// An array of strings that specify the versions of the package with the statuses
	// to update.
	//
	// This member is required.
	Versions []string

	// The 12-digit account number of the Amazon Web Services account that owns the
	// domain. It does not include dashes or spaces.
	DomainOwner *string

	// The package version’s expected status before it is updated. If expectedStatus
	// is provided, the package version's status is updated only if its status at the
	// time UpdatePackageVersionsStatus is called matches expectedStatus .
	ExpectedStatus types.PackageVersionStatus

	// The namespace of the package version to be updated. The package version
	// component that specifies its namespace depends on its type. For example:
	//   - The namespace of a Maven package version is its groupId .
	//   - The namespace of an npm package version is its scope .
	//   - Python and NuGet package versions do not contain a corresponding component,
	//   package versions of those formats do not have a namespace.
	//   - The namespace of a generic package is its namespace .
	Namespace *string

	// A map of package versions and package version revisions. The map key is the
	// package version (for example, 3.5.2 ), and the map value is the package version
	// revision.
	VersionRevisions map[string]string

	noSmithyDocumentSerde
}

type UpdatePackageVersionsStatusOutput struct {

	// A list of SuccessfulPackageVersionInfo objects, one for each package version
	// with a status that successfully updated.
	FailedVersions map[string]types.PackageVersionError

	// A list of PackageVersionError objects, one for each package version with a
	// status that failed to update.
	SuccessfulVersions map[string]types.SuccessfulPackageVersionInfo

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationUpdatePackageVersionsStatusMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpUpdatePackageVersionsStatus{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpUpdatePackageVersionsStatus{}, middleware.After)
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
	if err = addOpUpdatePackageVersionsStatusValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opUpdatePackageVersionsStatus(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opUpdatePackageVersionsStatus(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "codeartifact",
		OperationName: "UpdatePackageVersionsStatus",
	}
}
