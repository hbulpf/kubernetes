package priorities

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const N int = 2 //粒子个数，为方便演示，只取两个，观察其运动方向
var (
	c1    float64    = 2.0
	c2    float64    = 2.0
	x     [N]float64 //粒子位置
	y     [N]float64 //适应度值
	v     [N]float64 //粒子运行速度
	pBest [N]float64
	gBest float64 = math.Inf(-1)
	VMax  float64 = 0.1
)

//初始化
func init() {
	fmt.Println("PSO start to init...")
	//随机初始化值
	// r := rand.New(rand.NewSource(time.Now().Unix() % 10))
	rand.Seed((time.Now().UnixNano() / 100) % 100)
	x[0] = rand.Float64()*2 - 1
	x[1] = rand.Float64()*2 - 1
	v[0] = rand.Float64()
	v[1] = rand.Float64()

	baseFitnessFunc()
	// fmt.Println("---初始化当前个体极值，并找到群体极值---")
	for i := 0; i < N; i++ {
		pBest[i] = y[i]
		if y[i] > gBest {
			gBest = y[i]
		}
	}
	fmt.Printf("start \n gBest:%f \n x[0]:%f \n x[1]:%f \n v[0]:%f \n v[1]:%f\n",
		gBest, x[0], x[1], v[0], v[1])
}

//适应度函数
func baseFitnessFunc() {
	for i := 0; i < N; i++ {
		y[i] = -1 * x[i] * (x[i] - 2)
		// fmt.Printf("x[%d]=>y[%d]:%f=>%f \t\n", i, i, x[i], y[i])
	}
}

//粒子群算法计算
func CalPSO(maxEpoch int) (res float64) {
	fmt.Printf("x \t v \t pBest0 \t  y \t pBest1 \n")
	for i := 0; i < maxEpoch; i++ {
		// fmt.Printf("\n\n======%d======", i+1)
		w := 0.5
		for j := 0; j < N; j++ {
			rand.Seed((time.Now().UnixNano() / 100) % 100)
			v[j] = w*v[j] + c1*rand.Float64()*(pBest[j]-x[j]) + c2*rand.Float64()*(gBest-x[j])
			if v[j] < VMax {
				v[j] = VMax
			}
			// fmt.Printf("x[%d](%f)+v[%d](%f)=>%f \n", j, x[j], j, x[j], x[j]+v[j])
			x[j] += v[j]
			//越界判断
			if x[j] > 2 {
				x[j] = 2
			}
			if x[j] < -2 {
				x[j] = 2
			}
			// fmt.Printf("越界检查后:\tx[%d]=%f \n", j, x[j])
		}
		baseFitnessFunc()
		//更新个体极值和群体极值
		// fmt.Printf("\n第  %d 次更新个体极值和群体极值\n", i+1)
		for j := 0; j < N; j++ {
			// tmp := pBest[j]
			pBest[j] = math.Max(y[j], pBest[j])
			// fmt.Printf("%f \t %f \t %f \t  %f \t %f\n", x[j], v[j], tmp, y[j], pBest[j])
			if pBest[j] > gBest {
				gBest = pBest[j]
			}
		}
		// fmt.Printf("======%d======gBest:%f\n", i+1, gBest)
	}
	res = gBest
	return
}
