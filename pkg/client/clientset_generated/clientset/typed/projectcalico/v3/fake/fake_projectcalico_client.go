// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"

	v3 "github.com/projectcalico/apiserver/pkg/client/clientset_generated/clientset/typed/projectcalico/v3"
)

type FakeProjectcalicoV3 struct {
	*testing.Fake
}

func (c *FakeProjectcalicoV3) BGPConfigurations() v3.BGPConfigurationInterface {
	return &FakeBGPConfigurations{c}
}

func (c *FakeProjectcalicoV3) BGPPeers() v3.BGPPeerInterface {
	return &FakeBGPPeers{c}
}

func (c *FakeProjectcalicoV3) ClusterInformations() v3.ClusterInformationInterface {
	return &FakeClusterInformations{c}
}

func (c *FakeProjectcalicoV3) FelixConfigurations() v3.FelixConfigurationInterface {
	return &FakeFelixConfigurations{c}
}

func (c *FakeProjectcalicoV3) GlobalNetworkPolicies() v3.GlobalNetworkPolicyInterface {
	return &FakeGlobalNetworkPolicies{c}
}

func (c *FakeProjectcalicoV3) GlobalNetworkSets() v3.GlobalNetworkSetInterface {
	return &FakeGlobalNetworkSets{c}
}

func (c *FakeProjectcalicoV3) HostEndpoints() v3.HostEndpointInterface {
	return &FakeHostEndpoints{c}
}

func (c *FakeProjectcalicoV3) IPPools() v3.IPPoolInterface {
	return &FakeIPPools{c}
}

func (c *FakeProjectcalicoV3) KubeControllersConfigurations() v3.KubeControllersConfigurationInterface {
	return &FakeKubeControllersConfigurations{c}
}

func (c *FakeProjectcalicoV3) NetworkPolicies(namespace string) v3.NetworkPolicyInterface {
	return &FakeNetworkPolicies{c, namespace}
}

func (c *FakeProjectcalicoV3) NetworkSets(namespace string) v3.NetworkSetInterface {
	return &FakeNetworkSets{c, namespace}
}

func (c *FakeProjectcalicoV3) Profiles() v3.ProfileInterface {
	return &FakeProfiles{c}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeProjectcalicoV3) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
