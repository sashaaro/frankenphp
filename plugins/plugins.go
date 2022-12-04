package plugins

var FrankenPhpPlugins = make(map[string]FrankenPhpPlugin)

type FrankenPhpPlugin interface {
	Name() string
	Init(config interface{}) error
	// TODO OnRequest(request *http.Request) error
	// TODO OnResponse(response *http.Response) error
}
