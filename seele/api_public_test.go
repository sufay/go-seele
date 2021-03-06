/**
*  @file
*  @copyright defined in go-seele/LICENSE
 */
package seele

import (
	"bytes"
	"context"
	"os"
	"testing"

	"github.com/seeleteam/go-seele/common"
	"github.com/seeleteam/go-seele/core"
	"github.com/seeleteam/go-seele/crypto"
	"github.com/seeleteam/go-seele/log"
	"github.com/seeleteam/go-seele/node"
)

func getTmpConfig() *node.Config {
	common.IsShardDisabled = true
	acctAddr := crypto.MustGenerateRandomAddress()

	return &node.Config{
		SeeleConfig: node.SeeleConfig{
			TxConf:   *core.DefaultTxPoolConfig(),
			Coinbase: *acctAddr,
		},
	}
}

func Test_PublicSeeleAPI(t *testing.T) {
	conf := getTmpConfig()
	serviceContext := ServiceContext{
		DataDir: common.GetTempFolder(),
	}

	ctx := context.WithValue(context.Background(), "ServiceContext", serviceContext)
	dataDir := ctx.Value("ServiceContext").(ServiceContext).DataDir
	defer os.RemoveAll(dataDir)
	log := log.GetLogger("seele", true)
	ss, err := NewSeeleService(ctx, conf, log)
	if err != nil {
		t.Fatal()
	}

	api := NewPublicSeeleAPI(ss)
	var info MinerInfo
	api.GetInfo(nil, &info)

	if !bytes.Equal(conf.SeeleConfig.Coinbase[0:], info.Coinbase[0:]) {
		t.Fail()
	}
}
