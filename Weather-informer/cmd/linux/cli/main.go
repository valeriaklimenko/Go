package main

import (
"fmt"
"os"

"github.com/valeriaklimenko/Weather-informer/internal/pkg/app/cli"
"github.com/valeriaklimenko/Weather-informer/pkg/logger"
)

func main() {
log := logger.NewSimpleLogger()
app := cli.New(log)
err := app.Run()
if err != nil {
log.Error("Приложение завершилось с ошибкой")
fmt.Printf("Some error - %s\n", err.Error())
os.Exit(1)
}
log.Info("Приложение успешно завершило работу")
os.Exit(0)
}
