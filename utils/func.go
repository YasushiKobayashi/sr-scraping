package utils

func UniqStringArray(list []string) []string {
	m := make(map[string]struct{})

	newList := make([]string, 0)

	for _, element := range list {
		// mapでは、第二引数にその値が入っているかどうかの真偽値が入っている
		if _, ok := m[element]; ok == false {
			m[element] = struct{}{}
			newList = append(newList, element)
		}
	}

	return newList
}
