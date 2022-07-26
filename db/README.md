# Подключение к БД :

Для подключения к базе данных используйте команду
#### `В папке config:`

```bash
cp db_conf.example db_conf.go
```

Далее внутри файла используйте логин и пароль от бд.

***


## Установка `migrate` 

### `linux:` 
```bash
$ curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add -
$ echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list
$ apt-get update
$ apt-get install -y migrate
```

#### `alternative`
```bash
$ curl -L https://github.com/golang-migrate/migrate/releases/download/v4.6.2/migrate.linux-amd64.tar.gz | tar xvz
```


### `Windows:`

Using [scoop](https://scoop.sh/)

```bash
$ scoop install migrate
```

### `Mac:`
```bash
$ brew install golang-migrate
```

## `Migrations`

# Важно: 
##### Никодна не используйте `migrate down` без указания числа N.

### Создание миграций :
Из root каталога проекта :

```bash
$ migrate create -ext sql -dir db/migrations -seq migrate_name
```
Создает 2 файла `up/down` обязательно прописать откат.

```bash
$ migrate -source file://db/migrations -database postgres://postgres:1117111@localhost:5432/vectorquest up
```


### Откат миграции :

```bash
$ migrate -source file://db/migrations -database postgres://postgres:1117111@localhost:5432/vectorquest down 1
```

Откат последних N миграций :

```bash
$ migrate -source file://db/migrations -database postgres://postgres:1117111@localhost:5432/vectorquest down N
```
