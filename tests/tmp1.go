package main

import "fmt"

func main() {
	res := IsLocalForum("东城")
	fmt.Println(res)
}

func IsLocalForum(needleName string) map[string]interface{} {
	initReturn := map[string]interface{}{
		"is_local":   0,
		"local_name": "",
	}
	localForumArr := map[string][]string{
		"北京市": []string{"东城"},
	}
	for _, arr := range localForumArr {
		for _, tmpName := range arr {
			if tmpName == needleName {
				initReturn["is_local"] = 1
				initReturn["local_name"] = tmpName
				return initReturn
			}
		}
	}

	return initReturn
}
