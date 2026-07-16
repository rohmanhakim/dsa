package lc155

type MinStack struct {
	data    []int
	minVals []int
}

func Constructor() MinStack {
	return MinStack{
		data:    []int{},
		minVals: []int{},
	}
}

func (this *MinStack) Push(value int) {
	if len(this.minVals) == 0 || (len(this.minVals) > 0 && value <= this.minVals[len(this.minVals)-1]) {
		this.minVals = append(this.minVals, value)
	}
	this.data = append(this.data, value)
}

func (this *MinStack) Pop() {
	if len(this.data) > 0 {
		top := this.data[len(this.data)-1]
		if len(this.minVals) > 0 && top == this.minVals[len(this.minVals)-1] {
			this.minVals = this.minVals[:len(this.minVals)-1]
		}
		this.data = this.data[:len(this.data)-1]
	}
}

func (this *MinStack) Top() int {
	if len(this.data) > 0 {
		return this.data[len(this.data)-1]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.minVals) > 0 {
		return this.minVals[len(this.minVals)-1]
	}
	return 0
}
