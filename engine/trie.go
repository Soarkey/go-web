package engine

import "strings"

type node struct {
	pattern string  // 待匹配路由, 如 /p/:lang
	part    string  // 路由中的一部分, 例如 :lang
	childs  []*node // 子结点, 例如 [doc, tutorial, intro]
	isWild  bool    // 是否精确匹配, part 含有 : 或 * 时为 true
}

// matchChild 找到第一个匹配成功的结点, 用于插入
func (n *node) matchChild(part string) *node {
	for _, child := range n.childs {
		if child.part == part || child.isWild {
			return child
		}
	}
	return nil
}

// matchChilds 找到所有匹配成功的结点
func (n *node) matchChilds(part string) []*node {
	nodes := make([]*node, 0)
	for _, child := range n.childs {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}
	return nodes
}

// insert 插入 pattern parts height 字典树深度
func (n *node) insert(pattern string, parts []string, height int) {
	if len(parts) == height {
		n.pattern = pattern
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.childs = append(n.childs, child)
	}
	// 递归插入
	child.insert(pattern, parts, height+1)
}

// search 匹配
func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}
	part := parts[height]
	childs := n.matchChilds(part)
	for _, child := range childs {
		result := child.search(parts, height+1)
		if result != nil {
			return result
		}
	}
	return nil
}
