package main

import (
    "github.com/go-vgo/robotgo"
)

func main() {
    // 移动鼠标到指定坐标
    robotgo.MoveMouse(100, 200)
    
    // 鼠标移动并点击
    robotgo.Click()
    
    // 平滑移动
    robotgo.MoveMouseSmooth(300, 400)
}