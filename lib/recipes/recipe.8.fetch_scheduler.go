package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/scheduler"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	blocks "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/blocks/blocks.1"
	keyPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/key_pods/key_pods.1"

	log "github.com/sirupsen/logrus"
)

const (
	SchedulerKeyName   = "Tuition Class"
	SchedulerBlockName = "Final Exam"
	SchedulerPodName   = "Math Quiz"
	DueDate            = "2023-03-10T14:19:04.027Z"
)

func FetchScheduler() {
	log.Info("Objective: Set Due Date for Block and Pod, and Fetch Scheduler Events")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var key response.Key
	key, _ = recipes.AddCustomKey(user, SchedulerKeyName)
	log.Info("Set due date for block")
	err = setBlockDueDate(user, key)
	if err != nil {
		return
	}
	log.Info(".Block due date set successfully")

	log.Info("Set due date for pod")
	err = setPodDueDate(user, key)
	if err != nil {
		return
	}
	log.Info(".Pod due date set successfully")

	log.Info("Show due date events")
	err = fetchSchedulerEvents(user)
	if err != nil {
		return
	}
	log.Info(".Displayed block & pod due date events")
}

func setBlockDueDate(user response.User, key response.Key) error {
	block, err := recipes.AddBlock(user, SchedulerBlockName, key)
	if err != nil {
		return err
	}
	dueDate := DueDate
	_, err = blocks.UpdateBlock(
		user.JwtToken,
		blocks.UpdateBlockReqBody{DueDate: &dueDate},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID})
	if err != nil {
		return err
	}
	return nil
}

func setPodDueDate(user response.User, key response.Key) error {
	pod, err := recipes.AddPod(user, SchedulerPodName, key)
	if err != nil {
		return err
	}
	dueDate := DueDate
	_, err = keyPods.UpdateKeyPod(
		user.JwtToken,
		request.UpdatePodReqBody{DueDate: &dueDate},
		common.ResourceIdParam{PodId: pod.ID, KeyId: pod.Key.ID})
	if err != nil {
		return err
	}
	return nil
}

func fetchSchedulerEvents(user response.User) error {
	allEvents, err := scheduler.GetEventsForGivenDay(user.JwtToken, request.EventDateParam{StartDate: DueDate})
	if err != nil {
		return err
	}
	for _, blockEvent := range allEvents.DueDateEvent.Blocks {
		log.Printf(".Block %s is due on %s", blockEvent.Name, *blockEvent.DueDate)
	}
	for _, podEvent := range allEvents.DueDateEvent.Pods {
		log.Printf(".Pod %s is due on %s", podEvent.Name, podEvent.DueDate)
	}
	return nil
}
