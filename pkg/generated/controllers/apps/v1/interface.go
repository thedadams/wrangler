/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package v1

import (
	"github.com/rancher/lasso/pkg/controller"
	"github.com/rancher/wrangler/pkg/generic"
	"github.com/rancher/wrangler/pkg/schemes"
	v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

func init() {
	schemes.Register(v1.AddToScheme)
}

type Interface interface {
	DaemonSet() DaemonSetController
	Deployment() DeploymentController
	StatefulSet() StatefulSetController
}

func New(controllerFactory controller.SharedControllerFactory) Interface {
	return &version{
		controllerFactory: controllerFactory,
	}
}

type version struct {
	controllerFactory controller.SharedControllerFactory
}

func (v *version) DaemonSet() DaemonSetController {
	return &DaemonSetGenericController{
		generic.NewController[*v1.DaemonSet, *v1.DaemonSetList](schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "DaemonSet"}, "daemonsets", true, v.controllerFactory),
	}
}

func (v *version) Deployment() DeploymentController {
	return &DeploymentGenericController{
		generic.NewController[*v1.Deployment, *v1.DeploymentList](schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "Deployment"}, "deployments", true, v.controllerFactory),
	}
}

func (v *version) StatefulSet() StatefulSetController {
	return &StatefulSetGenericController{
		generic.NewController[*v1.StatefulSet, *v1.StatefulSetList](schema.GroupVersionKind{Group: "apps", Version: "v1", Kind: "StatefulSet"}, "statefulsets", true, v.controllerFactory),
	}
}
