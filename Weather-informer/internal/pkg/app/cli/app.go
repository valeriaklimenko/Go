package cli

import (
"encoding/json"
"errors"
"fmt"
"io"
"net/http"

"github.com/valeriaklimenko/Weather-informer/pkg/logger"
)

type cliApp struct {
log logger.Logger
}

func New(log logger.Logger) *cliApp {
return &cliApp{log: log}
}

func (c *cliApp) Run() error {
c.log.Info("Запуск приложения Weather Informer")
c.log.Debug("Настройка параметров запроса к API")

type Current struct {
Temp float32 `json:"temperature_2m"`
}
type Response struct {
Curr Current `json:"current"`
}
var response Response

params := fmt.Sprintf(
"latitude=%f&longitude=%f&current=temperature_2m",
53.6688,
23.8223,
)
url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?%s", params)

c.log.Debug("Отправка запроса к open-meteo API")
resp, err := http.Get(url)
if err != nil {
c.log.Error("Ошибка при получении данных погоды")
customErr := errors.New("can't get weather data from openmeteo")
return errors.Join(customErr, err)
}
defer func() {
if err := resp.Body.Close(); err != nil {
c.log.Error("Ошибка при закрытии тела ответа")
fmt.Printf("can't close body err - %s\n", err.Error())
}
}()

c.log.Debug("Чтение данных ответа")
data, err := io.ReadAll(resp.Body)
if err != nil {
c.log.Error("Ошибка при чтении данных")
customErr := errors.New("can't read data from response")
return errors.Join(customErr, err)
}

c.log.Debug("Парсинг JSON ответа")
if err := json.Unmarshal(data, &response); err != nil {
c.log.Error("Ошибка при парсинге JSON")
customErr := errors.New("can't unmarshal data from response")
return errors.Join(customErr, err)
}

c.log.Info("Успешное получение данных о погоде")
fmt.Printf("Температура воздуха - %.2f градусов цельсия\n", response.Curr.Temp)
return nil
}
