package recipes

import (
	"development/go/recipes/lib"
	blockPods2 "development/go/recipes/lib/endpoints/block_pods/block_pods.1"
	"development/go/recipes/lib/endpoints/scales"
	teacherKeys2 "development/go/recipes/lib/endpoints/teacher_keys/teacher_keys.2.teachers"
	recipes2 "development/go/recipes/lib/helpers/recipes"
	"development/go/recipes/lib/structs/common"
	request2 "development/go/recipes/lib/structs/request"
	response2 "development/go/recipes/lib/structs/response"
	log "github.com/sirupsen/logrus"
)

const (
	TeacherKeyName   = "School"
	TeacherBlockName = "Math"
	TeacherPodName   = "Final Exam"
)

func PublishStudentGrade() {
	log.Info("Objective: Assign and Publish Pod Grades for a Student")
	_, err := recipes2.ValidateDependencies()
	if err != nil {
		return
	}

	var user response2.User
	user, err = recipes2.SignIn(lib.ActiveUser, lib.Password)
	if err != nil {
		return
	}

	var key response2.Key
	key, err = recipes2.AddTeacherKey(user, TeacherKeyName)
	if err != nil {
		return
	}

	var block response2.Block
	block, err = recipes2.AddBlock(user, TeacherBlockName, key)
	if err != nil {
		return
	}

	var pod response2.Pod
	pod, err = recipes2.AddPodToBlock(user, TeacherPodName, block)
	if err != nil {
		return
	}

	err = recipes2.SearchUserAndShareBlock(user, block, "api_read_user", lib.ReadAcl)
	if err != nil {
		return
	}

	err = publishPod(user, pod)
	if err != nil {
		return
	}

	var readUser response2.User
	readUser, err = recipes2.SignIn(lib.ReadUser, lib.Password)
	if err != nil {
		return
	}

	log.Printf("Assign grade to %s pod for %s", pod.Name, lib.ReadUser)
	recipes2.SleepBefore()
	err = assignPodGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade assigned successfully to %s", lib.ReadUser)
	recipes2.SleepAfter()

	log.Printf("Publish grade for %s", readUser.Email)
	recipes2.SleepBefore()
	err = publishPodGrade(user, pod, readUser)
	if err != nil {
		return
	}
	log.Printf(".Grade published successfully for %s", lib.ReadUser)
	recipes2.SleepAfter()
}

func publishPod(user response2.User, pod response2.Pod) error {
	_, err := blockPods2.UpdateBlockPodCompletionStatus(
		user.JwtToken,
		request2.UpdatePodStatusReqBody{Completed: true},
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

func assignPodGrade(user response2.User, pod response2.Pod, student response2.User) error {
	resScales, err := scales.GetScales(user.JwtToken, scales.GetScalesParam{
		IncludeCounts: false,
		ExcludeEmpty:  true,
	})
	var alphabeticScale response2.Scale
	for _, scale := range resScales {
		if *scale.Type == lib.AlphabeticScaleType {
			alphabeticScale = scale
			break
		}
	}
	podId := pod.ID
	blockId := pod.Block.ID
	err = blockPods2.AddScaleToBlockPod(user.JwtToken, request2.ScaleIdParam{
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
		request2.UpdateScaleValueReqBody{ScaleValue: alphabeticScale.ScaleValues[0]},
		request2.ClassroomIdParam{
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

func publishPodGrade(user response2.User, pod response2.Pod, student response2.User) error {
	err := teacherKeys2.BulkPublishPodGradesForAStudent(
		user.JwtToken,
		teacherKeys2.PublishStudentGradeReqBody{PodIds: pod.ID},
		request2.ClassroomIdParam{
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
