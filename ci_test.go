package test

import (
	"fmt"
	"testing"

	"gopkg.in/check.v1"
)

func Test(t *testing.T) { check.TestingT(t) }

var _ = check.Suite(&TestSuite{})

type TestSuite struct {
}

func (cs *TestSuite) SetUpSuite(c *check.C) {
}

func (cs *TestSuite) TearDownSuite(c *check.C) {
}
