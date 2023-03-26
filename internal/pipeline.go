package internal

import (
	"context"
	"errors"
	"fmt"
	"regexp"

	"github.com/gnolang/gno/pkgs/bft/rpc/client"
	"github.com/gnolang/gno/pkgs/crypto/keys"
	"github.com/gnolang/supernova/internal/common"
	"github.com/gnolang/supernova/internal/distributor"
	"github.com/schollz/progressbar/v3"
)

var (
	errInvalidURL          = errors.New("invalid node URL specified")
	errInvalidMnemonic     = errors.New("invalid Mnemonic specified")
	errInvalidMode         = errors.New("invalid mode specified")
	errInvalidSubaccounts  = errors.New("invalid number of subaccounts specified")
	errInvalidTransactions = errors.New("invalid number of transactions specified")
	errInvalidBatchSize    = errors.New("invalid batch size specified")
)

var (
	urlRegex = regexp.MustCompile(`(https?://.*)(:(\d*)\/?(.*))?`)
)

type Pipeline struct {
	ctx context.Context

	cfg     *Config
	keybase keys.Keybase
	cli     client.Client
}

// NewPipeline creates a new pipeline instance
func NewPipeline(ctx context.Context, cfg *Config) *Pipeline {
	return &Pipeline{
		ctx:     ctx,
		cfg:     cfg,
		keybase: keys.NewInMemory(),
		cli:     client.NewHTTP(cfg.URL, ""),
	}
}

func (p *Pipeline) Execute() error {
	// Register the accounts with the keybase
	fmt.Println("Generating sub-accounts...")

	accounts := make([]keys.Info, p.cfg.SubAccounts+1)
	bar := progressbar.Default(int64(p.cfg.SubAccounts+1), "sub-accounts added")

	for i := 0; i < int(p.cfg.SubAccounts)+1; i++ {
		info, err := p.keybase.CreateAccount(
			fmt.Sprintf("%s%d", common.KeybasePrefix, i),
			p.cfg.Mnemonic,
			"",
			common.EncryptPassword,
			uint32(0),
			uint32(i),
		)
		if err != nil {
			return fmt.Errorf("unable to create account with keybase, %w", err)
		}

		accounts[i] = info
		_ = bar.Add(1)
	}

	// Distribution //

	fmt.Printf("\n💸 Starting Fund Distribution 💸\n\n")

	_, err := distributor.NewDistributor(
		newBroadcaster(p.cli),
		newStore(p.cli),
		newSigner(p.keybase),
	).Distribute(
		accounts,
		p.cfg.Transactions,
	)
	if err != nil {
		return fmt.Errorf("unable to distribute funds, %w", err)
	}

	return nil
}
