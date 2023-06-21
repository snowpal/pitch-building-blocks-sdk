package recipes

import (
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/common"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/request"
	"github.com/snowpal/pitch-building-blocks-sdk/lib/structs/response"
	"strings"

	blockCollaboration "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/collaboration/collaboration.1.blocks"
	blockPodCollaboration "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/collaboration/collaboration.2.block_pods"
	podCollaboration "github.com/snowpal/pitch-building-blocks-sdk/lib/endpoints/collaboration/collaboration.3.key_pods"
)

func SearchUserAndShareBlock(user response.User, block response.Block, searchEmail string, acl string) error {
	blockIdParam := common.ResourceIdParam{
		BlockId: block.ID,
		KeyId:   block.Key.ID,
	}
	searchedUsers, err := blockCollaboration.GetUsersThisBlockCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: strings.Split(searchEmail, "@")[0],
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}

	// For the purpose of this recipe, it does not matter which user from the list we end up picking, hence we go with
	// the first one.
	_, err = blockCollaboration.ShareBlockWithCollaborator(
		user.JwtToken,
		request.BlockAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: blockIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}

func SearchUserAndSharePod(user response.User, pod response.Pod, searchToken string, acl string) error {
	podIdParam := common.ResourceIdParam{
		PodId: pod.ID,
		KeyId: pod.Key.ID,
	}
	searchedUsers, err := podCollaboration.GetUsersThisKeyPodCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: strings.Split(searchToken, "@")[0],
			ResourceIds: podIdParam,
		})
	if err != nil {
		return err
	}
	_, err = podCollaboration.ShareKeyPodWithCollaborator(
		user.JwtToken,
		request.PodAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: podIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}

func SearchUserAndShareBlockPod(user response.User, blockPod response.Pod, searchToken string, acl string) error {
	blockPodIdParam := common.ResourceIdParam{
		PodId:   blockPod.ID,
		KeyId:   blockPod.Key.ID,
		BlockId: blockPod.Block.ID,
	}
	searchedUsers, err := blockPodCollaboration.GetUsersThisBlockPodCanBeSharedWith(
		user.JwtToken,
		common.SearchUsersParam{
			SearchToken: strings.Split(searchToken, "@")[0],
			ResourceIds: blockPodIdParam,
		})
	if err != nil {
		return err
	}
	_, err = blockPodCollaboration.ShareBlockPodWithCollaborator(
		user.JwtToken,
		request.PodAclReqBody{Acl: acl},
		common.AclParam{
			UserId:      searchedUsers[0].ID,
			ResourceIds: blockPodIdParam,
		})
	if err != nil {
		return err
	}
	return nil
}
