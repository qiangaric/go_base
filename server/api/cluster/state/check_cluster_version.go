package state

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/eip-work/kuboard-spray/api/ansible_rpc"
	"github.com/eip-work/kuboard-spray/api/cluster/cluster_common"
	"github.com/eip-work/kuboard-spray/api/command"
	"github.com/eip-work/kuboard-spray/common"
	"github.com/eip-work/kuboard-spray/constants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ClusterVersion map[string](map[string]ansible_rpc.AnsibleResultNode)

func CheckClusterVersion(c *gin.Context) {

	var request GetNodesRequest
	c.ShouldBindUri(&request)

	startTime := time.Now()

	cluster, err := cluster_common.ClusterMetadataByName(request.ClusterName)
	if err != nil {
		common.HandleError(c, http.StatusInternalServerError, "failed to get cluster metadata", err)
		return
	}

	playbook := common.MapGetString(cluster.ResourcePackage, "data.supported_playbooks.cluster_version_"+common.MapGetString(cluster.Inventory, "all.children.target.vars.container_manager"))
	if playbook == "" {
		common.HandleError(c, http.StatusInternalServerError, "current resource package doesnot support the operation to view cluster version", nil)
		return
	}

	cmd := command.Run{
		Cmd: "ansible-playbook",
		Args: []string{
			cluster.ResourcePackageDir + "/" + playbook,
			"-i", cluster.InventoryPath,
			"--forks", "200",
			"--timeout", "3",
			"-e", "kuboardspray_ssh_args='-o ConnectionAttempts=1 -o UserKnownHostsFile=/dev/null -F /dev/null'",
		},
		Env:     []string{"ANSIBLE_CONFIG=./ansible.cfg", "ANSIBLE_CACHE_PLUGIN_CONNECTION=" + constants.GET_DATA_CLUSTER_DIR() + "/" + request.ClusterName + "/fact"},
		Timeout: 60,
		Dir:     "./ansible-rpc",
	}

	stdout, stderr, err := cmd.Run()
	duration := time.Now().UnixNano() - startTime.UnixNano()
	logrus.Trace("duration: ", duration/1000000)
	if err != nil {
		common.HandleError(c, http.StatusInternalServerError, "failed to run", err)
		return
	}

	// logrus.Trace("stdout: ", string(stdout), "\nstderr: ", string(stderr))

	result := &ansible_rpc.AnsibleResult{}
	if err := json.Unmarshal(stdout, result); err != nil {
		common.HandleError(c, http.StatusInternalServerError, "failed to Unmarshal result: ["+string(stdout)+"]", err)
		logrus.Trace("stdout: ", string(stdout), "\nstderr: ", string(stderr))
		logrus.Trace("duration: ", duration/1000000)
		return
	}

	clusterVersion := ClusterVersion{}

	for _, task := range result.Plays[0].Tasks {
		for nodeName, node := range task.Hosts {
			if clusterVersion[nodeName] == nil {
				clusterVersion[nodeName] = make(map[string]ansible_rpc.AnsibleResultNode)
			}
			clusterVersion[nodeName][task.Task.Name] = node
		}
	}

	c.JSON(http.StatusOK, common.KuboardSprayResponse{
		Code:    http.StatusOK,
		Message: "success",
		Data:    clusterVersion,
	})
}
