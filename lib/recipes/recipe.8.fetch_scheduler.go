package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/blocks/blocks.1"
	"development/go/recipes/lib/endpoints/key_pods/key_pods.1"
	"development/go/recipes/lib/endpoints/scheduler"
	"development/go/recipes/lib/structs/common"

	recipes2 "development/go/recipes/lib/helpers/recipes"
	request2 "development/go/recipes/lib/structs/request"
	response2 "development/go/recipes/lib/structs/response"

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
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	key, err := recipes2.AddCustomKey(user, SchedulerKeyName)
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

func setBlockDueDate(user response2.User, key response2.Key) error {
	block, err := recipes2.AddBlock(user, SchedulerBlockName, key)
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

func setPodDueDate(user response2.User, key response2.Key) error {
	pod, err := recipes2.AddPod(user, SchedulerPodName, key)
	if err != nil {
		return err
	}
	dueDate := DueDate
	_, err = keyPods.UpdateKeyPod(
		user.JwtToken,
		request2.UpdatePodReqBody{DueDate: &dueDate},
		common.ResourceIdParam{PodId: pod.ID, KeyId: pod.Key.ID})
	if err != nil {
		return err
	}
	return nil
}

func fetchSchedulerEvents(user response2.User) error {
	allEvents, err := scheduler.GetEventsForGivenDay(user.JwtToken, request2.EventDateParam{StartDate: DueDate})
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
