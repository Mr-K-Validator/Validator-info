package cmd

import (
        common "github.com/Mr-K-Validator/Validator-info/rest/common"
)

func set_config() {

	common.Addr = restAddr
        common.OperAddr = operAddr

}
