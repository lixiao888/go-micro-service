module git.imooc.com/coding-447/paymentApi

go 1.15

require (
	git.imooc.com/coding-447/common v0.0.0-00010101000000-000000000000
	git.imooc.com/coding-447/payment v0.0.0-00010101000000-000000000000
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/golang/protobuf v1.4.0
	github.com/micro/go-micro/v2 v2.9.1
	github.com/micro/go-plugins/registry/consul/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/select/roundrobin/v2 v2.9.1
	github.com/micro/go-plugins/wrapper/trace/opentracing/v2 v2.9.1
	github.com/opentracing/opentracing-go v1.2.0
	github.com/plutov/paypal/v3 v3.1.0
)

replace git.imooc.com/coding-447/common => ../common

replace git.imooc.com/coding-447/payment => ../payment
