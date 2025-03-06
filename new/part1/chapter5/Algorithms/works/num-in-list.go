package works

func NumInList(list []int, num int) bool {
    if len(list) == 0 {
        return false
    }

    for _, number := range list {
        if num == number {
            return true
        }
    }

    return false
}
