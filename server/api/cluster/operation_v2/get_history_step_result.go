package operation_v2

import (
	"net/http"

	"github.com/eip-work/kuboard-spray/common"
	"github.com/eip-work/kuboard-spray/constants"
	"github.com/gin-gonic/gin"
)

type GetHistoryStepResultRequest struct {
	Cluster   string `uri:"cluster" binding:"required"`
	Operation string `uri:"operation" binding:"required"`
	Step      string `uri:"step" binding:"required"`
	Time      string `uri:"time" binding:"required"`
}

func GetHistoryStepResult(c *gin.Context) {
	var req GetHistoryStepResultRequest
	c.ShouldBindUri(&req)
	c.ShouldBindJSON(&req)

	historyDir := constants.GET_DATA_CLUSTER_DIR() + "/" + req.Cluster + "/history/" + req.Operation + "/" + req.Step + "/" + req.Time + "/result.yaml"

	result, err := common.ParseYamlFile(historyDir)
	if err == nil {
		c.JSON(http.StatusOK, common.KuboardSprayResponse{
			Code:    http.StatusOK,
			Message: "success",
			Data:    result,
		})
	} else {
		common.HandleError(c, http.StatusInternalServerError, "Cannot find history result", err)
	}
}
