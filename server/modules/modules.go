

package modules

import (
  "github.com/threatcode/threatcode-soc/module"
  "github.com/threatcode/threatcode-soc/server"
  "github.com/threatcode/threatcode-soc/server/modules/elastic"
  "github.com/threatcode/threatcode-soc/server/modules/elasticcases"
  "github.com/threatcode/threatcode-soc/server/modules/filedatastore"
  "github.com/threatcode/threatcode-soc/server/modules/generichttp"
  "github.com/threatcode/threatcode-soc/server/modules/influxdb"
  "github.com/threatcode/threatcode-soc/server/modules/kratos"
  "github.com/threatcode/threatcode-soc/server/modules/sostatus"
  "github.com/threatcode/threatcode-soc/server/modules/statickeyauth"
  "github.com/threatcode/threatcode-soc/server/modules/staticrbac"
  "github.com/threatcode/threatcode-soc/server/modules/thehive"
)

func BuildModuleMap(srv *server.Server) map[string]module.Module {
  moduleMap := make(map[string]module.Module)
  moduleMap["filedatastore"] = filedatastore.NewFileDatastore(srv)
  moduleMap["httpcase"] = generichttp.NewHttpCase(srv)
  moduleMap["influxdb"] = influxdb.NewInfluxDB(srv)
  moduleMap["kratos"] = kratos.NewKratos(srv)
  moduleMap["elastic"] = elastic.NewElastic(srv)
  moduleMap["elasticcases"] = elasticcases.NewElasticCases(srv)
  moduleMap["sostatus"] = sostatus.NewSoStatus(srv)
  moduleMap["statickeyauth"] = statickeyauth.NewStaticKeyAuth(srv)
  moduleMap["staticrbac"] = staticrbac.NewStaticRbac(srv)
  moduleMap["thehive"] = thehive.NewTheHive(srv)
  return moduleMap
}
