package main

import "fmt"

type treePoint struct { // BST
	number      int        // число в корневом узле
	leftPoint   *treePoint // узел со значением меньше текущего
	rightPoint  *treePoint // узел со значением больше текущего
	parentPoint *treePoint // родительский узел
}

func (root *treePoint) Add(Number int) { // добавление элемента

	switch {
	case Number == root.number:
		fmt.Println("Узел с таким номером уже существует")
	case Number > root.number:
		if root.rightPoint == nil {
			root.rightPoint = &treePoint{number: Number, leftPoint: nil, rightPoint: nil, parentPoint: root}
			return
		} else {
			root.rightPoint.Add(Number) // рекурсия
			return
		}
	case Number < root.number:
		if root.leftPoint == nil {
			root.leftPoint = &treePoint{number: Number, leftPoint: nil, rightPoint: nil, parentPoint: root}
			return
		} else {
			root.leftPoint.Add(Number)
			return
		}
	}
}

func (root *treePoint) Delete(Number int) { // удаление элемента

	exist, point := root.isExist(Number)

	if !exist { // проверка сществования узла перед удалением
		fmt.Println("Узла с таким номером не существует")
		return
	}

	if checkIsLast(point) { // если удаляемый узел не имеет дочерних
		if point.number < point.parentPoint.number {
			point.parentPoint.leftPoint = nil
		} else {
			point.parentPoint.rightPoint = nil
		}
		return
	}

	if checkIsOneDaughter(point) { // если у удаляемого узла только один дочерний
		if point.leftPoint != nil {
			if point.number < point.parentPoint.number {
				point.parentPoint.leftPoint = point.leftPoint
			} else {
				point.parentPoint.rightPoint = point.leftPoint
			}
		} else {
			if point.number < point.parentPoint.number {
				point.parentPoint.leftPoint = point.rightPoint
			} else {
				point.parentPoint.rightPoint = point.rightPoint
			}
		}
		return
	}

	minPoint := findMin(*point.rightPoint)
	minNumber := minPoint.number
	point.Delete(minNumber)
	point.number = minNumber
}

func (root *treePoint) isExist(Number int) (exist bool, point *treePoint) { // проверка наличия элемента и получение указателя на него

	switch {
	case Number == root.number:
		return true, root

	case Number > root.number:
		if root.rightPoint == nil {
			return false, nil
		}

		if Number == root.rightPoint.number {
			return true, root.rightPoint
		} else {
			return root.rightPoint.isExist(Number)
		}

	case Number < root.number:
		if root.leftPoint == nil {
			return false, nil
		}

		if Number == root.leftPoint.number {
			return true, root.leftPoint
		} else {
			return root.leftPoint.isExist(Number)
		}
	}

	return false, nil
}

func checkIsLast(point *treePoint) bool { // проверка, является ли узел последним (проверяем при удалении узла)
	if (point.leftPoint == nil) && (point.rightPoint == nil) {
		return true
	}
	return false
}

func checkIsOneDaughter(point *treePoint) bool { // проверка, что у узла есть всего один дочерний (проверяем при удалении узла)
	caseA := (point.leftPoint == nil) && (point.leftPoint != nil)
	caseB := (point.leftPoint != nil) && (point.rightPoint == nil)
	if caseA || caseB {
		return true
	}
	return false
}

func findMin(point treePoint) treePoint { // поиск минимального элемента в правом поддереве узла point
	if point.leftPoint == nil {
		return point
	}
	return findMin(*point.leftPoint)
}

func main() {
	var bst treePoint = treePoint{number: 9, leftPoint: nil, rightPoint: nil, parentPoint: nil}

	// Добавление элементов
	bst.Add(6)
	bst.Add(17)
	bst.Add(3)
	bst.Add(8)
	bst.Add(16)
	bst.Add(20)
	bst.Add(1)
	bst.Add(4)
	bst.Add(7)
	bst.Add(12)
	bst.Add(19)
	bst.Add(21)
	bst.Add(2)
	bst.Add(5)
	bst.Add(11)
	bst.Add(14)
	bst.Add(18)
	bst.Add(10)
	bst.Add(13)
	bst.Add(15)
	bst.Add(15) // Узел с таким номером уже существует

	// Проверка существования
	fmt.Println(bst.isExist(21)) // true &{21 <nil> <nil> 0x140000920e0}
	fmt.Println(bst.isExist(9))  // true &{9 0x14000092040 0x14000092060 <nil>}
	fmt.Println(bst.isExist(12)) // true &{12 0x14000092200 0x14000092220 0x140000920c0}
	fmt.Println(bst.isExist(52)) // false <nil>

	// Удаление элементов
	fmt.Println(bst.leftPoint.rightPoint.number) // 8
	bst.Delete(8)
	fmt.Println(bst.isExist(8))                  // false <nil>
	fmt.Println(bst.leftPoint.rightPoint.number) // 7

	fmt.Println(bst.rightPoint.rightPoint.number) // 20
	bst.Delete(20)
	fmt.Println(bst.isExist(20))                  // false <nil>
	fmt.Println(bst.rightPoint.rightPoint.number) // 21

	fmt.Println(bst.rightPoint.leftPoint.leftPoint.number) // 12
	bst.Delete(12)
	fmt.Println(bst.isExist(12))                           // false <nil>
	fmt.Println(bst.rightPoint.leftPoint.leftPoint.number) // 13
}
