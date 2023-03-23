package main

import (
	"development/go/recipes/lib/golang"
	"development/go/recipes/lib/golang/endpoints/block_pods.1"
	"development/go/recipes/lib/golang/endpoints/scales.1"
	teacherKeys2 "development/go/recipes/lib/golang/endpoints/teacher_keys.2.teachers"
	"development/go/recipes/lib/golang/helpers/recipes"
	"development/go/recipes/lib/golang/structs/common"
	"development/go/recipes/lib/golang/structs/request"
	"development/go/recipes/lib/golang/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	TeacherKeyName   = "School"
	TeacherBlockName = "Math"
	TeacherPodName   = "Final Exam"
)

func main() {
	log.Info("Objective: Assign and Publish Pod Grades for a Student")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(golang.ActiveUser, golang.Password)
	if err != nil {
		return
	}

	var key response.Key
	key, err = recipes.AddTeacherKey(user, TeacherKeyName)
	if err != nil {
		return
	}

	var block response.Block
	block, err = recipes.AddBlock(user, TeacherBlockName, key)
	if err != nil {
		return
	}

	var pod response.Pod
	pod, err = recipes.AddPodToBlock(user, TeacherPodName, block)
	if err != nil {
		return
	}

	err = recipes.SearchUserAndShareBlock(user, block, golang.ReadUserToken, golang.ReadAcl)
	if err != nil {
		return
	}

	err = publishPod(user, pod)
	if err != nil {
		return
	}

	var readUser response.User
	readUser, err = recipes.SignIn(golang.ReadUser, golang.Password)
	if err != nil {
		return
	}

	log.Printf("Assign grade to %s pod for %s", pod.Name, golang.ReadUser)
	recipes.SleepBefore()
	err = assignPodGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade assigned successfully to %s", golang.ReadUser)
	recipes.SleepAfter()

	log.Printf("Publish grade for %s", readUser.Email)
	recipes.SleepBefore()
	err = publishPodGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade published successfully for %s", golang.ReadUser)
	recipes.SleepAfter()
}

func publishPod(user response.User, pod response.Pod) error {
	_, err := block_pods.UpdateBlockPodCompletionStatus(
		user.JwtToken,
		request.UpdatePodStatusReqBody{Completed: true},
		common.ResourceIdParam{
			PodId:   pod.ID,
			BlockId: pod.Block.ID,
			KeyId:   pod.Key.ID,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func assignPodGrade(user response.User, pod response.Pod, student response.User) error {
	resScales, err := scales.GetScales(user.JwtToken, scales.GetScalesParam{
		IncludeCounts: false,
		ExcludeEmpty:  true,
	})
	var alphabeticScale response.Scale
	for _, scale := range resScales {
		if *scale.Type == golang.AlphabeticScaleType {
			alphabeticScale = scale
			break
		}
	}
	podId := pod.ID
	blockId := pod.Block.ID
	err = block_pods.AddScaleToBlockPod(user.JwtToken, request.ScaleIdParam{
		ScaleId: alphabeticScale.ID,
		PodId:   &podId,
		BlockId: &blockId,
		KeyId:   pod.Key.ID,
	})
	if err != nil {
		return err
	}
	_, err = teacherKeys2.AssignPodGradeForAStudentAsTeacher(
		user.JwtToken,
		request.UpdateScaleValueReqBody{ScaleValue: alphabeticScale.ScaleValues[0]},
		request.ClassroomIdParam{
			StudentId: student.ID,
			ResourceIds: common.ResourceIdParam{
				PodId:   podId,
				BlockId: blockId,
				KeyId:   pod.Key.ID,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func publishPodGrade(user response.User, pod response.Pod, student response.User) error {
	err := teacherKeys2.BulkPublishPodGradesForAStudent(
		user.JwtToken,
		teacherKeys2.PublishStudentGradeReqBody{PodIds: pod.ID},
		request.ClassroomIdParam{
			StudentId: student.ID,
			ResourceIds: common.ResourceIdParam{
				BlockId: pod.Block.ID,
				KeyId:   pod.Key.ID,
			},
		},
	)
	if err != nil {
		return err
	}
	return nil
}
