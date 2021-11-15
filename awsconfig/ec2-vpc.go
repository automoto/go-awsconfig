package awsconfig

import (
	"time"
)

type VPC struct {
	ConfigurationItem struct {
		Arn              string `json:"ARN"`
		AvailabilityZone string `json:"availabilityZone"`
		AwsAccountID     string `json:"awsAccountId"`
		AwsRegion        string `json:"awsRegion"`
		Configuration    struct {
			Description   string `json:"description"`
			GroupID       string `json:"groupId"`
			GroupName     string `json:"groupName"`
			IpPermissions []struct {
				FromPort   int      `json:"fromPort"`
				IpProtocol string   `json:"ipProtocol"`
				IpRanges   []string `json:"ipRanges"`
				Ipv4Ranges []struct {
					CidrIp string `json:"cidrIp"`
				} `json:"ipv4Ranges"`
				Ipv6Ranges       []interface{} `json:"ipv6Ranges"`
				PrefixListIds    []interface{} `json:"prefixListIds"`
				ToPort           int           `json:"toPort"`
				UserIDGroupPairs []interface{} `json:"userIdGroupPairs"`
			} `json:"ipPermissions"`
			IpPermissionsEgress []struct {
				IpProtocol string   `json:"ipProtocol"`
				IpRanges   []string `json:"ipRanges"`
				Ipv4Ranges []struct {
					CidrIp string `json:"cidrIp"`
				} `json:"ipv4Ranges"`
				Ipv6Ranges       []interface{} `json:"ipv6Ranges"`
				PrefixListIds    []interface{} `json:"prefixListIds"`
				UserIDGroupPairs []interface{} `json:"userIdGroupPairs"`
			} `json:"ipPermissionsEgress"`
			OwnerID string `json:"ownerId"`
			Tags    []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tags"`
			VpcID string `json:"vpcId"`
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
		ResourceName               string      `json:"resourceName"`
		ResourceType               string      `json:"resourceType"`
		SupplementaryConfiguration struct{}    `json:"supplementaryConfiguration"`
		Tags                       struct {
			Name string `json:"Name"`
		} `json:"tags"`
	} `json:"configurationItem"`
	ConfigurationItemDiff struct {
		ChangeType        string   `json:"changeType"`
		ChangedProperties struct{} `json:"changedProperties"`
	} `json:"configurationItemDiff"`
	MessageType              string    `json:"messageType"`
	NotificationCreationTime time.Time `json:"notificationCreationTime"`
	RecordVersion            string    `json:"recordVersion"`
}

