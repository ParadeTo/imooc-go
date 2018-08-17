package config

const (
	ItemSaverPort = 1234
	WorkerPort0   = 9000

	ElasticIndex = "dating_profile"

	ItemSaverRpc = "ItemSaverService.Save"
	CrawlServiceRpc = "CrawlService.Process"

	// Parser names
	ParseCity     = "ParseCity"
	ParseCityList = "ParseCityList"
	ParseProfile  = "ParseProfile"
	NilParser     = "NilParser"
)
