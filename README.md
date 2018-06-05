#### kubelet-version-set

Is run prior to the kubelet starting up, writes a file `/etc/kubernetes/kubelet.env` which is expected to be an environment file used by kubelet.

#### Modes

It run with two modes one with the kubeconfig passed which queries the api server to get the version then adds that to the file as `KUBELET_IMAGE_TAG`. 
The other needs to be passed with a flag `imagever` which sets the image version to the file with the same tag. 
