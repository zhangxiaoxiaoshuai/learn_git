//机器人在一个无限大小的 XY 网格平面上行走，从点 (0, 0) 处开始出发，面向北方。该机器人可以接收以下三种类型的命令 commands ：
//
//-2 ：向左转 90 度
//-1 ：向右转 90 度
//1 <= x <= 9 ：向前移动 x 个单位长度
//在网格上有一些格子被视为障碍物 obstacles 。第 i 个障碍物位于网格点  obstacles[i] = (xi, yi) 。
//
//机器人无法走到障碍物上，它将会停留在障碍物的前一个网格方块上，但仍然可以继续尝试进行该路线的其余部分。
//
//返回从原点到机器人所有经过的路径点（坐标为整数）的最大欧式距离的平方。（即，如果距离为 5 ，则返回 25 ）

func robotSim(commands []int, obstacles [][]int) int {
	// 方向：上右下左0123
	result, curDir, mObstacles := float64(0), 0, make(map[string]bool) // 结果 当前方向 障碍物哈希
	x, y, stepX, stepY := 0, 0, []int{0, 1, 0, -1}, []int{1, 0, -1, 0} // 当前的位置 以及 xy轴上各个方向移动的大小
	for _, v := range obstacles {
		mObstacles[fmt.Sprintf("%d-%d", v[0], v[1])] = true
	}
	for _, v := range commands {
		switch v {
		case -1:
			curDir = (curDir + 1) % 4
		case -2:
			curDir = (curDir + 3) % 4
		default:
			for i := 0; i < v; i++ {
				tempX, tempY := x+stepX[curDir], y+stepY[curDir]
				// 碰到障碍物，就不要移动了
				if mObstacles[fmt.Sprintf("%d-%d", tempX, tempY)] {
					break
				}
				x, y = tempX, tempY
				result = math.Max(float64(x*x+y*y), result)
			}
		}
	}
	return int(result)
}
