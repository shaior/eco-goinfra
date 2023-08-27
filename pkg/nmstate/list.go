package nmstate

import (
	"context"

	"github.com/golang/glog"
	nmstateV1 "github.com/nmstate/kubernetes-nmstate/api/v1"
	"github.com/openshift-kni/eco-goinfra/pkg/clients"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// ListPolicy returns a list of NodeNetworkConfigurationPolicy.
func ListPolicy(apiClient *clients.Settings) ([]*PolicyBuilder, error) {
	glog.V(100).Infof("Listing NodeNetworkConfigurationPolicy")

	policyList := &nmstateV1.NodeNetworkConfigurationPolicyList{}
	err := apiClient.Client.List(context.Background(), policyList)

	if err != nil {
		glog.V(100).Infof("Failed to list NodeNetworkConfigurationPolicy due to %s", err.Error())

		return nil, err
	}

	var networkConfigurationPolicyObjects []*PolicyBuilder

	for _, policy := range policyList.Items {
		copiedPolicy := policy
		policyBuilder := &PolicyBuilder{
			apiClient:  apiClient,
			Definition: &copiedPolicy,
			Object:     &copiedPolicy}

		networkConfigurationPolicyObjects = append(networkConfigurationPolicyObjects, policyBuilder)
	}

	return networkConfigurationPolicyObjects, nil
}

// ListNmState returns a NMState list.
func ListNmState(apiClient *clients.Settings) (nmstateV1.NMStateList, error) {
	nmStateList := &nmstateV1.NMStateList{}
	err := apiClient.List(context.TODO(), nmStateList, &client.ListOptions{})

	return *nmStateList, err
}
