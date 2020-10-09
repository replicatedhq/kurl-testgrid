package instances

import (
	kurlv1beta1 "github.com/replicatedhq/kurl/kurlkinds/pkg/apis/cluster/v1beta1"
	"github.com/replicatedhq/kurl/testgrid/tgrun/pkg/scheduler/types"
)

func init() {
	Instances = append(
		Instances,
		types.InstallerSpec{
			Kubernetes: kurlv1beta1.Kubernetes{
				Version: "1.18.9",
			},
			Weave: &kurlv1beta1.Weave{
				Version: "2.6.4",
			},
			Rook: &kurlv1beta1.Rook{
				Version: "1.0.4",
			},
			Contour: &kurlv1beta1.Contour{
				Version: "1.0.1",
			},
			Containerd: &kurlv1beta1.Containerd{
				Version: "1.2.13",
			},
			Prometheus: &kurlv1beta1.Prometheus{
				Version: "0.33.0",
			},
			Registry: &kurlv1beta1.Registry{
				Version: "2.7.1",
			},
			Velero: &kurlv1beta1.Velero{
				Version: "1.2.0",
			},
			Kotsadm: &kurlv1beta1.Kotsadm{
				Version: "1.19.0",
			},
		},
	)
}
