package awsconfig

import (
	"time"
)

type ENI struct {
	ConfigurationItem struct {
		Arn              string `json:"ARN"`
		AvailabilityZone string `json:"availabilityZone"`
		AwsAccountID     string `json:"awsAccountId"`
		AwsRegion        string `json:"awsRegion"`
		Configuration    struct {
			Association struct {
				AllocationID    interface{} `json:"allocationId"`
				AssociationID   interface{} `json:"associationId"`
				CarrierIp       interface{} `json:"carrierIp"`
				CustomerOwnedIp interface{} `json:"customerOwnedIp"`
				IpOwnerID       string      `json:"ipOwnerId"`
				PublicDnsName   string      `json:"publicDnsName"`
				PublicIp        string      `json:"publicIp"`
			} `json:"association"`
			Attachment struct {
				AttachTime          time.Time   `json:"attachTime"`
				AttachmentID        string      `json:"attachmentId"`
				DeleteOnTermination bool        `json:"deleteOnTermination"`
				DeviceIndex         int         `json:"deviceIndex"`
				InstanceID          interface{} `json:"instanceId"`
				InstanceOwnerID     string      `json:"instanceOwnerId"`
				NetworkCardIndex    int         `json:"networkCardIndex"`
				Status              string      `json:"status"`
			} `json:"attachment"`
			AvailabilityZone string `json:"availabilityZone"`
			Description      string `json:"description"`
			Groups           []struct {
				GroupID   string `json:"groupId"`
				GroupName string `json:"groupName"`
			} `json:"groups"`
			InterfaceType      string        `json:"interfaceType"`
			Ipv6Addresses      []interface{} `json:"ipv6Addresses"`
			MacAddress         string        `json:"macAddress"`
			NetworkInterfaceID string        `json:"networkInterfaceId"`
			OutpostArn         interface{}   `json:"outpostArn"`
			OwnerID            string        `json:"ownerId"`
			PrivateDnsName     string        `json:"privateDnsName"`
			PrivateIpAddress   string        `json:"privateIpAddress"`
			PrivateIpAddresses []struct {
				Association struct {
					AllocationID    interface{} `json:"allocationId"`
					AssociationID   interface{} `json:"associationId"`
					CarrierIp       interface{} `json:"carrierIp"`
					CustomerOwnedIp interface{} `json:"customerOwnedIp"`
					IpOwnerID       string      `json:"ipOwnerId"`
					PublicDnsName   string      `json:"publicDnsName"`
					PublicIp        string      `json:"publicIp"`
				} `json:"association"`
				Primary          bool   `json:"primary"`
				PrivateDnsName   string `json:"privateDnsName"`
				PrivateIpAddress string `json:"privateIpAddress"`
			} `json:"privateIpAddresses"`
			RequesterID      string        `json:"requesterId"`
			RequesterManaged bool          `json:"requesterManaged"`
			SourceDestCheck  bool          `json:"sourceDestCheck"`
			Status           string        `json:"status"`
			SubnetID         string        `json:"subnetId"`
			TagSet           []interface{} `json:"tagSet"`
			VpcID            string        `json:"vpcId"`
		} `json:"configuration"`
		ConfigurationItemCaptureTime time.Time     `json:"configurationItemCaptureTime"`
		ConfigurationItemStatus      string        `json:"configurationItemStatus"`
		ConfigurationItemVersion     string        `json:"configurationItemVersion"`
		ConfigurationStateID         int           `json:"configurationStateId"`
		ConfigurationStateMd5Hash    string        `json:"configurationStateMd5Hash"`
		RelatedEvents                []interface{} `json:"relatedEvents"`
		Relationships                []struct {
			Name         string      `json:"name"`
			ResourceID   string      `json:"resourceId"`
			ResourceName interface{} `json:"resourceName"`
			ResourceType string      `json:"resourceType"`
		} `json:"relationships"`
		ResourceCreationTime       interface{} `json:"resourceCreationTime"`
		ResourceID                 string      `json:"resourceId"`
		ResourceName               interface{} `json:"resourceName"`
		ResourceType               string      `json:"resourceType"`
		SupplementaryConfiguration struct{}    `json:"supplementaryConfiguration"`
		Tags                       struct{}    `json:"tags"`
	} `json:"configurationItem"`
	ConfigurationItemDiff struct {
		ChangeType        string   `json:"changeType"`
		ChangedProperties struct{} `json:"changedProperties"`
	} `json:"configurationItemDiff"`
	MessageType              string    `json:"messageType"`
	NotificationCreationTime time.Time `json:"notificationCreationTime"`
	RecordVersion            string    `json:"recordVersion"`
}
