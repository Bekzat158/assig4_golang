package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"

	"Bekzat/pkg/store/postgres"
	"Bekzat/pkg/tracing"
	"Bekzat/pkg/type/context"
	log "Bekzat/pkg/type/logger"
	deliveryHttp "Bekzat/services/contact/internal/delivery/http"

	repositoryStorage "Bekzat/services/contact/internal/repository/storage/postgres"
	useCaseContact "Bekzat/services/contact/internal/useCase/contact"
	useCaseGroup "Bekzat/services/contact/internal/useCase/group"
)

func init() {
	viper.SetConfigName(".env")
	viper.SetConfigType("dotenv")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	viper.SetDefault("SERVICE_NAME", "contactService")
}

func main() {
	conn, err := postgres.New(postgres.Settings{})
	if err != nil {
		panic(err)
	}
	defer conn.Pool.Close()

	closer, err := tracing.New(context.Empty())
	if err != nil {
		panic(err)
	}
	defer func() {
		if err = closer.Close(); err != nil {
			log.Error(err)
		}
	}()

	

	repoStorage, err := repositoryStorage.New(conn.Pool, repositoryStorage.Options{})
	if err != nil {
		panic(err)
	}
	var (
		ucContact = useCaseContact.New(repoStorage, useCaseContact.Options{})
		// ucGroup      = useCaseGroup.New(repoGroup, useCaseGroup.Options{})
		ucGroup      = useCaseGroup.New(repoStorage, useCaseGroup.Options{})
		listenerHttp = deliveryHttp.New(ucContact, ucGroup, deliveryHttp.Options{})
	)

	go func() {
		fmt.Printf("service started successfully on http port: %d", viper.GetUint("HTTP_PORT"))
		if err = listenerHttp.Run(); err != nil {
			panic(err)
		}
	}()

	signalCh := make(chan os.Signal, 1)
	signal.Notify(signalCh, syscall.SIGINT, syscall.SIGTERM)
	<-signalCh

}
