package goAcAutoMachine

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestAc(t *testing.T) {
	content := "1取消拦截2"
	ac := NewAcAutoMachine()
	ac.AddPattern("红领巾")
	ac.AddPattern("拦截")
	ac.AddPattern("取消拦截")
	ac.AddPattern("取消拦")
	ac.Build()
	results := ac.Query(content)
	fmt.Println("内容: " + content)
	for _, result := range results {
		fmt.Println(result)
	}
}

func TestObj(t *testing.T) {
	var result Result
	result = Result{}
	log.Infof("%v", result)
	result = Result{}
	log.Infof("%v", result)

	result.Start = 1
}
