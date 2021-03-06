// deepcopy-gen is a tool for auto-generating DeepCopy functions.
//
// Structs in the input directories with the below line in their comments
// will be ignored during generation.
// // +gencopy=false
package main

import (
	"strings"

	"github.com/golang/glog"

	"k8s.io/gengo/args"
	"k8s.io/gengo/examples/deepcopy-gen/generators"
	"k8s.io/gengo/generator"
)

func main() {
	arguments := args.Default()

	// Override defaults. These are Kubernetes specific input locations.
	arguments.InputDirs = []string{
		"github.com/openshift/origin/pkg/build/admission/testing",
		"github.com/openshift/origin/pkg/build/controller/build/defaults/api",
		"github.com/openshift/origin/pkg/build/controller/build/defaults/api/v1",
		"github.com/openshift/origin/pkg/build/controller/build/overrides/api",
		"github.com/openshift/origin/pkg/build/controller/build/overrides/api/v1",
		"github.com/openshift/origin/pkg/cmd/server/api",
		"github.com/openshift/origin/pkg/cmd/server/api/v1",
		"github.com/openshift/origin/pkg/cmd/server/api/v1/testing",
		"github.com/openshift/origin/pkg/cmd/util/pluginconfig/testing",
		"github.com/openshift/origin/pkg/image/admission/imagepolicy/api",
		"github.com/openshift/origin/pkg/image/admission/imagepolicy/api/v1",
		"github.com/openshift/origin/pkg/ingress/admission/api",
		"github.com/openshift/origin/pkg/ingress/admission/api/v1",
		"github.com/openshift/origin/pkg/project/admission/lifecycle/testing",
		"github.com/openshift/origin/pkg/project/admission/requestlimit/api",
		"github.com/openshift/origin/pkg/project/admission/requestlimit/api/v1",
		"github.com/openshift/origin/pkg/quota/admission/clusterresourceoverride/api",
		"github.com/openshift/origin/pkg/quota/admission/clusterresourceoverride/api/v1",
		"github.com/openshift/origin/pkg/quota/admission/runonceduration/api",
		"github.com/openshift/origin/pkg/quota/admission/runonceduration/api/v1",
		"github.com/openshift/origin/pkg/scheduler/admission/podnodeconstraints/api",
		"github.com/openshift/origin/pkg/scheduler/admission/podnodeconstraints/api/v1",
		"github.com/openshift/origin/pkg/template/servicebroker/apis/config",
		"github.com/openshift/origin/pkg/template/servicebroker/apis/config/v1",
		"github.com/openshift/origin/pkg/util/testing",
		"github.com/openshift/origin/test/integration/testing",
		// internal apis
		"github.com/openshift/origin/pkg/apps/apis/apps",
		"github.com/openshift/origin/pkg/authorization/apis/authorization",
		"github.com/openshift/origin/pkg/build/apis/build",
		"github.com/openshift/origin/pkg/image/apis/image",
		"github.com/openshift/origin/pkg/network/apis/network",
		"github.com/openshift/origin/pkg/oauth/apis/oauth",
		"github.com/openshift/origin/pkg/project/apis/project",
		"github.com/openshift/origin/pkg/quota/apis/quota",
		"github.com/openshift/origin/pkg/route/apis/route",
		"github.com/openshift/origin/pkg/security/apis/security",
		"github.com/openshift/origin/pkg/template/apis/template",
		"github.com/openshift/origin/pkg/user/apis/user",
	}

	arguments.GeneratedBuildTag = "ignore_autogenerated_openshift"
	arguments.GoHeaderFilePath = "hack/boilerplate.txt"
	arguments.OutputFileBaseName = "zz_generated.deepcopy"
	arguments.CustomArgs = &generators.CustomArgs{
		BoundingDirs: []string{
			"k8s.io/kubernetes",
			"github.com/openshift/origin",
		},
	}

	if err := arguments.Execute(
		generators.NameSystems(),
		generators.DefaultNameSystem(),
		func(context *generator.Context, arguments *args.GeneratorArgs) generator.Packages {
			pkgs := generators.Packages(context, arguments)
			var include generator.Packages
			for _, pkg := range pkgs {
				if strings.HasPrefix(pkg.Path(), "k8s.io/") {
					continue
				}
				include = append(include, pkg)
			}
			return include
		},
	); err != nil {
		glog.Fatalf("Error: %v", err)
	}
	glog.Info("Completed successfully.")
}
