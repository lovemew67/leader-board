package controllerv1

import (
	"log"
	"os"
	"testing"

	"github.com/lovemew67/public-misc/cornerstone"
)

var (
	ctx = cornerstone.NewContext()
)

func BeforeTest() {}

func TestMain(m *testing.M) {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.LstdFlags)
	BeforeTest()
	retCode := m.Run()
	AfterTest()
	os.Exit(retCode)
}

func AfterTest() {}
