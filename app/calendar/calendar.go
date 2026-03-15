package calendar

import (
	"errors"
	"fmt"

	"github.com/SamiRemi/project/app/events"
	"github.com/SamiRemi/project/app/validation"
	"github.com/araddon/dateparse"
)

var eventsMap = make(map[string]events.Event)

func AddEvent(title string, date string) (events.Event, error) {

	e, err := events.NewEvent(title, date)
	if err != nil {
		return e, err
	}
	if _, ok := eventsMap[title]; ok {
		return e, errors.New("Событие с именем " + title + " уже существует!")
	}
	if len(title) == 0 {
		return e, errors.New("Нельзя ввести пустое имя")
	}
	eventsMap[e.ID] = e
	fmt.Println("Событие добавлено:", e.Title)
	return e, nil
}

func ShowEvent() error {
	if len(eventsMap) == 0 {
		return errors.New("Список пуст")
	}
	for _, v := range eventsMap {
		utcTime := v.StartAt.UTC()
		fmt.Println(v.Title, "", utcTime.Format("02.01.2006 15:04"))
	}
	return nil
}

func DeleteEvent(title string) error {
	e := eventsMap[title]
	if _, ok := eventsMap[title]; !ok {
		return errors.New("Событие с именем " + title + " не существует")
	}
	delete(eventsMap, title)
	fmt.Println("=========================")
	fmt.Println("Событие :", e.Title)
	fmt.Println("С ID :", e.ID)
	fmt.Println("Удалено")
	fmt.Println("=========================")
	fmt.Println("")
	return nil
}

func EditEvent(name, newTitle, dateStr string) error {
	e := eventsMap[name]
	date, dateErr := dateparse.ParseAny(dateStr)
	if dateErr != nil {
		return dateErr
	}
	err := fullValidation(name, newTitle)
	if err != nil {
		return err
	}
	eventsMap[name] = events.Event{
		Title:   newTitle,
		StartAt: date,
	}
	fmt.Println("=========================")
	fmt.Println("Событие :", name)
	fmt.Println("С ID :", e.ID)
	fmt.Println("УСпешно изменено")
	fmt.Println("=========================")
	fmt.Println("")
	return nil
}

func fullValidation(name, title string) error {
	if name == "" {
		return errors.New("Имя события не может быть пустым")
	}
	if _, ok := eventsMap[name]; !ok {
		return errors.New("Событие не найдено ")
	}
	if ok := validation.IsValidTitle(title); !ok {
		return errors.New("Заголовок введён некорректно")
	}
	if eventsMap[name].Title == title {
		return errors.New("Новый заголовок совпадает с текущим")
	}
	return nil
}
