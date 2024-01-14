package snow

import "fmt"

func Snow() { //被calc包导入的一个包
	fmt.Println("笑死了...")
}

func init() {
	fmt.Println("snow.init()")
}
