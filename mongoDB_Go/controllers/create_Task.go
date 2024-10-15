package controllers

func CreateTask(task *Task) error {
	_, err := Collection.InsertOne(Ctx, task)
	return err
}
