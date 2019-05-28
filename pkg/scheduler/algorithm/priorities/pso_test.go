package priorities

import (
	"math"
	"testing"
)

func TestCalPSO(t *testing.T) {
	var experes float64 = 1.0
	res := CalPSO(100)
	if math.Abs(res-experes) > 0.03 {
		t.Fatalf("测试失败，希望: %v   实际: %v", experes, res)
	}
	t.Logf("PSO算法的k8s调度策略 测试成功, 希望: %v   实际: %v", experes, res)
}
