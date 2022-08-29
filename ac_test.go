package goAcAutoMachine

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestAc(t *testing.T) {
	content := "1取消拦截2JT11111111，JT11111111，JT11111111取消拦截，JT7778888改地址，JT7778888改地址，  JT7778888催促,JT7778888改地址， ,JT121233444取消改地址 ,111查重量 ,111查询重量,,JT11111111消拦截，"
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
