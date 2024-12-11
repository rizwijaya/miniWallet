package main

import (
	"errors"
	golog "log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	configLib "github.com/rizwijaya/miniWallet/infrastructures/config"
	log "github.com/rizwijaya/miniWallet/infrastructures/logger"
	middleware "github.com/rizwijaya/miniWallet/infrastructures/middlewares"

	database "github.com/rizwijaya/miniWallet/infrastructures/databases"
	mcache "github.com/rizwijaya/miniWallet/infrastructures/memcache"
	routerWalletAPIV1 "github.com/rizwijaya/miniWallet/modules/v1/wallet/routes"

	walletCtrl "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/controllers"
	walletRepo "github.com/rizwijaya/miniWallet/modules/v1/wallet/interfaces/repositories"
	walletUc "github.com/rizwijaya/miniWallet/modules/v1/wallet/usecases"
)

var (
	//controller
	walletController *walletCtrl.WalletController
)

func newApps(config configLib.LoadConfig) configLib.Routing {
	log.NewLogger(config)
	defer log.Sync()

	newRoute := configLib.Routing{}
	newRoute.Database = database.NewDatabase(config)
	newRoute.Memcache = mcache.NewMemcache(config)

	newRoute.Router = fiber.New(fiber.Config{
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError

			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}

			return ctx.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	newRoute.Router.Use(cors.New())

	return newRoute
}

func initEntity(apps configLib.Routing) {
	//Init Repository
	walletRepository := walletRepo.NewRepository(apps.Database, apps.Memcache)

	//Init Usecase
	walletUsecase := walletUc.NewUsecase(walletRepository)

	//Init Controller
	walletController = walletCtrl.NewController(walletUsecase)

	//Init Middleware
	middleware.NewMiddleware(apps.Database)
}

func main() {
	//Initialize config
	config, err := configLib.New()
	if err != nil {
		golog.Fatal(err)
	}

	//Create New Apps
	apps := newApps(config)
	api := apps.Router.Group("/api/v1")
	api.Get("/healthy", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(map[string]interface{}{
			"status":  200,
			"message": "Server running successfully",
		})
	})

	//Create New Entity
	initEntity(apps)

	//router apps
	routerWalletAPIV1.Router(walletController, api)

	//listen server
	err = apps.Router.Listen(config.App.Url + ":" + config.App.Port)
	if err != nil {
		log.Fatal(err)
	}
}
