package router
import(
	"github.com/go-chi/chi"
	"github.com/DerrickKirimi/Bookshelf/server/app"
)

func New() *chi.Mux {
	r := chi.NewRouter()
	r.MethodFunc("Get", "/", app.HandleIndex)
	return r
}