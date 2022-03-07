package domainv1

import (
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
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

func Test_All(t *testing.T) {
	req1 := InsertScoreV1ServiceRequest{}
	assert.NotNil(t, req1.String())

	req2 := ListTopKScoresV1ServiceRequest{}
	assert.NotNil(t, req2.String())
}
