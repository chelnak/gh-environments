package delete

import (
	"fmt"

	"github.com/chelnak/gh-environments/internal/client"
	"github.com/erikgeiser/promptkit/confirmation"
)

type DeleteOptions struct {
	Name  string
	Force bool
}

type deleteCmd struct {
	client client.Client
}

type DeleteCmd interface {
	Delete(opts *DeleteOptions)
}

func (s deleteCmd) Delete(opts *DeleteOptions) {
	if _, err := s.client.GetEnvironment(opts.Name); err != nil {
		fmt.Println(err)
		return
	}

	if !opts.Force {
		promptText := fmt.Sprintf("You are about to delete %s. Are you sure that you want to continue?", opts.Name)
		confirm := confirmation.New(promptText, confirmation.No)
		ready, err := confirm.RunPrompt()
		if err != nil {
			fmt.Println(err)
			return
		}

		if !ready {
			return
		}
	}

	if err := s.client.DeleteEnvironment(opts.Name); err != nil {
		fmt.Println(err)
		return
	}
}

func NewDeleteCmd(client client.Client) DeleteCmd {
	return &deleteCmd{
		client: client,
	}
}
