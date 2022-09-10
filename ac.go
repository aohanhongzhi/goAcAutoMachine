package goAcAutoMachine

type AcNode struct {
	fail      *AcNode
	isPattern bool
	node      map[rune]*AcNode
}

func newAcNode() *AcNode {
	return &AcNode{
		fail:      nil,
		isPattern: false,
		node:      map[rune]*AcNode{},
	}
}

type AcAutoMachine struct {
	root *AcNode
}

type Result struct {
	Key   string
	Start int
	End   int
}

func NewAcAutoMachine() *AcAutoMachine {
	return &AcAutoMachine{
		root: newAcNode(),
	}
}

func (ac *AcAutoMachine) AddPattern(pattern string) {
	chars := []rune(pattern)
	iter := ac.root
	for _, c := range chars {
		if _, ok := iter.node[c]; !ok {
			iter.node[c] = newAcNode()
		}
		iter = iter.node[c]
	}
	iter.isPattern = true
}

func (ac *AcAutoMachine) Build() {
	queue := []*AcNode{}
	queue = append(queue, ac.root)
	for len(queue) != 0 {
		parent := queue[0]
		queue = queue[1:]

		for char, child := range parent.node {
			if parent == ac.root {
				child.fail = ac.root
			} else {
				failAcNode := parent.fail
				for failAcNode != nil {
					if _, ok := failAcNode.node[char]; ok {
						child.fail = failAcNode.node[char]
						break
					}
					failAcNode = failAcNode.fail
				}
				if failAcNode == nil {
					child.fail = ac.root
				}
			}
			queue = append(queue, child)
		}
	}
}

func (ac *AcAutoMachine) Query(content string) (results []Result) {
	chars := []rune(content)
	iter := ac.root
	var start, end int
	for i, value := range chars {
		c := value // 方便debug，赋值每次都跳转到这里
		_, ok := iter.node[c]
		for !ok && iter != ac.root {
			iter = iter.fail
		}
		if _, ok = iter.node[c]; ok {
			if iter == ac.root { // this is the first match, record the start position
				start = i
			}
			iter = iter.node[c]
			if iter.isPattern {
				end = i // this is the end match, record one result
				result := Result{
					Key:   string([]rune(content)[start : end+1]),
					Start: start,
					End:   end + 1,
				}
				// results = append(results, string([]rune(content)[start:end+1]))
				results = append(results, result)
			}
		}
	}
	return
}

// 仅匹配最长的关键字
func (ac *AcAutoMachine) QueryLast(content string) (results []Result) {
	chars := []rune(content)
	iter := ac.root
	var start, lastStart, end int
	var lastResult Result // cant remove!!!
	var result Result
	for i, c := range chars {
		_, ok := iter.node[c]
		for !ok && iter != ac.root {
			iter = iter.fail
		}
		if _, ok = iter.node[c]; ok {
			if iter == ac.root { // this is the first match, record the start position
				start = i
				result = Result{}
			}
			iter = iter.node[c]
			if iter.isPattern {
				end = i // this is the end match, record one result
				result = Result{
					Key:   string([]rune(content)[start : end+1]),
					Start: start,
					End:   end + 1,
				}

				// lastStart变了
				if lastStart != start {

					if (lastResult != Result{} && lastResult.Start != start) {
						results = append(results, lastResult)
						lastResult = Result{}
					}
				}

				lastResult = result
				lastStart = start

			}
		}

		// 结束
		if i == len(chars)-1 {
			if (lastResult != Result{}) {
				results = append(results, lastResult)
			}
		}

	}
	return
}
