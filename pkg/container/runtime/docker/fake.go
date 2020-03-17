package docker

import (
	"context"
	"io"
	"time"

	dockertypes "github.com/docker/docker/api/types"
	containertypes "github.com/docker/docker/api/types/container"
	networktypes "github.com/docker/docker/api/types/network"
)

// FakeClient is a mock of Docker client.
type FakeClient struct { //nolint:dupl
	ContainerCreateF   func(ctx context.Context, config *containertypes.Config, hostConfig *containertypes.HostConfig, networkingConfig *networktypes.NetworkingConfig, containerName string) (containertypes.ContainerCreateCreatedBody, error)
	ContainerStartF    func(ctx context.Context, container string, options dockertypes.ContainerStartOptions) error
	ContainerStopF     func(ctx context.Context, container string, timeout *time.Duration) error
	ContainerInspectF  func(ctx context.Context, container string) (dockertypes.ContainerJSON, error)
	ContainerRemoveF   func(ctx context.Context, container string, options dockertypes.ContainerRemoveOptions) error
	CopyFromContainerF func(ctx context.Context, container, srcPath string) (io.ReadCloser, dockertypes.ContainerPathStat, error)
	CopyToContainerF   func(ctx context.Context, container, path string, content io.Reader, options dockertypes.CopyToContainerOptions) error
	ContainerStatPathF func(ctx context.Context, container, path string) (dockertypes.ContainerPathStat, error)
	ImageListF         func(ctx context.Context, options dockertypes.ImageListOptions) ([]dockertypes.ImageSummary, error)
	ImagePullF         func(ctx context.Context, ref string, options dockertypes.ImagePullOptions) (io.ReadCloser, error)
}

// ContainerCreate mocks Docker client ContainerCreate().
func (f *FakeClient) ContainerCreate(ctx context.Context, config *containertypes.Config, hostConfig *containertypes.HostConfig, networkingConfig *networktypes.NetworkingConfig, containerName string) (containertypes.ContainerCreateCreatedBody, error) {
	return f.ContainerCreateF(ctx, config, hostConfig, networkingConfig, containerName)
}

// ContainerStart mocks Docker client ContainerStart().
func (f *FakeClient) ContainerStart(ctx context.Context, container string, options dockertypes.ContainerStartOptions) error {
	return f.ContainerStartF(ctx, container, options)
}

// ContainerStop mocks Docker client ContainerStop().
func (f *FakeClient) ContainerStop(ctx context.Context, container string, timeout *time.Duration) error {
	return f.ContainerStopF(ctx, container, timeout)
}

// ContainerInspect mocks Docker client ContainerInspect().
func (f *FakeClient) ContainerInspect(ctx context.Context, container string) (dockertypes.ContainerJSON, error) {
	return f.ContainerInspectF(ctx, container)
}

// ContainerRemove mocks Docker client ContainerRemove().
func (f *FakeClient) ContainerRemove(ctx context.Context, container string, options dockertypes.ContainerRemoveOptions) error {
	return f.ContainerRemoveF(ctx, container, options)
}

// CopyFromContainer mocks Docker client CopyFromContainer().
func (f *FakeClient) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, dockertypes.ContainerPathStat, error) {
	return f.CopyFromContainerF(ctx, container, srcPath)
}

// CopyToContainer mocks Docker client CopyToContainer().
func (f *FakeClient) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options dockertypes.CopyToContainerOptions) error {
	return f.CopyToContainerF(ctx, container, path, content, options)
}

// ContainerStatPath mocks Docker client ContainerStatPath().
func (f *FakeClient) ContainerStatPath(ctx context.Context, container, path string) (dockertypes.ContainerPathStat, error) {
	return f.ContainerStatPathF(ctx, container, path)
}

// ImageList mocks Docker client ImageList().
func (f *FakeClient) ImageList(ctx context.Context, options dockertypes.ImageListOptions) ([]dockertypes.ImageSummary, error) {
	return f.ImageListF(ctx, options)
}

// ImagePull mocks Docker client ImagePull().
func (f *FakeClient) ImagePull(ctx context.Context, ref string, options dockertypes.ImagePullOptions) (io.ReadCloser, error) {
	return f.ImagePullF(ctx, ref, options)
}