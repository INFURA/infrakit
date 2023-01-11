package controllers

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// Reconciler holds generic reconciler information.
type Reconciler struct {
	Client client.Client
	Scheme *runtime.Scheme
}
