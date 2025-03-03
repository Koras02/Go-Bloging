package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

func (c *Character) AttackEnemy(enemy *Character) {
	damage := c.Attack - enemy.Defense
	if damage < 0 {
		damage = 0
	}
	enemy.Health -= damage
	fmt.Printf("%s가 %s에게 %d의 피해를 주었습니다!\n", c.Name, enemy.Name, damage)
}

func (c *Character) Defend() {
	c.Defense += 5
	fmt.Printf("%s가 방어 자세를 취했습니다! 방어력이 %d로 증가했습니다.\n", c.Name, c.Defense)
}

func (c *Character) Escape() bool {
	success := rand.Intn(2) // 50% 확률로 도망망
	if success == 0 {
		fmt.Printf("%s가 도망쳤습니다!\n", c.Name)
		return true
	} 
	fmt.Printf("%s가 도망치는 데 실패했습니다!\n", c.Name)
	return false
}

func (c *Character) isAlive() bool {
	return c.Health > 0
}

func (c *Character) ShowHealth() {
	fmt.Printf("%s의 현재 체력: %d\n", c.Name, c.Health)
}

func main() {
     rand.Seed(time.Now().UnixNano()) // 랜덤 시드 초기화


	link := Character{Name: "Link", Health: 100, Attack: 20, Defense: 5}
	ganondorf := Character{Name: "Ganondorf", Health: 120, Attack: 25, Defense: 10}

	for link.isAlive() && ganondorf.isAlive() {
		fmt.Println("\n 선택")
		fmt.Println("1: 공격")
		fmt.Println("2: 방어")
		fmt.Println("3: 도망가기")

		var choice int 
		fmt.Scan(&choice)

		switch choice {
		case 1:
			link.AttackEnemy(&ganondorf)
		case 2:
			link.Defend()
		case 3:
			if link.Escape() {
				fmt.Println("전투에서 도망쳤습니다.")
				return
			}
		default:
			fmt.Println("잘못된 선택입니다. 다시 선택하세요.")
			continue
		}
		// 공격후 현재 체력 출력 
		link.ShowHealth()
		ganondorf.ShowHealth()


		if !ganondorf.isAlive() {
			fmt.Printf("%s가 쓰러졌습니다!\n",ganondorf.Name)
			break
		}
		ganondorf.AttackEnemy(&link)
		if !link.isAlive() {
			fmt.Printf("%s가 쓰러졌습니다.", link.Name)
			break
		}

	}
}