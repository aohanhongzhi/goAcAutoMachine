package goAcAutoMachine

import (
	"fmt"
	"strings"
	"testing"
)

func TestAc(t *testing.T) {
	content := "JT11111111，JT11111111，JT11111111取消拦截，JT7778888改地址，JT7778888改地址，  JT7778888催促,JT7778888改地址， ,JT121233444取消改地址 ,111查重量 ,111查询重量,,JT11111111消拦截，"
	//content := "JT1取消拦截,JT44取消改地址"
	//content := "JT11111111JT11111111撤消改地址，JT11111111，JT11111111取消拦截，JT7778888改地址，JT7778888改地址，  JT7778888催促,JT7778888改地址， ,JT121233444取消改地址 ,111查重量 ,111查询重量,,JT11111111消拦截，"
	//content := "JT11111111撤,JT11111111撤消改地址,JT11111111撤,JT11111111撤消改地址"
	//content = "JT11111111撤,JT11111111撤消改地，,JT11111111撤,JT11111111撤消改地址"

	//content := "JT11111111JT11111111撤消改地址，JT11111111，JT11111111取消拦截，JT7778888改地址，JT7778888改地址，  JT7778888催促,JT7778888改地址， ,JT121233444取消改地址 ,111查重量 ,111查询重量,,JT11111111消拦截，"
	//content := "JT11111111撤,JT11111111撤消改地址,JT11111111撤,JT11111111撤消改地址"
	//content = "JT11111111撤,JT11111111撤消改地，,JT11111111撤,JT11111111撤消改地址"
	content = "取消改地址"

	ac := NewAcAutoMachine()
	ac.AddPattern("拦截")
	ac.AddPattern("取消拦截")
	ac.AddPattern("取消拦")
	ac.AddPattern("改地址")

	ac.AddPattern("撤")
	ac.AddPattern("撤消改")
	ac.AddPattern("撤消改地址")
	ac.AddPattern("取消改地址")
	ac.AddPattern("取消改地")
	ac.AddPattern("消改地")
	ac.AddPattern("重量")
	ac.AddPattern("查重量")
	ac.AddPattern("催促")
	ac.AddPattern("撤")
	ac.Build()
	results := ac.Query(content)
	fmt.Println("内容: " + content)
	for _, result := range results {
		fmt.Println(result)
	}

	//results1 := ac.QueryLast(content)
	//fmt.Println("=============")
	//for _, result := range results1 {
	//	fmt.Println(result)
	//}
}

func TestAcAndContain(t *testing.T) {

	var whyWord = []string{}

	whyWord = append(whyWord, "能不能派送")
	whyWord = append(whyWord, "是不是丢了")
	whyWord = append(whyWord, "停发了吗")
	whyWord = append(whyWord, "还没有退回")
	whyWord = append(whyWord, "还没退回")
	whyWord = append(whyWord, "为何退回")
	whyWord = append(whyWord, "为什么退回")
	whyWord = append(whyWord, "什么退回")
	whyWord = append(whyWord, "怎么退回")
	whyWord = append(whyWord, "怎么回事")
	whyWord = append(whyWord, "什么情况")
	whyWord = append(whyWord, "什么原因")
	whyWord = append(whyWord, "查物流")
	whyWord = append(whyWord, "物流信息")
	whyWord = append(whyWord, "更新物流")
	whyWord = append(whyWord, "查快递")
	whyWord = append(whyWord, "查一下")
	whyWord = append(whyWord, "位置")
	whyWord = append(whyWord, "快递到哪")
	whyWord = append(whyWord, "核实物流")
	whyWord = append(whyWord, "核实一下物流")
	whyWord = append(whyWord, "核实一下信息")
	whyWord = append(whyWord, "核实退回原因")
	whyWord = append(whyWord, "核实退回")
	whyWord = append(whyWord, "核实好了吗")
	whyWord = append(whyWord, "核实什么情况没有转出")
	whyWord = append(whyWord, "不走件")
	whyWord = append(whyWord, "物流不动")
	whyWord = append(whyWord, "物流不更新")
	whyWord = append(whyWord, "没更新物流")
	whyWord = append(whyWord, "天没更新") // 六天没更新
	whyWord = append(whyWord, "不更新物流信息")
	whyWord = append(whyWord, "这个件啥情况")
	whyWord = append(whyWord, "是丢了吗")
	whyWord = append(whyWord, "核实了吗")
	whyWord = append(whyWord, "退回了吗")
	whyWord = append(whyWord, "这个退回了")
	whyWord = append(whyWord, "退回了吧")
	whyWord = append(whyWord, "退回来了吧")
	whyWord = append(whyWord, "退回来了吗")
	whyWord = append(whyWord, "有回复吗")
	whyWord = append(whyWord, "拦截成功了吗")
	whyWord = append(whyWord, "是否拦截")
	whyWord = append(whyWord, "是否已经拦截")
	whyWord = append(whyWord, "是否已经退回")
	whyWord = append(whyWord, "是否拦截")
	whyWord = append(whyWord, "咋回事")
	whyWord = append(whyWord, "停发吗")
	whyWord = append(whyWord, "是否退回")
	whyWord = append(whyWord, "显示取消寄件")
	whyWord = append(whyWord, "中转错了吗")
	whyWord = append(whyWord, "错分了吗")
	whyWord = append(whyWord, "咋疑难")
	whyWord = append(whyWord, "疑难核实")
	whyWord = append(whyWord, "核实疑难")
	whyWord = append(whyWord, "核实下疑难件")
	whyWord = append(whyWord, "疑难原因")
	whyWord = append(whyWord, "查的怎么样了")
	whyWord = append(whyWord, "退了吗")
	whyWord = append(whyWord, "退回单号")
	whyWord = append(whyWord, "收件人没收到")
	whyWord = append(whyWord, "没收到") // FIXME  没收到 退回 。这句话，应该强调退回
	whyWord = append(whyWord, "未收到")
	whyWord = append(whyWord, "是在退回")   // 是在退回吗
	whyWord = append(whyWord, "没找到")    // 是在退回吗
	whyWord = append(whyWord, "没有退回信息") // 没有退回信息
	whyWord = append(whyWord, "正常在运输吗")
	whyWord = append(whyWord, "在正常运输吗")
	whyWord = append(whyWord, "在运输吗")
	whyWord = append(whyWord, "退回么")
	whyWord = append(whyWord, "退回了么")
	whyWord = append(whyWord, "不动了")
	whyWord = append(whyWord, "没有收到货物")
	whyWord = append(whyWord, "为啥取消寄件")
	whyWord = append(whyWord, "为啥取消拦截")
	whyWord = append(whyWord, "为啥拦截")
	whyWord = append(whyWord, "为啥退回")
	whyWord = append(whyWord, "签收未收到")
	whyWord = append(whyWord, "发货了吗")
	var whyWordAC *AcAutoMachine

	whyWordAC = NewAcAutoMachine()
	for _, ww := range whyWord {
		whyWordAC.AddPattern(ww)
	}
	whyWordAC.Build()
	msg := "不更新物流已退款找到退回找不到记丢件;"
	msg = "长时间不更新物流找到退回找不到记丢件;"
	whyWordContain := false
	acWhyWordContain := false
	for _, value := range whyWord {
		if strings.Contains(msg, value) {
			whyWordContain = true
			println("包含的是", value)
			break
		}
	}

	result := whyWordAC.Query(msg)
	if result != nil {
		println("最短匹配到了")
	}
	results := whyWordAC.QueryLast(msg)
	if results != nil {
		acWhyWordContain = true
	}

	if acWhyWordContain != whyWordContain {
		println("Ac自动机(", acWhyWordContain, ")和Contains(", whyWordContain, ")算法不一样 ", msg)
	}

}

func TestSimple(t *testing.T) {
	machine := NewAcAutoMachine()

	machine.AddPattern("更新")
	machine.AddPattern("不更新")
	machine.AddPattern("不更")

	machine.Build()

	msg := "不更新物流"
	results := machine.Query(msg)
	for _, result := range results {
		println("====>", result.Key)
	}
	last := machine.QueryLast(msg)

	for _, l := range last {
		println("====>", l.Key)
	}
}
