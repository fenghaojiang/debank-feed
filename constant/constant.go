package constant

import "github.com/ethereum/go-ethereum/crypto"

var (
	TranferEvent    = crypto.Keccak256Hash([]byte("Transfer(address,address,uint256)"))
	ApprovalEvent   = crypto.Keccak256Hash([]byte("Approval(address,address,uint256)"))
	WithdrawalEvent = crypto.Keccak256Hash([]byte("Withdrawal(address,uint256)"))
	DepositEvent    = crypto.Keccak256Hash([]byte("Deposit(address,uint256)"))
)
