// goprice
package goprice

import (
	"fmt"
	"time"
)

type Price struct {
	Contragent      string               //Код контрагента
	CreateAt        time.Time            //Время создания объекта
	Categories      map[string]*Category //Категории прайса
	categoriesStack map[int]*Category    //Стек текущих категорий
}

type Category struct {
	Name       string                //Наименование категории
	Url        string                //URL
	Categories map[string]*Category  //Подкатегории
	Items      map[string]*PriceItem //Позиции категории
}

type PriceItem struct {
	Code   string  //Кодовое обозначение позиции
	Name   string  //Наименование позиции
	Price  float64 //Цена позиции
	Valuta string  //Валюта цены
	Stock  string  //Количество у контрагента
}

/**
 * Возвращает инициализированные по умолчанию опции запроса
 */
func GetPrice(name string) *Price {
	price := new(Price)
	price.Contragent = name
	price.CreateAt = time.Now()
	price.Categories = make(map[string]*Category)
	price.categoriesStack = make(map[int]*Category)
	return price
}

func (p *Price) AddCategory(pCategory *Category) {
	p.Categories[pCategory.Name] = pCategory
}

func (p *Price) PushCategory(pCategory *Category) {
	var length = len(p.categoriesStack)
	fmt.Println("Длина стека: ", length)
	p.categoriesStack[length+1] = pCategory
}

func (p *Price) PullCategory() {
	var length = len(p.categoriesStack)
	fmt.Println("Длина стека: ", length)
	delete(p.categoriesStack, length)
}

func (p *Price) GetCurrentCategory() *Category {
	var length = len(p.categoriesStack)
	fmt.Println("Длина стека: ", length)
	return p.categoriesStack[length]

}

func CreateCategory(name string) *Category {
	category := new(Category)
	category.Name = name
	category.Categories = make(map[string]*Category)
	category.Items = make(map[string]*PriceItem)
	return category
}

func (cat *Category) SetUrl(url string) {
	cat.Url = url
}

func (cat *Category) AddCategory(pCategory *Category) {
	cat.Categories[pCategory.Name] = pCategory
}

func (cat *Category) AddItem(pItem *PriceItem) {
	cat.Items[pItem.Code] = pItem
}

func CreatePriceItem(code string, name string) *PriceItem {
	item := new(PriceItem)
	item.Name = name
	item.Code = code
	return item
}
