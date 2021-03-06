package contract_reboot

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	"github.com/rzry/airdrop/internal/contract_reboot/fflag"
	"github.com/rzry/airdrop/internal/contract_reboot/options"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/blindbox"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/boxpromotion"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/flip"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/gov"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/handingPro"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/lock"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/nft"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/prizepool"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/token"
	log "github.com/rzry/kwstart/kwlog"
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

var once sync.Once

func NewApp(basename string) {
	opts := options.NewOptions(basename)
	fflag.InitFlag()
	once.Do(func() {
		client, err := ethclient.Dial(fflag.Url)
		if err != nil {
			os.Exit(1)
		}
		opts.Client = client
		initInstance(client, opts)
	})
	run(opts)
}

func initInstance(c *ethclient.Client, opts *options.Options) {
	boxer, err := blindbox.NewStore(common.HexToAddress(*fflag.BoxAddr), c)
	boxproer, err := boxpromotion.NewStore(common.HexToAddress(*fflag.BoxProAddr), c)
	fliper, err := flip.NewStore(common.HexToAddress(*fflag.FlipAddr), c)
	gover, err := gov.NewStore(common.HexToAddress(*fflag.GovAddr), c)
	hander, err := handingPro.NewStore(common.HexToAddress(*fflag.HandAddr), c)
	nfter, err := nft.NewStore(common.HexToAddress(*fflag.NftAddr), c)
	pooler, err := prizepool.NewStore(common.HexToAddress(*fflag.PoolAddr), c)
	ptoken, err := token.NewStore(common.HexToAddress(*fflag.PTokenAddr), c)
	locked, err := lock.NewStore(common.HexToAddress(*fflag.LockAddr), c)
	//
	if err != nil {
		log.Fatal("init err", zap.Error(err))
	}
	opts.Ier = &options.Instance{
		Box:     boxer,
		BoxPro:  boxproer,
		Flip:    fliper,
		Gov:     gover,
		Handing: hander,
		Nft:     nfter,
		Pool:    pooler,
		Ptoken:  ptoken,
		Lock:    locked,
		//Ktoken:  ktoken,
	}
	return
}

func run(opts *options.Options) {
	log.Init(opts.Log)
	defer log.Flush()
	prompt := promptui.Select{
		Label: "What do you want to do?",
		Items: []string{
			"???????????????,????????????,????????????",
			"???????????????",
			"???????????????",
			"????????????",
			"????????????",
			"????????????",
			"????????????",
			"????????????",
		},
	}
	_, result, err := prompt.Run()
	if err != nil {
		os.Exit(3)
	}
	ctx := context.Background()
	switch result {
	case "???????????????,????????????,????????????":
		AddBoxPro(opts)
		time.Sleep(3*time.Second)
		CreateSeriesV2(opts,1)
		time.Sleep(3*time.Second)
		TestDraw(opts,1)
	case "???????????????":
		InitOne(opts,ctx)
	case "???????????????":
		AddBoxPro(opts)
	case "????????????":
		CreateSeriesV2(opts,1)
	case "????????????":
		TestDraw(opts,1)
	case "????????????":
		QueryHanding(opts)
	case "????????????":
		SendAward(opts)
	case "????????????":
		Received(opts)
	}

}
