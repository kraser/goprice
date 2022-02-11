// goprice
package goprice

import (
	"fmt"
	"time"
)

type Price struct {
	Contragent string              //Код контрагента
	CreateAt   time                //Время создания объекта
	Categories map[string]Category //Категории прайса
}

type Category struct {
	Name       string              //Наименование категории
	Url        string              //URL
	Categories map[string]Category //Подкатегории
}

/**
 * Возвращает инициализированные по умолчанию опции запроса
 */
func GetPrice() *Price {
	price := new(Price)
	price.Contragent = "test"
	price.CreateAt = time.Now()
	price.Categories = make(map[string]Category)
	return price
}
