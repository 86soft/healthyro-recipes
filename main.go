package main

import (
	"fmt"
	"net/http"
)

func main() {
	/*ctx := context.Background()
	application := app.Application{
		Commands: app.Commands{
			CreateRecipe:             command.CreateRecipeHandler{},
			UpdateRecipeTitle:        command.UpdateRecipeTitleHandler{},
			UpdateRecipeDescription:  command.UpdateRecipeDescriptionHandler{},
			UpdateRecipeExternalLink: command.UpdateRecipeExternalLinkHandler{},
			DeleteRecipe:             command.DeleteRecipeHandler{},
		},
		Queries:  app.Queries{},
	}(ctx)
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":80", nil)*/
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Dzialaaaaaaaaaaa")
}
