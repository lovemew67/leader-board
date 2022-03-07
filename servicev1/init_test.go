package servicev1

import (
	"log"
	"os"
	"testing"
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
