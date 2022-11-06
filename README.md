#  Бот

## <a href="https://">Курс</a>

В качестве хранилища используется <a href="https://github.com/boltdb/bolt">Bolt DB</a>.

Чтобы реализовать авторизацию пользователей, вместе с ботом запускается HTTP сервер на порту 80, на который происходит редирект от Pocket при успешной авторизации пользователя. 

Когда сервер принимает запрос, он генерирует Access Token через Pocket API для пользователя и сохраняет его в хранилище.

### Стек:
- Go 1.19
- BoltDB
- Docker
