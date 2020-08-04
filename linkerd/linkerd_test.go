package linkerd

import(
	"context"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"testing"
	"github.com/layer5io/meshery-linkerd/meshes"
)

func TestExecuteInstall(t *testing.T){
	os.Setenv("KUBECONFIG", "/Users/aisuko/Documents/rke/kube_config_cluster.yml")
	os.Setenv("CURRENTCONTEXT", "local")
	kubectlConfig := os.Getenv("KUBECONFIG")
	contextName := os.Getenv("CURRENTCONTEXT")
	byteKubeconfig, err := ioutil.ReadFile(kubectlConfig)
	if err != nil {
		t.Fatal(err)
	}
	client, err := newClient(byteKubeconfig, contextName)


	assert.NoError(t,err)
	con:=context.Background()
	arReq:=&meshes.ApplyRuleRequest{
		Namespace:"linkerd",
		DeleteOp:true,
	}
	err=client.executeInstall(con,arReq)

	assert.NoError(t,err)

}