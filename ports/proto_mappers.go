package ports

import (
	"github.com/86soft/healthyro-recipes/domain"
	pb "github.com/86soft/healthyro/recipe"
)

func MapRecipesToProto(rs []domain.Recipe) (res []*pb.Recipe) {
	for _, r := range rs {
		res = append(res, MapRecipeToProto(&r))
	}
	return res
}

func MapRecipeToProto(r *domain.Recipe) *pb.Recipe {
	return &pb.Recipe{
		Uuid:        r.ID(),
		Title:       r.Title(),
		Description: r.Description(),
	}
}
