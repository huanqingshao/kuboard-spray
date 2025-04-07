package operation_v2

import (
	"net/http"
	"os"

	"github.com/eip-work/kuboard-spray/common"
	"github.com/eip-work/kuboard-spray/constants"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ListExecuteHistoryRequest struct {
	Cluster   string `uri:"cluster" binding:"required"`
	Operation string `uri:"operation" binding:"required"`
	Step      string `uri:"step" binding:"required"`
}

type HistoryItem struct {
	Time   string `json:"time" binding:"required"`
	Status string `json:"status" binding:"required"`
}

func ListOperationStepHistory(c *gin.Context) {
	var req ListExecuteHistoryRequest
	c.ShouldBindUri(&req)
	c.ShouldBindJSON(&req)

	historyDir := constants.GET_DATA_CLUSTER_DIR() + "/" + req.Cluster + "/history/" + req.Operation + "/" + req.Step
	logrus.Info(historyDir)

	result := []HistoryItem{}
	if common.PathExists(historyDir) {
		entries, err := os.ReadDir(historyDir)
		if err != nil {
			common.HandleError(c, http.StatusInternalServerError, "cannot read dir: "+historyDir, err)
			return
		}
		for _, entry := range entries {
			if entry.IsDir() {
				result = append(result, HistoryItem{
					Time:   entry.Name(),
					Status: "unknown",
				})
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "success",
		"data": gin.H{
			"history": result,
		},
	})

}
