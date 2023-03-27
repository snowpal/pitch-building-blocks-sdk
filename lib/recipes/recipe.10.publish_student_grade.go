package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/scales"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/helpers/recipes"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"

	blockPods "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/block_pods/block_pods.1"
	teacherKeys "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/teacher_keys/teacher_keys.2.teachers"

	log "github.com/sirupsen/logrus"
)

const (
	TeacherKeyName   = "School"
	TeacherBlockName = "Math"
	TeacherPodName   = "Final Exam"
)

func PublishStudentGrade() {
	log.Info("Objective: Assign and Publish Pod Grades for a Student")
	_, err := recipes.ValidateDependencies()
	if err != nil {
		return
	}

	var user response.User
	user, err = recipes.SignIn(lib.ActiveUser, lib.Password)
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

	err = recipes.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}

	err = publishPod(user, pod)
	if err != nil {
		return
	}

	var readUser response.User
	readUser, err = recipes.SignIn(lib.ReadUser, lib.Password)
	if err != nil {
		return
	}

	log.Printf("Assign grade to %s pod for %s", pod.Name, lib.ReadUser)
	recipes.SleepBefore()
	err = assignPodGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade assigned successfully to %s", lib.ReadUser)
	recipes.SleepAfter()

	log.Printf("Publish grade for %s", readUser.Email)
	recipes.SleepBefore()
	err = publishPodGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade published successfully for %s", lib.ReadUser)
	recipes.SleepAfter()
}

func publishPod(user response.User, pod response.Pod) error {
	_, err := blockPods.UpdateBlockPodCompletionStatus(
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
	resScales, _ := scales.GetScales(user.JwtToken, scales.GetScalesParam{
		IncludeCounts: false,
		ExcludeEmpty:  true,
	})
	var alphabeticScale response.Scale
	for _, scale := range resScales {
		if *scale.Type == lib.AlphabeticScaleType {
			alphabeticScale = scale
			break
		}
	}
	podId := pod.ID
	blockId := pod.Block.ID
	err := blockPods.AddScaleToBlockPod(user.JwtToken, request.ScaleIdParam{
		ScaleId: alphabeticScale.ID,
		PodId:   &podId,
		BlockId: &blockId,
		KeyId:   pod.Key.ID,
	})
	if err != nil {
		return err
	}
	_, err = teacherKeys.AssignPodGradeForAStudentAsTeacher(
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
	err := teacherKeys.BulkPublishPodGradesForAStudent(
		user.JwtToken,
		teacherKeys.PublishStudentGradeReqBody{PodIds: pod.ID},
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
