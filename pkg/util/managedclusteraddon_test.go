// Copyright (c) 2020 Red Hat, Inc.

package util

import (
	"context"
	"testing"

	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes/scheme"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"

	addonv1alpha1 "github.com/open-cluster-management/api/addon/v1alpha1"
)

const (
	namespace = "test-namespace"
)

func TestManagedClusterAddon(t *testing.T) {
	s := scheme.Scheme
	addonv1alpha1.AddToScheme(s)
	c := fake.NewFakeClient()
	err := CreateManagedClusterAddonCR(c, namespace)
	if err != nil {
		t.Fatalf("Failed to create managedclusteraddon: (%v)", err)
	}
	addon := &addonv1alpha1.ManagedClusterAddOn{}
	err = c.Get(context.TODO(),
		types.NamespacedName{
			Name:      ManagedClusterAddonName,
			Namespace: namespace,
		},
		addon,
	)
	if err != nil {
		t.Fatalf("Failed to get managedclusteraddon: (%v)", err)
	}
}