package repl

import (
	"encoding/hex"
	"fmt"

	"github.com/spacemeshos/CLIWallet/log"
	apitypes "github.com/spacemeshos/api/release/go/spacemesh/v1"
	"github.com/spacemeshos/ed25519"
	gosmtypes "github.com/spacemeshos/go-spacemesh/common/types"
)

func (r *repl) chooseAccount() {
	accs := r.client.ListAccounts()
	if len(accs) == 0 {
		r.createAccount()
		return
	}

	fmt.Println(printPrefix, "Choose an account to load:")
	accName := multipleChoice(accs)
	account, err := r.client.GetAccount(accName)
	if err != nil {
		panic("wtf")
	}
	fmt.Printf("%s Loaded account alias: `%s`, address: %s \n", printPrefix, account.Name, account.Address().String())

	r.client.SetCurrentAccount(account)
}

func (r *repl) createAccount() {
	fmt.Println(printPrefix, "Create a new account")
	alias := inputNotBlank(createAccountMsg)

	ac := r.client.CreateAccount(alias)
	err := r.client.StoreAccounts()
	if err != nil {
		log.Error("failed to create account: %v", err)
		return
	}

	fmt.Printf("%s Created account alias: `%s`, address: %s \n", printPrefix, ac.Name, ac.Address().String())
	r.client.SetCurrentAccount(ac)
}

func (r *repl) accountInfo() {
	acc := r.client.CurrentAccount()
	if acc == nil {
		r.chooseAccount()
		acc = r.client.CurrentAccount()
	}

	address := gosmtypes.BytesToAddress(acc.PubKey)

	state, err := r.client.AccountState(address)
	if err != nil {
		log.Error("failed to get account info: %v", err)
		return
	}

	fmt.Println(printPrefix, "Local alias:", acc.Name)
	fmt.Println(printPrefix, "Address:", address.String())
	fmt.Println(printPrefix, "Balance:", state.StateCurrent.Balance.Value, coinUnitName)
	fmt.Println(printPrefix, "Nonce:", state.StateCurrent.Counter)
	fmt.Println(printPrefix, "Projected Balance:", state.StateProjected.Balance.Value, coinUnitName)
	fmt.Println(printPrefix, "Projected Nonce:", state.StateProjected.Counter)

	fmt.Println(printPrefix, fmt.Sprintf("Public key: 0x%s", hex.EncodeToString(acc.PubKey)))
	fmt.Println(printPrefix, fmt.Sprintf("Private key: 0x%s", hex.EncodeToString(acc.PrivKey)))
}

// printAccountRewards prints all rewards awarded to an account
func (r *repl) printAccountRewards() {
	acc := r.client.CurrentAccount()
	if acc == nil {
		r.chooseAccount()
		acc = r.client.CurrentAccount()
	}

	// todo: request offset and total from user
	rewards, total, err := r.client.AccountRewards(acc.Address(), 0, 10000)
	if err != nil {
		log.Error("failed to list transactions: %v", err)
		return
	}

	fmt.Println(printPrefix, fmt.Sprintf("Total rewards: %d", total))
	for _, r := range rewards {
		printReward(r)
		fmt.Println(printPrefix, "-----")
	}
}

func printReward(r *apitypes.Reward) {

}

func (r *repl) sign() {
	acc := r.client.CurrentAccount()
	if acc == nil {
		r.chooseAccount()
		acc = r.client.CurrentAccount()
	}

	msgStr := inputNotBlank(msgSignMsg)
	msg, err := hex.DecodeString(msgStr)
	if err != nil {
		log.Error("failed to decode msg hex string: %v", err)
		return
	}

	signature := ed25519.Sign2(acc.PrivKey, msg)

	fmt.Println(printPrefix, fmt.Sprintf("signature (in hex): %x", signature))
}

func (r *repl) textsign() {
	acc := r.client.CurrentAccount()
	if acc == nil {
		r.chooseAccount()
		acc = r.client.CurrentAccount()
	}

	msg := inputNotBlank(msgTextSignMsg)
	signature := ed25519.Sign2(acc.PrivKey, []byte(msg))

	fmt.Println(printPrefix, fmt.Sprintf("signature (in hex): %x", signature))
}

/*
func (r *repl) unlockAccount() {
	passphrase := r.commandLineParams(1, r.input)
	err := r.client.Unlock(passphrase)
	if err != nil {
		log.Debug(err.Error())
		return
	}

	acctCmd := r.commands[3]
	r.executor(fmt.Sprintf("%s %s", acctCmd.text, passphrase))
}

func (r *repl) lockAccount() {
	passphrase := r.commandLineParams(2, r.input)
	err := r.client.Lock(passphrase)
	if err != nil {
		log.Debug(err.Error())
		return
	}

	acctCmd := r.commands[3]
	r.executor(fmt.Sprintf("%s %s", acctCmd.text, passphrase))
}*/
