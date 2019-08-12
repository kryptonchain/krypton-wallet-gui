// Copyright (c) 2018, The TurtleCoin Developers
// Copyright (c) 2019, The KryptonChain Developers
//
// Please see the included LICENSE file for more information.
//

package walletdmanager

const (
	// DefaultTransferFee is the default fee. It is expressed in ZOD
	DefaultTransferFee float64 = 0.1

	logWalletdCurrentSessionFilename     = "krypton-service-session.log"
	logWalletdAllSessionsFilename        = "krypton-service.log"
	logTurtleCoindCurrentSessionFilename = "kryptond-session.log"
	logTurtleCoindAllSessionsFilename    = "kryptond.log"
	walletdLogLevel                      = "3" // should be at least 3 as I use some logs messages to confirm creation of wallet
	walletdCommandName                   = "krypton-cli"
	turtlecoindCommandName               = "kryptond"
)
