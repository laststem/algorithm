/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type Codec struct {
    
}

func Constructor() Codec {
    return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
    if root == nil {
        return ""
    }
    return fmt.Sprintf("%d,%s,%s", root.Val, this.serialize(root.Left), this.serialize(root.Right))
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
    numbers := strings.Split(data, ",")
    queue := &Queue{data: numbers}
    for _, number := range numbers {
        queue.push(number)
    }
    deserializer := &De{}
    return deserializer.deserialize(queue)
}

type Queue struct {
    data []string
}

func (q *Queue) push(data string) {
    q.data = append(q.data, data)   
}

func (q *Queue) empty() bool {
    return len(q.data) == 0
}

func (q *Queue) pop() string {
    ret := q.data[0]
    q.data = q.data[1:]
    return ret
}

type De struct {}

func (d *De) deserialize(queue *Queue) *TreeNode {
    if queue.empty() {
        return nil
    }
    val := queue.pop()
    if val == "" {
        return nil
    }
    v, _ := strconv.Atoi(val)
    node := &TreeNode{Val: v}
    node.Left = d.deserialize(queue)    
    node.Right = d.deserialize(queue)
    return node
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * data := obj.serialize(root);
 * ans := obj.deserialize(data);
 */
