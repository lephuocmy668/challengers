package kar

// https://leetcode.com/problems/keys-and-rooms
// Runtime: 4 ms, faster than 98.78% of Go online submissions for Keys and Rooms.
// Memory Usage: 4.2 MB, less than 50.00% of Go online submissions for Keys and Rooms.
func canVisitAllRooms(rooms [][]int) bool {
	// cache is used to check room is already visited or not
	cache := make(map[int]bool)
	// roomCount represents number of room need to visited
	roomCount := len(rooms)
	// stack is used to implement DFS
	stack := []int{0}

	for len(stack) > 0 {
		// pop room from stack to visit
		n := len(stack) - 1
		currRoom := stack[n]
		stack = stack[:n]

		// if current room isn't already visited, do dfs
		if _, ok := cache[currRoom]; !ok {
			roomCount--
			cache[currRoom] = true
			for _, room := range rooms[currRoom] {
				stack = append(stack, room)
			}
		}
	}

	return roomCount == 0
}
