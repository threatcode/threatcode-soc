

package elasticcases

import (
	"github.com/threatcode/threatcode-soc/model"
)

func convertToElasticCase(inputCase *model.Case) (*ElasticCase, error) {
	outputCase := NewElasticCase()
	outputCase.Title = inputCase.Title
	outputCase.Description = inputCase.Description
	outputCase.Tags = append(outputCase.Tags, "SecurityOnion")
	outputCase.Connector = &ElasticConnector{
		Id:     "none",
		Name:   "none",
		Type:   ".none",
		Fields: nil,
	}
	outputCase.Settings = &ElasticSettings{
		SyncAlerts: true,
	}
	outputCase.Owner = "securitySolution"
	return outputCase, nil
}

func convertFromElasticCase(inputCase *ElasticCase) (*model.Case, error) {
	outputCase := model.NewCase()
	outputCase.Title = inputCase.Title
	outputCase.Description = inputCase.Description
	outputCase.Id = inputCase.Id
	outputCase.Status = inputCase.Status
	if inputCase.CreatedDate != nil {
		outputCase.CreateTime = inputCase.CreatedDate
	}
	if inputCase.ModifiedDate != nil {
		outputCase.UpdateTime = inputCase.ModifiedDate
	}
	if inputCase.ClosedDate != nil {
		outputCase.CompleteTime = inputCase.ClosedDate
	}
	return outputCase, nil
}
