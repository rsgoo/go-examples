package unit

import "testing"

func TestMyGetSum(t *testing.T) {
	t.Log("start to test func MyGetSum....")
	sum := MyGetSum(10)
	if sum != 55 {
		t.Errorf("期望结果是%d, 实际结果是%d\n", 55, sum)
		t.FailNow()
	}
	t.Log("finish to test func MyGetSum....")
}

func TestMyGetRecursiveSum(t *testing.T) {
	t.Log("start to test func MyGetSum....")
	sum := MyGetRecursiveSum(100)
	if sum != MyGetSum(10) {
		t.Fatalf("期望结果是%d, 实际结果是%d\n", MyGetSum(100), sum)
	}
	t.Log("finish to test func MyGetSum....")
}
