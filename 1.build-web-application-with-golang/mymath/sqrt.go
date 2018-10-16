package mymath

//这样我的应用包目录和代码已经新建完毕，注意：一般建议package的名称和目录名保持一致
func Sqrt(x float64) float64 {
	z := 0.0
	for i := 0; i < 1000; i++ {
		z -= (z*z - x) / (2 * x)
	}
	return z
}

