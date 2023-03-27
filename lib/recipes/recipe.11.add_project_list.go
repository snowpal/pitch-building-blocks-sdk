package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	log "github.com/sirupsen/logrus"
	projectKeys "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/project_keys/project_keys.1"
	projectLists "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/project_keys/project_keys.2.lists"
)

const (
	ProjectKeyName   = "Go Development"
	ProjectBlockName = "Status API"
	ProjectPodName   = "Define Endpoints"
	ProjectList1Name = "Statuses"
	ProjectList2Name = "Teams"
)

func AddProjectList() {
	log.Info("Objective: Add Project Lists, Project Pods, Move Pod between Lists")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	projectKey, err := recipes.AddProjectKey(user, ProjectKeyName)
	if err != nil {
		return
	}

	projectBlock, err := recipes.AddBlock(user, ProjectBlockName, projectKey)
	if err != nil {
		return
	}

	log.Info("Add 2 project lists")
	recipes.SleepBefore()
	projectList1, err := addProjectList(user, ProjectList1Name, projectBlock)
	if err != nil {
		return
	}
	projectList2, err := addProjectList(user, ProjectList2Name, projectBlock)
	if err != nil {
		return
	}
	log.Printf(".Both project lists, %s and %s created successfully", projectList1.Name, projectList2.Name)
	recipes.SleepAfter()

	log.Info("Add a project pod into a project list")
	recipes.SleepBefore()
	projectPod, err := addProjectPod(user, ProjectPodName, projectList1)
	if err != nil {
		return
	}
	log.Printf(".Project pod %s created inside %s successfully", projectPod.Name, projectList1.Name)
	recipes.SleepAfter()

	log.Printf("Move project pod %s between project lists", projectPod.Name)
	recipes.SleepBefore()
	err = movePodBetweenLists(user, projectList1, projectList2, projectPod)
	if err != nil {
		return
	}
	log.Printf(".Project pod %s moved from list %s to list %s successfully", projectPod.Name,
		projectList1.Name, projectList2.Name)
	recipes.SleepAfter()
}

func addProjectPod(user response.User, podName string, projectList response.ProjectList) (response.Pod, error) {
	newPod, err := projectKeys.AddProjectPod(
		user.JwtToken,
		request.AddPodReqBody{Name: podName},
		request.ProjectListIdParam{
			ProjectListId: projectList.ID,
			BlockId:       projectList.Block.ID,
			KeyId:         projectList.Key.ID,
		},
	)
	if err != nil {
		return newPod, err
	}
	return newPod, nil
}

func addProjectList(user response.User, projectListName string, block response.Block) (response.ProjectList, error) {
	newProjectList, err := projectLists.AddProjectBlockList(
		user.JwtToken,
		request.AddProjectListReqBody{Name: projectListName},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID},
	)
	if err != nil {
		return newProjectList, err
	}
	return newProjectList, nil
}

func movePodBetweenLists(
	user response.User,
	list1 response.ProjectList,
	list2 response.ProjectList,
	pod response.Pod,
) error {
	err := projectLists.BulkMovePodsInProjectList(user.JwtToken, request.CopyMoveProjectListPodsParam{
		ProjectListId:       list1.ID,
		BlockId:             pod.Block.ID,
		KeyId:               pod.Key.ID,
		TargetProjectListId: list2.ID,
		TargetBlockId:       list2.Block.ID,
		TargetKeyId:         list2.Key.ID,
		PodIds:              []string{pod.ID},
	})
	if err != nil {
		return err
	}
	return nil
}
