package recipes

import (
	"development/go/recipes/lib"
	"development/go/recipes/lib/endpoints/project_keys/project_keys.2.lists"
	"development/go/recipes/lib/structs/common"

	projectKeys1 "development/go/recipes/lib/endpoints/project_keys/project_keys.1"
	recipes2 "development/go/recipes/lib/helpers/recipes"
	request2 "development/go/recipes/lib/structs/request"
	response2 "development/go/recipes/lib/structs/response"

	log "github.com/sirupsen/logrus"
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
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	user, err := recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	projectKey, err := recipes2.AddProjectKey(user, ProjectKeyName)
	if err != nil {
		return
	}

	projectBlock, err := recipes2.AddBlock(user, ProjectBlockName, projectKey)
	if err != nil {
		return
	}

	log.Info("Add 2 project lists")
	recipes2.SleepBefore()
	projectList1, err := addProjectList(user, ProjectList1Name, projectBlock)
	if err != nil {
		return
	}
	projectList2, err := addProjectList(user, ProjectList2Name, projectBlock)
	if err != nil {
		return
	}
	log.Printf(".Both project lists, %s and %s created successfully", projectList1.Name, projectList2.Name)
	recipes2.SleepAfter()

	log.Info("Add a project pod into a project list")
	recipes2.SleepBefore()
	projectPod, err := addProjectPod(user, ProjectPodName, projectList1)
	if err != nil {
		return
	}
	log.Printf(".Project pod %s created inside %s successfully", projectPod.Name, projectList1.Name)
	recipes2.SleepAfter()

	log.Printf("Move project pod %s between project lists", projectPod.Name)
	recipes2.SleepBefore()
	err = movePodBetweenLists(user, projectList1, projectList2, projectPod)
	if err != nil {
		return
	}
	log.Printf(".Project pod %s moved from list %s to list %s successfully", projectPod.Name,
		projectList1.Name, projectList2.Name)
	recipes2.SleepAfter()
}

func addProjectPod(user response2.User, podName string, projectList response2.ProjectList) (response2.Pod, error) {
	newPod, err := projectKeys1.AddProjectPod(
		user.JwtToken,
		request2.AddPodReqBody{Name: podName},
		request2.ProjectListIdParam{
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

func addProjectList(user response2.User, projectListName string, block response2.Block) (response2.ProjectList, error) {
	newProjectList, err := projectKeys.AddProjectBlockList(
		user.JwtToken,
		request2.AddProjectListReqBody{Name: projectListName},
		common.ResourceIdParam{BlockId: block.ID, KeyId: block.Key.ID},
	)
	if err != nil {
		return newProjectList, err
	}
	return newProjectList, nil
}

func movePodBetweenLists(
	user response2.User,
	list1 response2.ProjectList,
	list2 response2.ProjectList,
	pod response2.Pod,
) error {
	err := projectKeys.BulkMovePodsInProjectList(user.JwtToken, request2.CopyMoveProjectListPodsParam{
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
