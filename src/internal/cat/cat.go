package cat

import "fmt"

type Cat struct {
	Name          string   `bson:"name,omitempty" json:"name"`
	Breed         string   `bson:"breed,omitempty" json:"breed"`
	Fur           string   `bson:"fur,omitempty" json:"fur"`
	Lives         int      `bson:"lives,omitempty" json:"lives"`
	Age           int      `bson:"age,omitempty" json:"age"`
	FavoriteFoods []string `bson:"favorite_foods" json:"favorite_foods"`
}

func CreateCat(age int, name string, breed string, lives int, fur string, favoriteFoods []string) Cat {
	return Cat{
		Age:           age,
		Name:          name,
		Breed:         breed,
		Lives:         lives,
		Fur:           fur,
		FavoriteFoods: favoriteFoods,
	}
}

func (c *Cat) PrintCat() string {
	return fmt.Sprintf("%s %s, %d, %s", c.Name, c.Breed, c.Age, c.Fur)
}
