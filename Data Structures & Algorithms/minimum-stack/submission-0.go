type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	minValue := this.data[0]

	for i := 1; i < len(this.data); i++ {
		minValue = min(minValue, this.data[i])
	}
	return minValue
}
