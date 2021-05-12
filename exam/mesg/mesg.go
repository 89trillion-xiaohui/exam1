package mesg

type Soldier struct {
	Id           int `json:"id"`    //士兵id
	Rarity       int `json:"rarity"`   //士兵稀有度
	Unlockarena  int `json:"unlockarena"`   //解锁阶段
	Combatpoints int `json:"combatpoints"`  //战力
	Name         string `json:"name"` //名字
	Cvc 		 int `json:"cvc"` //cvc client version code
}

// PrintRarity 输入士兵id获取稀有度
func PrintRarity (id int) int  {
	var soldier Soldier
	if id == soldier.Id {
		return soldier.Rarity
	}
	return -1
}

// PrintCombatPoints 输入士兵id获取战力
func PrintCombatPoints (id int) int  {
	var soldier Soldier
	if id == soldier.Id {
		return soldier.Combatpoints
	}
	return -1
}

// PrintCvcSoldier 输入稀有度，当前解锁阶段和cvc（client version code），获取该稀有度cvc合法且已解锁的所有士兵
func PrintCvcSoldier(rarity int, unlockarena  int, cvc int)  string {
	var soldier Soldier
	if (rarity == soldier.Rarity)&&(unlockarena == 1)&&(cvc >= 10){
		return soldier.Name
	}
	return "Error"
}


func main() {

}
