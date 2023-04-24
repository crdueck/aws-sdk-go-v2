// Code generated by smithy-go-codegen DO NOT EDIT.

package quicksight

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/quicksight/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates a template either from a TemplateDefinition or from an existing Amazon
// QuickSight analysis or template. You can use the resulting template to create
// additional dashboards, templates, or analyses. A template is an entity in Amazon
// QuickSight that encapsulates the metadata required to create an analysis and
// that you can use to create s dashboard. A template adds a layer of abstraction
// by using placeholders to replace the dataset associated with the analysis. You
// can use templates to create dashboards by replacing dataset placeholders with
// datasets that follow the same schema that was used to create the source analysis
// and template.
func (c *Client) CreateTemplate(ctx context.Context, params *CreateTemplateInput, optFns ...func(*Options)) (*CreateTemplateOutput, error) {
	if params == nil {
		params = &CreateTemplateInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateTemplate", params, optFns, c.addOperationCreateTemplateMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateTemplateOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateTemplateInput struct {

	// The ID for the Amazon Web Services account that the group is in. You use the ID
	// for the Amazon Web Services account that contains your Amazon QuickSight
	// account.
	//
	// This member is required.
	AwsAccountId *string

	// An ID for the template that you want to create. This template is unique per
	// Amazon Web Services Region; in each Amazon Web Services account.
	//
	// This member is required.
	TemplateId *string

	// The definition of a template. A definition is the data model of all features in
	// a Dashboard, Template, or Analysis. Either a SourceEntity or a Definition must
	// be provided in order for the request to be valid.
	Definition *types.TemplateVersionDefinition

	// A display name for the template.
	Name *string

	// A list of resource permissions to be set on the template.
	Permissions []types.ResourcePermission

	// The entity that you are using as a source when you create the template. In
	// SourceEntity , you specify the type of object you're using as source:
	// SourceTemplate for a template or SourceAnalysis for an analysis. Both of these
	// require an Amazon Resource Name (ARN). For SourceTemplate , specify the ARN of
	// the source template. For SourceAnalysis , specify the ARN of the source
	// analysis. The SourceTemplate ARN can contain any Amazon Web Services account
	// and any Amazon QuickSight-supported Amazon Web Services Region. Use the
	// DataSetReferences entity within SourceTemplate or SourceAnalysis to list the
	// replacement datasets for the placeholders listed in the original. The schema in
	// each dataset must match its placeholder. Either a SourceEntity or a Definition
	// must be provided in order for the request to be valid.
	SourceEntity *types.TemplateSourceEntity

	// Contains a map of the key-value pairs for the resource tag or tags assigned to
	// the resource.
	Tags []types.Tag

	// A description of the current template version being created. This API operation
	// creates the first version of the template. Every time UpdateTemplate is called,
	// a new version is created. Each version of the template maintains a description
	// of the version in the VersionDescription field.
	VersionDescription *string

	noSmithyDocumentSerde
}

type CreateTemplateOutput struct {

	// The ARN for the template.
	Arn *string

	// The template creation status.
	CreationStatus types.ResourceStatus

	// The Amazon Web Services request ID for this operation.
	RequestId *string

	// The HTTP status of the request.
	Status int32

	// The ID of the template.
	TemplateId *string

	// The ARN for the template, including the version information of the first
	// version.
	VersionArn *string

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateTemplateMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsRestjson1_serializeOpCreateTemplate{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsRestjson1_deserializeOpCreateTemplate{}, middleware.After)
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
	if err = addOpCreateTemplateValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateTemplate(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateTemplate(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "quicksight",
		OperationName: "CreateTemplate",
	}
}
