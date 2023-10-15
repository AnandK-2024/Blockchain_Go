package core

import (
	"errors"
	"fmt"
	"sync"

	"github.com/AnandK-2024/Blockchain/types"
)

var (
	ErrorAccountNotFound        = errors.New("account not found")
	ErrorInsufficientBalance = errors.New("insufficient  account balance")
)

// a account should have address and balance
type Account struct {
	address types.Address
	balance uint64
}

func (a Account) String() string {
	return fmt.Sprintf("%d", a.balance)
}

// accountstate that store accounts of users
type AccountState struct {
	lock    sync.RWMutex
	account map[types.Address]*Account
}

// create a new account state that manage all accounts
func NewAccountState() *AccountState {
	return &AccountState{
		account: make(map[types.Address]*Account),
	}
}

// create new account
func (s *AccountState) CreateAccount(address types.Address) *Account {
	s.lock.Lock()
	defer s.lock.Unlock()
	acc := &Account{address: address, balance: 0}
	s.account[address] = acc
	return acc
}

// get account details with address
func (s *AccountState) GetAccount(address types.Address) (*Account, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.getAccountWithoutLock(address)
}

func (s *AccountState) getAccountWithoutLock(address types.Address) (*Account, error) {
	account, ok := s.account[address]
	if !ok {
		return nil, ErrorAccountNotFound
	}
	return account, nil
}

// get balance of accout with address
func (s *AccountState) GetBalance(address types.Address) (uint64, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	account, ok := s.getAccountWithoutLock(address)
	if ok != nil {
		return 0, ErrorAccountNotFound
	}
	return account.balance, nil

}

// transfer crypto from one account to another account
func (s *AccountState) Transfer(from types.Address, to types.Address, amount uint64) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	// get account of sender
	sender, err := s.getAccountWithoutLock(from)
	if err != nil {
		return err
	}

	if sender.balance < amount {
		return ErrorInsufficientBalance
	}
	if sender.balance != 0 {
		sender.balance -= amount
	}

	// get account of reciever
	reciever, err := s.getAccountWithoutLock(to)
	if err != nil {
		// if account not found then create a new account for reciever
		s.CreateAccount(to)
		// get account for reciever
		reciever, _ = s.getAccountWithoutLock(to)
	}
	reciever.balance += amount
	return nil

}
