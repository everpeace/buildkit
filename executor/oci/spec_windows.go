// +build windows

package oci

import (
	"github.com/containerd/containerd/contrib/seccomp"
	"github.com/containerd/containerd/oci"
	"github.com/docker/docker/pkg/idtools"
	"github.com/moby/buildkit/solver/pb"
	"github.com/moby/buildkit/util/system"
	"github.com/pkg/errors"
)

func generateMountOpts(resolvConf, hostsFile string) ([]oci.SpecOpts, error) {
	return nil, nil
}

// generateSecurityOpts may affect mounts, so must be called after generateMountOpts
func generateSecurityOpts(mode pb.SecurityMode) ([]oci.SpecOpts, error) {
	if mode == pb.SecurityMode_INSECURE {
		return nil, errors.New("no support for running in insecure mode on Windows")
	} else if system.SeccompSupported() && mode == pb.SecurityMode_SANDBOX {
		// TODO: Can LCOW support seccomp? Does that even make sense?
		return []oci.SpecOpts{seccomp.WithDefaultProfile()}, nil
	}
	return nil, nil
}

// generateProcessModeOpts may affect mounts, so must be called after generateMountOpts
func generateProcessModeOpts(mode ProcessMode) ([]oci.SpecOpts, error) {
	if mode == NoProcessSandbox {
		return nil, errors.New("no support for NoProcessSandbox on Windows")
	}
	return nil, nil
}

func generateIDmapOpts(idmap *idtools.IdentityMapping) ([]oci.SpecOpts, error) {
	if idmap == nil {
		return nil, nil
	}
	return nil, errors.New("no support for IdentityMapping on Windows")
}