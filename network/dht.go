package network

import (
	"context"

	lp2pcore "github.com/libp2p/go-libp2p-core"
	lp2phost "github.com/libp2p/go-libp2p-core/host"
	lp2pdht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/zarbchain/zarb-go/logger"
)

type dhtService struct {
	ctx       context.Context
	host      lp2phost.Host
	kademlia  *lp2pdht.IpfsDHT
	bootstrap *bootstrap
	logger    *logger.Logger
}

func newDHTService(ctx context.Context, host lp2phost.Host, protocolID lp2pcore.ProtocolID, conf *BootstrapConfig, logger *logger.Logger) *dhtService {
	opts := []lp2pdht.Option{
		lp2pdht.Mode(lp2pdht.ModeAuto),
		lp2pdht.ProtocolPrefix(protocolID),
	}

	kademlia, err := lp2pdht.New(ctx, host, opts...)
	if err != nil {
		logger.Panic("Unable to start DHT service: %v", err)
		return nil
	}

	bootstrap := newBootstrap(ctx,
		host, host.Network(), kademlia,
		conf, logger)

	return &dhtService{
		ctx:       ctx,
		host:      host,
		kademlia:  kademlia,
		bootstrap: bootstrap,
		logger:    logger,
	}
}

func (dht *dhtService) Start() error {
	dht.bootstrap.Start()
	return nil
}

func (dht *dhtService) Stop() {
	if err := dht.kademlia.Close(); err != nil {
		dht.logger.Error("Unable to close Kademlia", "err", err)
	}

	dht.bootstrap.Stop()
}