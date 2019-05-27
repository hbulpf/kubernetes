package priorities

import (
	"math"
	"testing"
)

func TestFsp(t *testing.T) {
	var experes float64 = 1.0
	res := CalFspPSO(100)
	if math.Abs(res-experes) > 0.03 {
		t.Fatalf("测试失败，希望: %v   实际: %v", experes, res)
	}
	t.Logf("基于PSO算法 的k8s调度算法 fsp 测试成功, 希望: %v   实际: %v", experes, res)
}
