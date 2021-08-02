/*
Score:
执行用时：0 ms, 在所有 Go 提交中击败了100.00% 的用户
内存消耗：2.2 MB, 在所有 Go 提交中击败了100.00% 的用户

Points:
1.

Thoughts:
1. 这涉及到线性代数， 顺时针旋转90的矩阵是[[0, 1], [-1, 0]], 前提是以原点为中心旋转
2. 为了减少辅助变量， 只用一个tmp应当每次把一个圈上的数旋转
3. 真正用起来会涉及到浮点运算， 所以将这个过程通俗化， 原点平移至中心， 旋转， 原点平移回左上角
4. 并且， 值得注意的是， y轴的方向与数学中的笛卡尔坐标方向相反, 这意味者视觉上来看是顺时针， 在笛卡尔坐标中是逆时针操作的， 所以旋转矩阵是[[0,-1],[1,0]]
5. 并且x,y 是按[y][x]来操作的
6. 还需要注意， 长度为2的矩阵， length其实为1，因为这个矩阵其实只有4格子， 一个格子对应一个点
7. 总的说来说，就是matrix[y][x]--> matrix[length-x][y]
8. 并不能按一个顺序， 接连交换值来实现目的， 需要四个点之一的点与顺时针方向的值作交换， 然后还是这个点与逆时针方向的值作交换， 用刚被交换的点接着进行逆时针交换


给定一个 n × n 的二维矩阵 matrix 表示一个图像。请你将图像顺时针旋转 90 度。

你必须在 原地 旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要 使用另一个矩阵来旋转图像。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-image
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

*/
package leetcode

import "fmt"

func rotate(matrix [][]int) {
	tmp_x, tmp_y, length, x, y := 0, 0, len(matrix), 0, 0
	for i := 0; i < length/2; i++ {
		fmt.Println("i: ", i)
		//遍历
		// x = length - i - 1

		for j := length - i - 1; j > i; j-- {
			fmt.Println("\tj: ", j)
			x = length - i - 1
			y = j
			fmt.Println(x, y)
			//一圈四个值只需要交换三次
			//还不能顺时针换

			tmp_x, tmp_y = length-y-1, x //视觉上顺时针交换值
			// tmp_x = tmp_y * tmp_x // for test

			fmt.Printf("(%d, %d) --> (%d, %d)", x, y, tmp_x, tmp_y)

			matrix[tmp_y][tmp_x], matrix[y][x] = matrix[y][x], matrix[tmp_y][tmp_x]
			fmt.Println(matrix)
			tmp_x, tmp_y = y, length-x-1 //视觉上逆时针交换值
			fmt.Printf("(%d, %d) --> (%d, %d)", x, y, tmp_x, tmp_y)

			matrix[tmp_y][tmp_x], matrix[y][x] = matrix[y][x], matrix[tmp_y][tmp_x]
			fmt.Println(matrix)
			x, y = tmp_x, tmp_y
			tmp_x, tmp_y = y, length-x-1 //视觉上接着逆时针
			fmt.Printf("(%d, %d) --> (%d, %d)", x, y, tmp_x, tmp_y)

			matrix[tmp_y][tmp_x], matrix[y][x] = matrix[y][x], matrix[tmp_y][tmp_x]
			fmt.Println(matrix)

		}

	}
	fmt.Println(matrix)

}
