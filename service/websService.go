/*PACKAGE*/
package WebService

/**/

/*IMPORTS FROM PACKAGE*/
import (
	"context"
	WebRespository "seachProfile/domain"
	"seachProfile/models"
)

/**/

func GetUsersWebsService(ctx context.Context) ([]models.Webs, error) {
	return WebRespository.GetUsersWebsPopularRepository(ctx)
}

func CreateUsersService(ctx context.Context, web models.Webs) (int, error) {
	return WebRespository.CreateWebRepository(ctx, web)
}
