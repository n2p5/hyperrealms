// A generated module for Hyperrealms functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/hyperrealms/internal/dagger"
)

const (
	goVersion        = "1.23.5"
	goModCacheVolume = "hyperrealms-mod-cache"
	containerSrcDir  = "/src"
	containerModDir  = "/go/pkg/mod"
)

type Hyperrealms struct{}

// Return the result of running unit tests
func (m *Hyperrealms) Test(ctx context.Context, src *dagger.Directory) (string, error) {
	return m.BuildEnv(src).
		WithExec([]string{"go", "test", ".", "-v"}).
		Stdout(ctx)
}

func (m *Hyperrealms) Build(ctx context.Context, src *dagger.Directory) (string, error) {
	return m.BuildEnv(src).
		WithExec([]string{"go", "build", "-o", "app"}).
		Stdout(ctx)
}

func (m *Hyperrealms) BuildEnv(src *dagger.Directory) *dagger.Container {
	goModCache := dag.CacheVolume(goModCacheVolume)
	return dag.Container().
		From("golang:"+goVersion).
		WithDirectory(containerSrcDir, src).
		WithWorkdir(containerSrcDir).
		WithMountedCache("/go/pkg/mod", goModCache).
		WithEnvVariable("GOMODCACHE", containerModDir)
}
