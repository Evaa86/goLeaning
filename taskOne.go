/*
Сформировать данные для отправки заказа из
магазина по накладной и вывести на экран:
1) Наименование товара (минимум 1, максимум 100)
2) Количество (только числа)
3) ФИО покупателя (только буквы)
4) Контактный телефон (10 цифр)
5) Адрес(индекс(ровно 6 цифр), город, улица, дом, квартира)

Эти данные не могут быть пустыми.
Проверить правильность заполнения полей.

реализовать несколько методов у типа "Накладная"

createReader == NewReader
*/
package main

import (
	"fmt"
)

type Invoice struct {
	ProductName  string
	Quantity     int
	CustomerName string
	ContactPhone string
	Address      Address
}

type Address struct {
	Index     string
	City      string
	Street    string
	House     string
	Apartment string
}

func NewInvoice(productName string, quantity int, customerName string, contactPhone string, address Address) *Invoice {
	return &Invoice{ 
		productName,
		quantity,
		customerName,
		contactPhone,
		address,
	}
}

func (inv *Invoice) Validate() error {
	if inv.ProductName == "" {
		return fmt.Errorf("наименование товара не может быть пустым")
	}

	if inv.Quantity <= 0 {
		return fmt.Errorf("количество товара должно быть больше нуля")
	}

	if inv.CustomerName == "" {
		return fmt.Errorf("ФИО покупателя не может быть пустым")
	}

	if len(inv.ContactPhone) != 11 {
		return fmt.Errorf("контактный телефон должен содержать ровно 11 цифр") 
	}

	if inv.Address.Index == "" || len(inv.Address.Index) != 6 {
		return fmt.Errorf("некорректно заполнен индекс")
	}

	if inv.Address.City == "" {
		return fmt.Errorf("город не может быть пустым")
	}

	if inv.Address.Street == "" {
		return fmt.Errorf("улица не может быть пустой")
	}

	if inv.Address.House == "" {
		return fmt.Errorf("номер дома не может быть пустым")
	}

	if inv.Address.Apartment == "" {
		return fmt.Errorf("номер квартиры не может быть пустым")
	}

	return nil
}

func main() {
	address := Address{
		Index:     "123456",
		City:      "Москва",
		Street:    "Улица Пушкина",
		House:     "1",
		Apartment: "1",
	}

	invoice := NewInvoice("Кресло", 2, "Петров Иван Петрович", "81234567890", address)
	if err := invoice.Validate(); err != nil {
		fmt.Println("Ошибка при создании накладной:", err)
		return
	}

	fmt.Println("\nДанные заказа:")
	fmt.Printf("Наименование товара: %s\n", invoice.ProductName)
	fmt.Printf("Количество: %d\n", invoice.Quantity)
	fmt.Printf("ФИО покупателя: %s\n", invoice.CustomerName)
	fmt.Printf("Контактный телефон: %s\n", invoice.ContactPhone)
	fmt.Println("\nАдрес:")
	fmt.Printf("Индекс: %s\n", invoice.Address.Index)
	fmt.Printf("Город: %s\n", invoice.Address.City)
	fmt.Printf("Улица: %s\n", invoice.Address.Street)
	fmt.Printf("Дом: %s\n", invoice.Address.House)
	fmt.Printf("Квартира: %s\n", invoice.Address.Apartment)
}
