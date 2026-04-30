
type MinStack struct {
	min int;
	data []int;
	mins []int;
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min : int(^uint(0) >> 1),
		mins: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if val<=this.min {
		this.min=val
		this.mins=append(this.mins,val)
	}


}

func (this *MinStack) Pop() {
	val := this.Top();
	if val==this.min {
		if len(this.mins)-2>=0 {
			this.min=this.mins[len(this.mins)-2]
		}
		if len(this.mins)==1 {
			this.min=int(^uint(0) >> 1)
		}
		if len(this.mins)-1>=0 {
			this.mins=this.mins[:len(this.mins)-1]
		}
	}

	this.data = this.data[:len(this.data)-1]

	if len(this.mins)==0 && len(this.data)>0 {
		this.mins=append(this.mins,this.data[0])
		this.min=this.data[0]
	}
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1];
}

func (this *MinStack) GetMin() int {
	return this.min
}
