package task

func Init() {
	initClient()
	go registerWorker()
}
