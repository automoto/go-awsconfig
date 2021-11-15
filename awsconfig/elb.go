package awsconfig

import "time"

type T struct {
	ConfigurationItem struct {
		Arn                          string        `json:"ARN"`
		AvailabilityZone             interface{}   `json:"availabilityZone"`
		AwsAccountID                 string        `json:"awsAccountId"`
		AwsRegion                    string        `json:"awsRegion"`
		Configuration                interface{}   `json:"configuration"`
		ConfigurationItemCaptureTime time.Time     `json:"configurationItemCaptureTime"`
		ConfigurationItemStatus      string        `json:"configurationItemStatus"`
		ConfigurationItemVersion     string        `json:"configurationItemVersion"`
		ConfigurationStateID         int           `json:"configurationStateId"`
		ConfigurationStateMd5Hash    string        `json:"configurationStateMd5Hash"`
		RelatedEvents                []interface{} `json:"relatedEvents"`
		Relationships                []interface{} `json:"relationships"`
		ResourceCreationTime         interface{}   `json:"resourceCreationTime"`
		ResourceID                   string        `json:"resourceId"`
		ResourceName                 string        `json:"resourceName"`
		ResourceType                 string        `json:"resourceType"`
		SupplementaryConfiguration   struct{}      `json:"supplementaryConfiguration"`
		Tags                         struct{}      `json:"tags"`
	} `json:"configurationItem"`
	ConfigurationItemDiff struct {
		ChangeType        string `json:"changeType"`
		ChangedProperties struct {
			Configuration struct {
				ChangeType    string `json:"changeType"`
				PreviousValue struct {
					AvailabilityZones         []string      `json:"availabilityZones"`
					BackendServerDescriptions []interface{} `json:"backendServerDescriptions"`
					CanonicalHostedZoneName   string        `json:"canonicalHostedZoneName"`
					CanonicalHostedZoneNameID string        `json:"canonicalHostedZoneNameID"`
					CreatedTime               int           `json:"createdTime"`
					Dnsname                   string        `json:"dnsname"`
					ListenerDescriptions      []struct {
						Listener struct {
							InstancePort     int    `json:"instancePort"`
							InstanceProtocol string `json:"instanceProtocol"`
							LoadBalancerPort int    `json:"loadBalancerPort"`
							Protocol         string `json:"protocol"`
						} `json:"listener"`
						PolicyNames []interface{} `json:"policyNames"`
					} `json:"listenerDescriptions"`
					LoadBalancerName string `json:"loadBalancerName"`
					Policies         struct {
						AppCookieStickinessPolicies []interface{} `json:"appCookieStickinessPolicies"`
						LbcookieStickinessPolicies  []interface{} `json:"lbcookieStickinessPolicies"`
						OtherPolicies               []interface{} `json:"otherPolicies"`
					} `json:"policies"`
					Scheme              string   `json:"scheme"`
					SecurityGroups      []string `json:"securityGroups"`
					SourceSecurityGroup struct {
						GroupName  string `json:"groupName"`
						OwnerAlias string `json:"ownerAlias"`
					} `json:"sourceSecurityGroup"`
					Subnets []string `json:"subnets"`
					Vpcid   string   `json:"vpcid"`
				} `json:"previousValue"`
				UpdatedValue interface{} `json:"updatedValue"`
			} `json:"Configuration"`
			Relationships_0 struct {
				ChangeType    string `json:"changeType"`
				PreviousValue struct {
					Name         string      `json:"name"`
					ResourceID   string      `json:"resourceId"`
					ResourceName interface{} `json:"resourceName"`
					ResourceType string      `json:"resourceType"`
				} `json:"previousValue"`
				UpdatedValue interface{} `json:"updatedValue"`
			} `json:"Relationships.0"`
			Relationships_1 struct {
				ChangeType    string `json:"changeType"`
				PreviousValue struct {
					Name         string      `json:"name"`
					ResourceID   string      `json:"resourceId"`
					ResourceName interface{} `json:"resourceName"`
					ResourceType string      `json:"resourceType"`
				} `json:"previousValue"`
				UpdatedValue interface{} `json:"updatedValue"`
			} `json:"Relationships.1"`
			Relationships_2 struct {
				ChangeType    string `json:"changeType"`
				PreviousValue struct {
					Name         string      `json:"name"`
					ResourceID   string      `json:"resourceId"`
					ResourceName interface{} `json:"resourceName"`
					ResourceType string      `json:"resourceType"`
				} `json:"previousValue"`
				UpdatedValue interface{} `json:"updatedValue"`
			} `json:"Relationships.2"`
			Relationships_3 struct {
				ChangeType    string `json:"changeType"`
				PreviousValue struct {
					Name         string      `json:"name"`
					ResourceID   string      `json:"resourceId"`
					ResourceName interface{} `json:"resourceName"`
					ResourceType string      `json:"resourceType"`
				} `json:"previousValue"`
				UpdatedValue interface{} `json:"updatedValue"`
			} `json:"Relationships.3"`
			SupplementaryConfiguration_LoadBalancerAttributes struct {
				ChangeType    string `json:"changeType"`
				PreviousValue struct {
					AccessLog struct {
						Enabled bool `json:"enabled"`
					} `json:"accessLog"`
					AdditionalAttributes []struct {
						Key   string `json:"key"`
						Value string `json:"value"`
					} `json:"additionalAttributes"`
					ConnectionDraining struct {
						Enabled bool    `json:"enabled"`
						Timeout float64 `json:"timeout"`
					} `json:"connectionDraining"`
					ConnectionSettings struct {
						IdleTimeout float64 `json:"idleTimeout"`
					} `json:"connectionSettings"`
					CrossZoneLoadBalancing struct {
						Enabled bool `json:"enabled"`
					} `json:"crossZoneLoadBalancing"`
				} `json:"previousValue"`
				UpdatedValue interface{} `json:"updatedValue"`
			} `json:"SupplementaryConfiguration.LoadBalancerAttributes"`
			SupplementaryConfiguration_LoadBalancerPolicies struct {
				ChangeType    string        `json:"changeType"`
				PreviousValue []interface{} `json:"previousValue"`
				UpdatedValue  interface{}   `json:"updatedValue"`
			} `json:"SupplementaryConfiguration.LoadBalancerPolicies"`
			SupplementaryConfiguration_Tags struct {
				ChangeType    string        `json:"changeType"`
				PreviousValue []interface{} `json:"previousValue"`
				UpdatedValue  interface{}   `json:"updatedValue"`
			} `json:"SupplementaryConfiguration.Tags"`
		} `json:"changedProperties"`
	} `json:"configurationItemDiff"`
	MessageType              string    `json:"messageType"`
	NotificationCreationTime time.Time `json:"notificationCreationTime"`
	RecordVersion            string    `json:"recordVersion"`
}

