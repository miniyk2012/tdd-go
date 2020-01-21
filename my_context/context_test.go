package my_context

import "testing"


type StubStore struct {
	response string
}

func (s *StubStore) Fetch() string {
	return s.response
}

func TestHandler(t *testing.T) {

}
