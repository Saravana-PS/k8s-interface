package v1

import (
	"encoding/json"

	"github.com/armosec/k8s-interface/cloudsupport/mockobjects"
	"github.com/aws/aws-sdk-go/service/eks"
)

func NewEKSSupportMock() *EKSSupportMock {
	return &EKSSupportMock{}
}

type EKSSupportMock struct {
}

// Get descriptive info about cluster running in EKS.
func (eksSupportM *EKSSupportMock) GetClusterDescribe(currContext string, region string) (*eks.DescribeClusterOutput, error) {
	describeClusterOutput := &eks.DescribeClusterOutput{}
	err := json.Unmarshal([]byte(mockobjects.EksDescriptor), describeClusterOutput)
	return describeClusterOutput, err
}

// getName get cluster name from describe
func (eksSupportM *EKSSupportMock) GetName(describe *eks.DescribeClusterOutput) string {
	return *describe.Cluster.Name
}
