package core

import "fmt"

// state hold variable of smart contract as key,value pair
type state struct {
	data map[string][]byte
}

// create new state
func (s *state) NewState() *state {
	return &state{
		data: map[string][]byte{},
	}
}

//put data in state
func (s *state) Put(key, val []byte) error {
	s.data[string(key)] = val
	return nil
}

// delete data from state
func (s *state) Delete(key []byte) error {
	delete(s.data, string(key))
	return nil
}

// search data from state
func (s *state) Get(key []byte) ([]byte, error) {
	val, ok := s.data[string(key)]
	if !ok {
		return nil, fmt.Errorf("given key %d not found", key)
	}
	return val, nil

}
