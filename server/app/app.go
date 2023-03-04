//this is not the most efficient way of packaging Go applications
//we assume all API always depends on the same set of dependencies
//With this structure, we can easily load application dependencies inside handlers.
/* For the moment it’s only about a Syslog logger.
 Same way, we can add the DB connections, Redis connections, messaging systems like NATS and etc.
*/

package app

import (
	"github.com/DerrickKirimi/Bookshelf/util"
)

type App struct {
	logger *logger.Logger
}

func New(logger *logger.Logger) *App {
	return &App{logger : logger}
}

func (app *App) Logger() *logger.Logger {
	return app.logger
}