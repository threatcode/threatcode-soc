

package elasticcases

import (
  "context"
  "github.com/threatcode/threatcode-soc/model"
  "github.com/threatcode/threatcode-soc/server"
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestCreateUnauthorized(tester *testing.T) {
  casestore := NewElasticCasestore(server.NewFakeUnauthorizedServer())
  casestore.Init("some/url", "someusername", "somepassword", true)
  socCase := model.NewCase()
  newCase, err := casestore.Create(context.Background(), socCase)
  assert.Error(tester, err)
  assert.Nil(tester, newCase)
}

func TestCreate(tester *testing.T) {
  casestore := NewElasticCasestore(server.NewFakeAuthorizedServer(nil))
  casestore.Init("some/url", "someusername", "somepassword", true)
  caseResponse := `
    {
      "id": "a123",
      "title": "my title"
    }`
  casestore.client.MockStringResponse(caseResponse, 200, nil)
  socCase := model.NewCase()
  newCase, err := casestore.Create(context.Background(), socCase)
  assert.NoError(tester, err)

  assert.Equal(tester, "my title", newCase.Title)
  assert.Equal(tester, "a123", newCase.Id)
}
