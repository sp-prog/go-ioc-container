# go-ioc-container
Контейнер инверсии управления для go

Инверсия управления - это соглащение о использовании внутри данного объекта только интерфейсов других объектов

Внедрение зависимостей (пассивная инверсия зависимостей) - это частный случай инверсии управления. Некое соглашение о создании зависимостей вне объекта. Созданные вне объекта зависимости могут передаваться как через конструктор, так и через методы или любым другим доступным способом.

Поиск зависимостей (активная инверсия зависимостей) - другой частный случай инверсии управления. Некое соглашение о создании зависимостей внутри объекта через специальный поставщик объектов (тот самый контейнер IoC). Поставщик передается в объект любым доступным способом

Инверсия зависимостей - это буква "D" в слове SOLID

Все это нужно для:
- упрощения тестирования. А в некоторых случаях и просто для появления возможности написать тесты
- ослабления связанности компонентов

Подробнее можно посмотреть тут:
- https://ru.stackoverflow.com/questions/713050/%D0%A0%D0%B0%D0%B7%D0%BD%D0%B8%D1%86%D0%B0-%D0%BC%D0%B5%D0%B6%D0%B4%D1%83-%D0%B8%D0%BD%D0%B2%D0%B5%D1%80%D1%81%D0%B8%D0%B5%D0%B9-%D1%83%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D0%B5%D0%BD%D0%B8%D1%8F-%D0%B8-%D0%B2%D0%BD%D0%B5%D0%B4%D1%80%D0%B5%D0%BD%D0%B8%D0%B5%D0%BC-%D0%B7%D0%B0%D0%B2%D0%B8%D1%81%D0%B8%D0%BC%D0%BE%D1%81%D1%82%D0%B5%D0%B9
- Spring 4 для профессионалов Криса Шеффера и компании

Проект создан не "от хорошей жизни", а пришлось.
Известные на момент написания этого куска текста реализации Go DI IoC не поддерживали возможность создания
объектов со временем жизни "Transient", т.е. время жизни на период существования блока кода,
куда данная зависимость была передана механизмом IoC.

Это реализации, которые "не прошли":
- google wire https://pkg.go.dev/github.com/google/wire - это вообще требует вызова генератора кода
- uber dig и fx https://github.com/uber-go/dig , https://github.com/uber-go/fx
- sarulabs DI https://github.com/sarulabs/di

Исходный код проект основан на примерах из книги "Pro Go Полное руководство ..." Адама Фримана, т.е. на этом: https://github.com/Apress/pro-go/tree/main/32%20-%20Platform%20-%20Part%201

# Установка

Предполагалось выпожить всю эту прелесть на GitFlic, но там - па-ба-ба-бам! - для создания публичного репозитория нужен бубен от яндекса и еще какие манипуляции с нижними полушариями мозга.

Поэтому, для установки необходимо сделать это:

import "github.com/sp-prog/go-ioc-container"

# Подключение пакета к проекту

Тут будет описано как подключить сию писанину к бекенд-приложению. Скорее-всего, только к fiber'у.
Почему именно к нему? Можно частично понять, просмотрев этот ресурс: https://github.com/mingrammer/go-web-framework-stars?tab=readme-ov-file
Почему "скорее-всего"? На некоторых ресурсах пишут, что fiber при загрузке с/на сервер(а) больших файлов может немножечко упасть из-за нехватким памяти. Т.е. не умеет работать с потоками. Если опасения подтвердятся, то придется что-то делать...

