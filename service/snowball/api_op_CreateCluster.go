// Code generated by smithy-go-codegen DO NOT EDIT.

package snowball

import (
	"context"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/snowball/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Creates an empty cluster. Each cluster supports five nodes. You use the
// CreateJob action separately to create the jobs for each of these nodes. The
// cluster does not ship until these five node jobs have been created.
func (c *Client) CreateCluster(ctx context.Context, params *CreateClusterInput, optFns ...func(*Options)) (*CreateClusterOutput, error) {
	if params == nil {
		params = &CreateClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "CreateCluster", params, optFns, c.addOperationCreateClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*CreateClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type CreateClusterInput struct {

	// The ID for the address that you want the cluster shipped to.
	//
	// This member is required.
	AddressId *string

	// The type of job for this cluster. Currently, the only job type supported for
	// clusters is LOCAL_USE . For more information, see
	// "https://docs.aws.amazon.com/snowball/latest/snowcone-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide or
	// "https://docs.aws.amazon.com/snowball/latest/developer-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide.
	//
	// This member is required.
	JobType types.JobType

	// The shipping speed for each node in this cluster. This speed doesn't dictate
	// how soon you'll get each Snowball Edge device, rather it represents how quickly
	// each device moves to its destination while in transit. Regional shipping speeds
	// are as follows:
	//   - In Australia, you have access to express shipping. Typically, Snow devices
	//   shipped express are delivered in about a day.
	//   - In the European Union (EU), you have access to express shipping. Typically,
	//   Snow devices shipped express are delivered in about a day. In addition, most
	//   countries in the EU have access to standard shipping, which typically takes less
	//   than a week, one way.
	//   - In India, Snow devices are delivered in one to seven days.
	//   - In the United States of America (US), you have access to one-day shipping
	//   and two-day shipping.
	//
	//   - In Australia, you have access to express shipping. Typically, devices
	//   shipped express are delivered in about a day.
	//   - In the European Union (EU), you have access to express shipping. Typically,
	//   Snow devices shipped express are delivered in about a day. In addition, most
	//   countries in the EU have access to standard shipping, which typically takes less
	//   than a week, one way.
	//   - In India, Snow devices are delivered in one to seven days.
	//   - In the US, you have access to one-day shipping and two-day shipping.
	//
	// This member is required.
	ShippingOption types.ShippingOption

	// The type of Snow Family devices to use for this cluster. For cluster jobs,
	// Amazon Web Services Snow Family currently supports only the EDGE device type.
	// For more information, see
	// "https://docs.aws.amazon.com/snowball/latest/snowcone-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide or
	// "https://docs.aws.amazon.com/snowball/latest/developer-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide.
	//
	// This member is required.
	SnowballType types.SnowballType

	// An optional description of this specific cluster, for example Environmental
	// Data Cluster-01 .
	Description *string

	// Force to create cluster when user attempts to overprovision or underprovision a
	// cluster. A cluster is overprovisioned or underprovisioned if the initial size of
	// the cluster is more (overprovisioned) or less (underprovisioned) than what
	// needed to meet capacity requirement specified with OnDeviceServiceConfiguration .
	ForceCreateJobs bool

	// The forwarding address ID for a cluster. This field is not supported in most
	// regions.
	ForwardingAddressId *string

	// If provided, each job will be automatically created and associated with the new
	// cluster. If not provided, will be treated as 0.
	InitialClusterSize *int32

	// The KmsKeyARN value that you want to associate with this cluster. KmsKeyARN
	// values are created by using the CreateKey (https://docs.aws.amazon.com/kms/latest/APIReference/API_CreateKey.html)
	// API action in Key Management Service (KMS).
	KmsKeyARN *string

	// Lists long-term pricing id that will be used to associate with jobs
	// automatically created for the new cluster.
	LongTermPricingIds []string

	// The Amazon Simple Notification Service (Amazon SNS) notification settings for
	// this cluster.
	Notification *types.Notification

	// Specifies the service or services on the Snow Family device that your
	// transferred data will be exported from or imported into. Amazon Web Services
	// Snow Family device clusters support Amazon S3 and NFS (Network File System).
	OnDeviceServiceConfiguration *types.OnDeviceServiceConfiguration

	// Allows you to securely operate and manage Snow devices in a cluster remotely
	// from outside of your internal network. When set to INSTALLED_AUTOSTART , remote
	// management will automatically be available when the device arrives at your
	// location. Otherwise, you need to use the Snowball Client to manage the device.
	RemoteManagement types.RemoteManagement

	// The resources associated with the cluster job. These resources include Amazon
	// S3 buckets and optional Lambda functions written in the Python language.
	Resources *types.JobResource

	// The RoleARN that you want to associate with this cluster. RoleArn values are
	// created by using the CreateRole (https://docs.aws.amazon.com/IAM/latest/APIReference/API_CreateRole.html)
	// API action in Identity and Access Management (IAM).
	RoleARN *string

	// If your job is being created in one of the US regions, you have the option of
	// specifying what size Snow device you'd like for this job. In all other regions,
	// Snowballs come with 80 TB in storage capacity. For more information, see
	// "https://docs.aws.amazon.com/snowball/latest/snowcone-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide or
	// "https://docs.aws.amazon.com/snowball/latest/developer-guide/snow-device-types.html"
	// (Snow Family Devices and Capacity) in the Snowcone User Guide.
	SnowballCapacityPreference types.SnowballCapacity

	// The tax documents required in your Amazon Web Services Region.
	TaxDocuments *types.TaxDocuments

	noSmithyDocumentSerde
}

type CreateClusterOutput struct {

	// The automatically generated ID for a cluster.
	ClusterId *string

	// List of jobs created for this cluster. For syntax, see
	// ListJobsResult$JobListEntries (https://docs.aws.amazon.com/snowball/latest/api-reference/API_ListJobs.html#API_ListJobs_ResponseSyntax)
	// in this guide.
	JobListEntries []types.JobListEntry

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationCreateClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpCreateCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpCreateCluster{}, middleware.After)
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
	if err = addOpCreateClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opCreateCluster(options.Region), middleware.Before); err != nil {
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

func newServiceMetadataMiddleware_opCreateCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "snowball",
		OperationName: "CreateCluster",
	}
}
