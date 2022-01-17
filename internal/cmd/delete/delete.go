package delete

import (
	"fmt"
	"log"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/erikgeiser/promptkit/confirmation"
)

type DeleteOptions struct {
	Name string
}

type deleteService struct {
	client client.Client
}

type DeleteService interface {
	Delete(opts *DeleteOptions)
}

func (s deleteService) Delete(opts *DeleteOptions) {
	promptText := fmt.Sprintf("You are about to delete %s. Are you sure that you want to continue?", opts.Name)
	confirm := confirmation.New(promptText, confirmation.No)
	ready, err := confirm.RunPrompt()
	if err != nil {
		log.Fatal(err)
	}

	if ready {
		err = s.client.DeleteEnvironment(opts.Name)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func NewDeleteService(client client.Client) DeleteService {
	return &deleteService{
		client: client,
	}
}
