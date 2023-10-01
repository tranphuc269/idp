package graphql

import "idp_system/app/framework/webreq"

func NewClientFactoryFake(handleFunc webreq.TransportHandleFunc) ClientFactory {
	return NewClientFactory(webreq.NewHTTPFake(handleFunc))
}
