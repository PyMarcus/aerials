package models

type Timer struct{
	CurrentTicks int 
	TargetTicks int 
}

func (t *Timer) Update(){
	if t.CurrentTicks < t.TargetTicks{
		t.CurrentTicks++
	}
}

func (t *Timer) Ready() bool{
	return t.CurrentTicks >= t.TargetTicks
}

func (t *Timer) Reset(){
	t.CurrentTicks = 0
}