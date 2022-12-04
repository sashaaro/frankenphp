package httpclient

// #cgo CFLAGS: -I/usr/local/include/php -I/usr/local/include/php/Zend -I/usr/local/include/php/TSRM -I/usr/local/include/php/main
// #include <stdlib.h>
// #include <stdint.h>
// #include <php_variables.h>
import "C"
import (
	"bufio"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/dunglas/frankenphp/plugins"
	"golang.org/x/net/http2"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var pluginName = "httpclient"

type Plugin struct {
	// http2clients
	clients map[string]*http.Client
}

func (plugin Plugin) Name() string {
	return pluginName
}

var defaultClientName = "default"

func (plugin Plugin) Init(config interface{}) error {
	// TODO create multiple clients and replace default arguments to value from config

	// TODO make from config

	// if config is empty - create default client
	clientConfigs := []*HttpClientConfig{
		&HttpClientConfig{
			name: defaultClientName,
			// insecureSkipVerify: true,
			allowHTTP: true,
		},
	}

	plugin.clients = make(map[string]*http.Client)

	for _, conf := range clientConfigs {
		plugin.clients[conf.name] = CreateFrankenHttp2Client(conf)
	}

	return nil
}

type HttpClientConfig struct {
	name               string
	insecureSkipVerify bool
	allowHTTP          bool
}

func CreateFrankenHttp2Client(config *HttpClientConfig) *http.Client {
	transport := &http2.Transport{
		AllowHTTP: config.allowHTTP,
	}
	transport.DialTLSContext = func(ctx context.Context, network, addr string, cfg *tls.Config) (net.Conn, error) {
		return net.Dial(network, addr)
	}

	if config.insecureSkipVerify {
		transport.TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	return &http.Client{
		Transport: transport,
	}
}

//export go_frankenphp_client_send_request
func go_frankenphp_client_send_request(request *C.char) *C.char {
	httpClientModule := plugins.FrankenPhpPlugins[pluginName]
	module, ok := httpClientModule.(Plugin)
	if !ok {
		panic("TODO")
	}

	http2Client := module.clients[defaultClientName] // take from argument

	if http2Client == nil {
		//fc.Logger.Error(fmt.Sprintf("http.client Frankenphp http client not exist."))
		// TODO exception
		return C.CString(fmt.Sprintf("Error Frankenphp httpclient is not supported %s"))
	}

	raw := C.GoString(request)

	req, err := http.ReadRequest(bufio.NewReader(strings.NewReader(raw)))
	//req, err := http.NewRequest("GET", "https://httpbin.org/headers", nil)

	if err != nil {
		//fc.Logger.Error(fmt.Sprintf("http.client failed to parse request. %s. Raw request:\n%s", err.Error(), raw))
		// todo throw exception
		return C.CString(fmt.Sprintf("Error parsing request: %s. Raw request:\n%s", err.Error(), raw))
	}

	req.RequestURI = ""
	req.URL.Host = req.Host

	if req.URL.Scheme == "" {
		req.URL.Scheme = "https"
	}

	// req.ProtoMajor = 2
	// req.ProtoMinor = 0

	//fc.Logger.Debug(fmt.Sprintf("http.client request %s %s", req.Method, req.URL.String()))

	res, err := http2Client.Do(req)
	if err != nil {
		//fc.Logger.Error(fmt.Sprintf("http.client Failed do request. %s", err.Error()))

		return C.CString(fmt.Sprintf("Error failed request: %s", err.Error()))
	}
	//fc.Logger.Debug(fmt.Sprintf("http.client Response status %s", res.StatusCode))

	b, err := httputil.DumpResponse(res, true)

	if err != nil {
		// TODO
	}

	return C.CString(string(b))
}
