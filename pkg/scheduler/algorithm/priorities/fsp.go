package priorities

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const N int = 2 //粒子个数，为方便演示，只取两个，观察其运动方向
var (
	clusterSize int        = 64                             //集群规模
	C1max       float64    = 2.5                            //个体学习因子最小值
	C1min       float64    = 0.01                           //个体学习因子最大值
	C2max       float64    = 2.5                            //群体学习因子最大值
	C2min       float64    = 0.01                           //群体学习因子最小值
	X           [N]float64                                  //粒子位置
	Y           [N]float64                                  //适应度值
	V           [N]float64                                  //粒子运行速度
	u           float64    = 0.38                           //是惩罚系数， 控制粒子群聚集的程度
	R           float64    = .0                             //粒子运行速度惩罚项
	penalty     float64    = math.Log(float64(clusterSize)) //是否惩罚的阈值
	PBest       [N]float64
	GBest       float64 = math.Inf(-1)
	VMax        float64 = 0.1
)

//初始化
func init() {
	fmt.Println("FSP start to init...")

	//随机初始化值
	rand.Seed((time.Now().UnixNano() / 100) % 100)
	X[0] = rand.Float64()*2 - 1
	X[1] = rand.Float64()*2 - 1
	V[0] = rand.Float64()
	V[1] = rand.Float64()

	fspFitnessFunc()
	// fmt.Println("---初始化当前个体极值，并找到群体极值---")
	for i := 0; i < N; i++ {
		PBest[i] = Y[i]
		if Y[i] > GBest {
			GBest = Y[i]
		}
	}
	fmt.Printf("start \n GBest:%f \n X[0]:%f \n X[1]:%f \n V[0]:%f \n V[1]:%f\n",
		GBest, X[0], X[1], V[0], V[1])
}

//适应度函数
func fspFitnessFunc() {
	for i := 0; i < N; i++ {
		Y[i] = -1 * X[i] * (X[i] - 2)
		// fmt.Printf("x[%d]=>y[%d]:%f=>%f \t\n", i, i, X[i], Y[i])
	}
}

//计算方差
func getSPow2(Num [N]float64) float64 {
	if len(Num) < 1 {
		return .0
	}
	length := len(Num)
	sum := .0 //计算数组和
	for _, ele := range Num {
		sum += ele
	}
	ave := sum / float64(length)
	sum = 0 //sum 清零计算平方和
	for _, ele := range Num {
		sum += (ele - ave) * (ele - ave)
	}
	return sum / float64(length) //得到方差
}

//粒子群计算
func CalFspPSO(maxEpoch int) (res float64) {
	tmaxpow := float64(maxEpoch * maxEpoch)
	for i := 1; i <= maxEpoch; i++ {
		// fmt.Printf("\n\n======%d======", i)
		w := 0.5
		for j := 0; j < N; j++ {
			C1 := C1max - (C1max-C1min)*float64(i*i)/tmaxpow //计算C1
			C2 := C2min + (C2max-C2min)*float64(i*i)/tmaxpow //计算C1
			R = 0
			spow2 := getSPow2(X)
			if spow2 > penalty {
				R = (spow2 - penalty*penalty) * u / penalty //计算惩罚项
			}
			rand.Seed((time.Now().UnixNano() / 100) % 100)
			V[j] = w*V[j] + C1*rand.Float64()*(PBest[j]-X[j]) + C2*rand.Float64()*(GBest-X[j]) + R
			if V[j] < VMax {
				V[j] = VMax
			}
			// fmt.Printf("x[%d](%f)+v[%d](%f)=>%f \n", j, X[j], j, X[j], X[j]+V[j])
			X[j] += V[j]
			//越界判断
			if X[j] > 2 {
				X[j] = 2
			}
			if X[j] < -2 {
				X[j] = 2
			}
			// fmt.Printf("越界检查后:\tx[%d]=%f \n", j, X[j])
		}
		fspFitnessFunc()
		//更新个体极值和群体极值
		// fmt.Printf("\n第  %d 次更新个体极值和群体极值\n", i)
		for j := 0; j < N; j++ {
			// tmp := PBest[j]
			PBest[j] = math.Max(Y[j], PBest[j])
			// fmt.Printf("%f \t %f \t %f \t  %f \t %f\n", V[j], V[j], tmp, Y[j], PBest[j])
			if PBest[j] > GBest {
				GBest = PBest[j]
			}
		}
		// fmt.Printf("======%d======gbest:%f\n", i, GBest)
	}
	res = GBest
	return
}
