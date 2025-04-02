package operation_v2

import (
	"net/http"

	"github.com/eip-work/kuboard-spray/api/cluster/cluster_common"
	"github.com/eip-work/kuboard-spray/api/command"
	"github.com/eip-work/kuboard-spray/common"
	"github.com/gin-gonic/gin"
)

type ExecuteStepRequest struct {
	Cluster      string `uri:"cluster" binding:"required"`
	Operation    string `uri:"operation" binding:"required"`
	Step         string `uri:"step" binding:"required"`
	Fork         int    `json:"fork"`
	Verbose      string `json:"verbose"`
	ExcludeNodes string `json:"nodes_to_exclude"`
}

func ExecuteStep(c *gin.Context) {
	var req ExecuteStepRequest
	c.ShouldBindUri(&req)
	c.ShouldBindJSON(&req)

	inventory, _, err := updateResourcePackageVarsToInventory(req)

	if err != nil {
		common.HandleError(c, http.StatusInternalServerError, "failed to process inventory", err)
		return
	}

	postExec := func(status command.ExecuteExitStatus) (string, error) {
		return "\n执行成功", nil
	}

	playbook := "operations/" + req.Operation + "/" + req.Step + "/playbook.yaml"

	cmd := command.Execute{
		OwnerType: "cluster",
		OwnerName: req.Cluster,
		Cmd:       "ansible-playbook",
		Args: func(execute_dir string) []string {
			result := []string{"-i", execute_dir + "/inventory.yaml", playbook}
			result = appendCommonParams(result, req, false)
			return result
		},
		Dir:      cluster_common.ResourcePackageDirForInventory(inventory),
		Type:     req.Operation,
		PreExec:  func(execute_dir string) error { return common.SaveYamlFile(execute_dir+"/inventory.yaml", inventory) },
		PostExec: postExec,
	}

	if err := cmd.Exec(); err != nil {
		common.HandleError(c, http.StatusInternalServerError, "Faild to InstallCluster. ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"pid": cmd.R_Pid,
		},
	})
}
