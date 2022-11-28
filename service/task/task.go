package task

func Task() {
	tk1 := task.NewTask("tk1", "0 12 * * * *", func(ctx context.Context) error { fmt.Println("tk1"); return nil })
}