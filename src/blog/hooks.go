//add your own func before start server
package routers

//define our HookFunc base struct
type hookFunc func() error

var (
	hooks = make([]hookFunc, 0) //hook function slice to store the hookfunc
)

//main method to add your hookFunc to system
func RegisterHookFunc(hook hookFunc)  {
	hooks = append(hooks,hook)
}

func startBeforeHttpRun()  {
	//add hookFunc


	//start hookFunc
	for _,hk :=range hooks{
		if err:=hk();err !=nil{
			panic(err)
		}
	}
}