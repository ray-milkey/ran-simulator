module github.com/onosproject/ran-simulator

go 1.16

require (
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cenkalti/backoff/v4 v4.1.0
	github.com/garyburd/redigo v1.1.1-0.20170914051019-70e1b1943d4f // indirect
	github.com/google/uuid v1.2.0
	github.com/gregjones/httpcache v0.0.0-20190611155906-901d90724c79 // indirect
	github.com/onosproject/helmit v0.6.19
	github.com/onosproject/onos-api/go v0.8.0
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm v0.8.4
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_kpm_v2_go v0.8.4
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_mho_go v0.8.4
	github.com/onosproject/onos-e2-sm/servicemodels/e2sm_rc_pre_go v0.8.4
	github.com/onosproject/onos-e2t v0.10.3
	github.com/onosproject/onos-lib-go v0.8.11
	github.com/onosproject/onos-ric-sdk-go v0.8.0
	github.com/onosproject/onos-test v0.6.4
	github.com/onosproject/rrm-son-lib v0.0.2
	github.com/pmcxs/hexgrid v0.0.0-20190126214921-42796ac894ab
	github.com/spf13/cobra v1.2.1
	github.com/spf13/viper v1.9.0
	github.com/stretchr/testify v1.7.0
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	googlemaps.github.io/maps v1.3.2
	gopkg.in/yaml.v2 v2.4.0
	gotest.tools v2.2.0+incompatible
)

replace github.com/docker/docker => github.com/docker/engine v1.4.2-0.20200229013735-71373c6105e3

replace github.com/pmcxs/hexgrid v0.0.0-20190126214921-42796ac894ab => github.com/SeanCondon/hexgrid v0.0.0-20200424141352-c3819a378a18
