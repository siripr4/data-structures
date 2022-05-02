package queue

type Queue []string

func (q *Queue) IsEmpty() bool{
	return len(*q) == 0
}

func (q *Queue) Push(element string){
	*q = append(*q, element)
}

func (q *Queue) Dequeue()(string, bool){
	if q.IsEmpty(){
		return "", false
	}
	element := (*q)[0]
	*q = (*q)[1:]
	return element, true
}