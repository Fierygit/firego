package todolist

func ReverseTodoList(todolist []TodoModel) {
	low, high := 0, len(todolist)-1
	for low < high {
		todolist[low], todolist[high] = todolist[high], todolist[low]
		low++
		high--
	}
}
