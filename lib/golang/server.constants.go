package golang

const XApiKey = "wf8sHELzWp9MGizZME5Zsjk4IntZS0e8mdYMYjjg"

const (
	HostProd    = "https://gateway.snowpal.com/"
	HostStaging = "https://gateway-staging.snowpal.com/"
	HostDev     = "https://gateway-dev.snowpal.com/"
)

const (
	RouteAttributesGetDisplayableAttributesOfKey   = "app/resource/attributes"
	RouteAttributesUpdateKeyDisplayAttributes      = "keys/%s/attributes"
	RouteAttributesUpdateBlockDisplayAttributes    = "blocks/%s/attributes"
	RouteAttributesUpdateKeyPodDisplayAttributes   = "pods/%s/attributes"
	RouteAttributesUpdateBlockPodDisplayAttributes = "block-pods/%s/attributes"
)

const (
	RouteBlocksGetBlocksInAKey                          = "keys/%s/blocks?filter=%s&batchIndex=%s&aclWriteOrHigher=%s"
	RouteBlocksAddBlock                                 = "keys/%s/blocks"
	RouteBlocksGetBlocksLinkedToPod                     = ""
	RouteBlocksAddBlockBasedOnTemplate                  = ""
	RouteBlocksLinkBlockToKey                           = ""
	RouteBlocksUnlinkBlockFromKey                       = ""
	RouteBlocksGetAllBlocksAvailableToBeLinkedToThisKey = ""
	RouteBlocksGetBlock                                 = ""
	RouteBlocksUpdateBlock                              = ""
	RouteBlocksAddBlockTypeToBlock                      = ""
	RouteBlocksDeleteBlockTypeFromBlock                 = ""
	RouteBlocksAddScaleToBlock                          = ""
	RouteBlocksDeleteScaleFromBlock                     = ""
	RouteBlocksAddScaleValueToBlock                     = ""
	RouteBlocksUpdateBlockDescription                   = ""
	RouteBlocksArchiveBlock                             = ""
	RouteBlocksUnarchiveBlock                           = ""
	RouteBlocksGetAllArchivedBlocks                     = ""
	RouteBlocksBulkArchiveBlocks                        = ""
	RouteBlocksAllowArchivalOfBlock                     = ""
	RouteBlocksCopyBlock                                = ""
	RouteBlocksMoveBlock                                = ""
)

const (
	RouteBlocksGetBlockAttachments   = ""
	RouteBlocksAddBlockAttachment    = ""
	RouteBlocksRenameBlockAttachment = ""
	RouteBlocksDeleteBlockAttachment = ""
)

const (
	RouteBlocksGetLinkedBlockPods           = ""
	RouteBlocksGetScaleValuesForScale       = ""
	RouteBlocksGetTaskStatusForBlock        = ""
	RouteBlocksGetBlockGradesForAllStudents = ""
)

const (
	RouteBlocksGetBlockChecklists     = ""
	RouteBlocksAddBlockChecklist      = ""
	RouteBlocksReorderBlockChecklists = ""
	RouteBlocksRenameChecklistTitle   = ""
	RouteBlocksDeleteChecklist        = ""
	RouteBlocksAddChecklistItem       = ""
	RouteBlocksUpdateChecklistItem    = ""
	RouteBlocksDeleteChecklistItem    = ""
	RouteBlocksReorderChecklistItems  = ""
)

const (
	RouteBlocksGetBlockComments   = ""
	RouteBlocksAddBlockComment    = ""
	RouteBlocksUpdateBlockComment = ""
	RouteBlocksDeleteBlockComment = ""
)

const (
	RouteBlocksGetBlockNotes   = ""
	RouteBlocksAddBlockNote    = ""
	RouteBlocksUpdateBlockNote = ""
	RouteBlocksDeleteBlockNote = ""
)

const (
	RouteBlocksGetBlockTasks     = ""
	RouteBlocksAddBlockTask      = ""
	RouteBlocksUpdateBlockTask   = ""
	RouteBlocksDeleteBlockTask   = ""
	RouteBlocksAssignBlockTask   = ""
	RouteBlocksUnassignBlockTask = ""
	RouteBlocksReorderBlockTasks = ""
)

const (
	RouteBlockPodsGetBlockPods                   = ""
	RouteBlockPodsAddBlockPod                    = ""
	RouteBlockPodsAddBlockPodBasedOnTemplate     = ""
	RouteBlockPodsLinkPodToBlock                 = ""
	RouteBlockPodsUnlinkPodFromBlock             = ""
	RouteBlockPodsGetBlockPod                    = ""
	RouteBlockPodsUpdateBlockPod                 = ""
	RouteBlockPodsUpdateBlockPodCompletionStatus = ""
	RouteBlockPodsAddPodTypeToBlockPod           = ""
	RouteBlockPodsRemovePodTypeFromBlockPod      = ""
	RouteBlockPodsAddScaleToBlockPod             = ""
	RouteBlockPodsDeleteScaleFromBlockPod        = ""
	RouteBlockPodsUpdateBlockPodScaleValue       = ""
	RouteBlockPodsArchiveBlockPod                = ""
	RouteBlockPodsGetArchivedBlockPods           = ""
	RouteBlockPodsGetPodsAvailableToBeLinked     = ""
	RouteBlockPodsUnarchiveBlockPod              = ""
	RouteBlockPodsBulkArchiveBlockPods           = ""
	RouteBlockPodsUpdateBlockPodDescription      = ""
	RouteBlockPodsAllowArchivalOfBlockPod        = ""
	RouteBlockPodsCopyBlockPod                   = ""
	RouteBlockPodsMoveBlockPod                   = ""
)

const (
	RouteBlockPodsGetBlockPodAttachments   = ""
	RouteBlockPodsAddBlockPodAttachment    = ""
	RouteBlockPodsRenameBlockPodAttachment = ""
	RouteBlockPodsDeleteBlockPodAttachment = ""
)

const (
	RouteBlockPodsGetBlockPodTasksForCharts       = ""
	RouteBlockPodsGetBlockPodGradesForAllStudents = ""
)

const (
	RouteBlockPodsGetBlockPodChecklists         = ""
	RouteBlockPodsAddBlockPodChecklist          = ""
	RouteBlockPodsReorderBlockPodChecklists     = ""
	RouteBlockPodsDeleteBlockPodChecklist       = ""
	RouteBlockPodsRenameBlockPodChecklist       = ""
	RouteBlockPodsAddBlockPodChecklistItem      = ""
	RouteBlockPodsUpdateBlockPodChecklistItem   = ""
	RouteBlockPodsDeleteBlockPodChecklistItem   = ""
	RouteBlockPodsReorderBlockPodChecklistItems = ""
)

const (
	RouteBlockPodsGetBlockPodComments   = ""
	RouteBlockPodsAddBlockPodComment    = ""
	RouteBlockPodsUpdateBlockPodComment = ""
	RouteBlockPodsDeleteBlockPodComment = ""
)

const (
	RouteBlockPodsGetBlockPodNotes   = ""
	RouteBlockPodsAddBlockPodNote    = ""
	RouteBlockPodsUpdateBlockPodNote = ""
	RouteBlockPodsDeleteBlockPodNote = ""
)

const (
	RouteBlockPodsGetBlockPodTasks     = ""
	RouteBlockPodsAddBlockPodTask      = ""
	RouteBlockPodsUpdateBlockPodTask   = ""
	RouteBlockPodsDeleteBlockPodTask   = ""
	RouteBlockPodsAssignBlockPodTask   = ""
	RouteBlockPodsUnassignBlockPodTask = ""
	RouteBlockPodsReorderBlockPodTasks = ""
)

const (
	RouteBlockTypesGetBlockTypes           = ""
	RouteBlockTypesAddBlockType            = ""
	RouteBlockTypesUpdateBlockType         = ""
	RouteBlockTypesDeleteBlockType         = ""
	RouteBlockTypesGetBlocksUsingBlockType = ""
)

const (
	RouteCollaboration1Blocks                             = ""
	RouteCollaborationGetBlockCollaborators               = ""
	RouteCollaborationUpdateCollaboratorsAccessLevel      = ""
	RouteCollaborationUnshareBlockFromCollaborator        = ""
	RouteCollaborationShareBlockWithUser                  = ""
	RouteCollaborationShareBlockWithUserAlongWithPods     = ""
	RouteCollaborationGetAllUsersThisBlockCanBeSharedWith = ""
	RouteCollaborationBulkShareBlocksWithCollaborators    = ""
	RouteCollaborationLeaveBlock                          = ""
)

const (
	RouteCollaborationGetBlockPodCollaborators               = ""
	RouteCollaborationShareBlockPodWithUser                  = ""
	RouteCollaborationUnshareBlockPodFromUser                = ""
	RouteCollaborationBulkShareBlockPodsWithCollaborators    = ""
	RouteCollaborationGetAllUsersThisBlockPodCanBeSharedWith = ""
	RouteCollaborationUpdateBlockPodACL                      = ""
	RouteCollaborationLeaveBlockPod                          = ""
)

const (
	RouteCollaborationGetKeyPodCollaborators               = ""
	RouteCollaborationShareKeyPodWithUser                  = ""
	RouteCollaborationBulkSharePodsWithCollaborators       = ""
	RouteCollaborationUnshareKeyPodFromUser                = ""
	RouteCollaborationGetAllUsersThisKeyPodCanBeSharedWith = ""
	RouteCollaborationUpdatePodAcl                         = ""
	RouteCollaborationLeavePod                             = ""
)

const (
	RouteCommentsGetRecentComments = ""
)

const (
	RouteConversationsGetUnreadConversationsCount         = ""
	RouteConversationsGetUserConversations                = ""
	RouteConversationsAddPrivateOrGroupConversation       = ""
	RouteConversationsGetConversationForGivenUsernames    = ""
	RouteConversationsSendMessageToAnExistingConversation = ""
	RouteConversationsGetConversation                     = ""
	RouteConversationsDeleteConversation                  = ""
	RouteConversationsLeaveConversation                   = ""
	RouteConversationsArchiveConversation                 = ""
)

const (
	RouteDashboardGetDashboardDetails              = ""
	RouteDashboardGetRecentlyModifiedBlocksAndPods = ""
	RouteDashboardGetUnreadCount                   = ""
	RouteDashboardGetRecentlyModifiedKeys          = ""
	RouteDashboardGetPodsAndTasksDueShortly        = ""
	RouteDashboardGetBlocksDueShortly              = ""
	RouteDashboardGetUnreadNotifications           = ""
	RouteDashboardGetUnreadConversations           = ""
)

const (
	RouteDashboardGetUserKeysBlocksAndPods           = ""
	RouteDashboardGetSystemKeysBlocksAndPods         = ""
	RouteDashboardGetFilteredUserKeysBlocksAndPods   = ""
	RouteDashboardGetFilteredSystemKeysBlocksAndPods = ""
	RouteDashboardGetBlocksBasedOnBlockTypes         = ""
	RouteDashboardGetPodsBasedOnPodTypes             = ""
	RouteDashboardGetBlocksAndPodsBasedOnScales      = ""
	RouteDashboardGetTasksByStatus                   = ""
)

const (
	RouteFavoritesGetFavorites          = ""
	RouteFavoritesAddKeyAsFavorite      = ""
	RouteFavoritesAddBlockAsFavorite    = ""
	RouteFavoritesAddPodAsFavorite      = ""
	RouteFavoritesAddBlockPodAsFavorite = ""
	RouteFavoritesDeleteFavorite        = ""
)

const (
	RouteFollowersAddUserToFollowUsList = ""
	RouteFollowersGetFollowers          = ""
)

const (
	RouteKeysGetAllKeys              = "keys"
	RouteKeysAddKey                  = ""
	RouteKeysGetKey                  = ""
	RouteKeysUpdateKey               = ""
	RouteKeysGetArchivedKeys         = ""
	RouteKeysGetKeysAPodIsLinkedTo   = ""
	RouteKeysGetKeysABlockIsLinkedTo = ""
	RouteKeysGetKeysFilteredByType   = ""
	RouteKeysBulkArchiveKeys         = ""
	RouteKeysArchiveKey              = ""
	RouteKeysUnarchiveKey            = ""
	RouteKeysUpdateKeyDescription    = ""
)

const (
	RouteKeysGetBlocksAndPodsAssociatedWithKey           = ""
	RouteKeysGetFilteredUserKeysBlocksAndPodsForGivenKey = ""
	RouteKeysGetBlockTypesAndBlocksBasedOnThemInKey      = ""
	RouteKeysGetPodsBasedOnPodTypesInKey                 = ""
	RouteKeysGetScalesAlongWithBlocksAndPodsBasedOnThem  = ""
	RouteKeysGetLinkedResources                          = ""
	RouteKeysGetKeyPodAndBlockScaleValues                = ""
	RouteKeysGetTaskStatus                               = ""
)

const (
	RouteKeysGetKeyChecklists         = ""
	RouteKeysAddKeyChecklist          = ""
	RouteKeysReorderKeyChecklists     = ""
	RouteKeysRenameChecklist          = ""
	RouteKeysDeleteKeyChecklist       = ""
	RouteKeysAddKeyChecklistItem      = ""
	RouteKeysUpdateKeyChecklistItem   = ""
	RouteKeysDeleteKeyChecklistItem   = ""
	RouteKeysReorderKeyChecklistItems = ""
)

const (
	RouteKeysGetKeyNotes   = ""
	RouteKeysAddKeyNote    = ""
	RouteKeysUpdateKeyNote = ""
	RouteKeysDeleteKeyNote = ""
)

const (
	RouteKeysGetKeyTasks    = ""
	RouteKeysAddKeyTask     = ""
	RouteKeysUpdateKeyTask  = ""
	RouteKeysDeleteKeyTask  = ""
	RouteKeysReorderKeyTask = ""
)

const (
	RouteKeyPodsGetKeyPods                    = ""
	RouteKeyPodsAddKeyPod                     = ""
	RouteKeyPodsGetKeyPodsAvailableToBeLinked = ""
	RouteKeyPodsLinkKeyPodToKey               = ""
	RouteKeyPodsUnlinkKeyPodFromKey           = ""
	RouteKeyPodsGetKeyPod                     = ""
	RouteKeyPodsUpdateKeyPod                  = ""
	RouteKeyPodsUpdateKeyPodsCompletionStatus = ""
	RouteKeyPodsUpdateKeyPodsScaleValue       = ""
	RouteKeyPodsAddPodTypeToKeyPod            = ""
	RouteKeyPodsRemovePodTypeFromKeyPod       = ""
	RouteKeyPodsAddScaleToPod                 = ""
	RouteKeyPodsDeleteScaleFromKeyPod         = ""
	RouteKeyPodsArchiveKeyPod                 = ""
	RouteKeyPodsGetArchivedPods               = ""
	RouteKeyPodsUnarchiveKeyPod               = ""
	RouteKeyPodsBulkArchiveKeyPods            = ""
	RouteKeyPodsUpdateKeyPodDescription       = ""
	RouteKeyPodsAllowArchivalOfKeyPod         = ""
	RouteKeyPodsCopyKeyPod                    = ""
	RouteKeyPodsMoveKeyPod                    = ""
)

const (
	RouteKeyPodsGetKeyPodAttachments   = ""
	RouteKeyPodsAddAttachmentToKeyPod  = ""
	RouteKeyPodsRenameKeyPodAttachment = ""
	RouteKeyPodsDeleteKeyPodAttachment = ""
)

const (
	RouteKeyPods3Checklists                 = ""
	RouteKeyPodsGetKeyPodChecklists         = ""
	RouteKeyPodsAddKeyPodChecklist          = ""
	RouteKeyPodsReorderKeyPodChecklists     = ""
	RouteKeyPodsDeleteKeyPodChecklist       = ""
	RouteKeyPodsRenameKeyPodChecklist       = ""
	RouteKeyPodsAddKeyPodChecklistItem      = ""
	RouteKeyPodsUpdateKeyPodChecklistItem   = ""
	RouteKeyPodsDeleteKeyPodChecklistItem   = ""
	RouteKeyPodsReorderKeyPodChecklistItems = ""
)

const (
	RouteKeyPodsGetKeyPodComments   = ""
	RouteKeyPodsAddKeyPodComment    = ""
	RouteKeyPodsUpdateKeyPodComment = ""
	RouteKeyPodsDeleteKeyPodComment = ""
)

const (
	RouteKeyPodsGetKeyPodNotes   = ""
	RouteKeyPodsAddKeyPodNote    = ""
	RouteKeyPodsUpdateKeyPodNote = ""
	RouteKeyPodsDeleteKeyPodNote = ""
)

const (
	RouteKeyPodsGetKeyPodTasks     = ""
	RouteKeyPodsAddKeyPodTask      = ""
	RouteKeyPodsUpdateKeyPodTask   = ""
	RouteKeyPodsDeleteKeyPodTask   = ""
	RouteKeyPodsAssignKeyPodTask   = ""
	RouteKeyPodsUnassignKeyPodTask = ""
	RouteKeyPodsReorderKeyPodTasks = ""
)

const (
	RouteNotificationsGetAllNotifications           = ""
	RouteNotificationsGetUnreadNotifications        = ""
	RouteNotificationsGetUnreadNotificationCount    = ""
	RouteNotificationsMarkNotificationAsRead        = ""
	RouteNotificationsMarkNotificationsAsReadInBulk = ""
)

const (
	RoutePodTypesGetPodTypes         = ""
	RoutePodTypesAddPodType          = ""
	RoutePodTypesUpdatePodType       = ""
	RoutePodTypesDeletePodType       = ""
	RoutePodTypesGetPodsUsingPodType = ""
)

const (
	RouteProfileGetUsersIntroductionProfile   = ""
	RouteProfileGetUsersProfile               = ""
	RouteProfileUpdateUsersProfile            = ""
	RouteProfileUpdateUsername                = ""
	RouteProfileBlocksUserFromSendingMessages = ""
	RouteProfileUnblocksUser                  = ""
)

const (
	RouteProjectKeysAddAProjectPod               = ""
	RouteProjectKeysAddProjectPodBasedOnTemplate = ""
	RouteProjectKeysLinkProjectPodToBlock        = ""
	RouteProjectKeysReorderProjectPods           = ""
	RouteProjectKeysCopyProjectPod               = ""
	RouteProjectKeysMoveProjectPod               = ""
	RouteProjectKeysCopyProjectBlock             = ""
	RouteProjectKeysAssignProjectPod             = ""
	RouteProjectKeysUnassignProjectPod           = ""
)

const (
	RouteProjectKeysAddProjectBlockList       = ""
	RouteProjectKeysGetProjectLists           = ""
	RouteProjectKeysCopyAllPodsInProjectList  = ""
	RouteProjectKeysBulkCopyPodsInProjectList = ""
	RouteProjectKeysMoveAllPodsInProjectList  = ""
	RouteProjectKeysBulkMovePodsInProjectList = ""
	RouteProjectKeysMoveProjectList           = ""
	RouteProjectKeysGetProjectList            = ""
	RouteProjectKeysRenameProjectList         = ""
	RouteProjectKeysReorderProjectList        = ""
	RouteProjectKeysArchiveProjectList        = ""
)

const (
	RouteRegistrationRegisterNewUserByEmail = "app/users/sign-up"
	RouteRegistrationSignInByEmail          = "app/users/sign-in"
	RouteRegistrationResetPassword          = ""
	RouteRegistrationActivateUser           = "app/user-verified/%s"
)

const (
	RouteRelationsGetRelationsMatchingSearchToken = ""
	RouteRelationsGetRelationsForKey              = ""
	RouteRelationsGetRelationsForBlock            = ""
	RouteRelationsGetRelationsForPod              = ""
	RouteRelationsGetRelationsForBlockPod         = ""
	RouteRelationsRelateKeyToKey                  = ""
	RouteRelationsUnrelateKeyFromKey              = ""
	RouteRelationsRelateBlockToKey                = ""
	RouteRelationsUnrelateBlockFromKey            = ""
	RouteRelationsRelatePodToKey                  = ""
	RouteRelationsUnrelatePodFromKey              = ""
	RouteRelationsRelatePodToBlock                = ""
	RouteRelationsUnrelatePodFromBlock            = ""
	RouteRelationsRelateBlockToBlock              = ""
	RouteRelationsUnrelateBlockFromBlock          = ""
	RouteRelationsRelatePodToPod                  = ""
	RouteRelationsUnrelatePodFromPod              = ""
)

const (
	RouteScalesGetScales           = ""
	RouteScalesAddScale            = ""
	RouteScalesGetScale            = ""
	RouteScalesUpdateScale         = ""
	RouteScalesDeleteScale         = ""
	RouteScalesGetBlocksUsingScale = ""
	RouteScalesGetPodsUsingScale   = ""
)

const (
	RouteSchedulerGetAllEventsInGivenWindow = ""
	RouteSchedulerGetAllEventsForGivenDay   = ""
	RouteSchedulerGetStandaloneEvents       = ""
	RouteSchedulerAddStandaloneEvent        = ""
	RouteSchedulerUpdateStandaloneEvent     = ""
	RouteSchedulerDeleteStandaloneEvent     = ""
)

const (
	RouteSearchSearchKeyBlockOrPodByToken = ""
	RouteSearchSearchUserByToken          = ""
)

const (
	RouteTeacherKeysGetAttachmentSubmissionsAsStudent = ""
	RouteTeacherKeysGetCommentSubmissionsAsStudent    = ""
	RouteTeacherKeysGetAllStudentsInABlock            = ""
	RouteTeacherKeys2Teachers                         = ""
)

const (
	RouteTeacherKeysGetStudentAttachmentSubmissionsAsTeacher  = ""
	RouteTeacherKeysGetStudentCommentSubmissionsAsTeacher     = ""
	RouteTeacherKeysAddAttachmentToTeacherPodAsTeacher        = ""
	RouteTeacherKeysAddCommentToTeacherPodAsTeacher           = ""
	RouteTeacherKeysGetBlockAndPodsGradesForAStudentAsTeacher = ""
	RouteTeacherKeysPublishStudentGradesForABlock             = ""
	RouteTeacherKeysBulkPublishPodGradesForAStudent           = ""
	RouteTeacherKeysBulkPublishPodGradesForStudents           = ""
	RouteTeacherKeysGetBlockGradesForAllStudents              = ""
	RouteTeacherKeysGetPodGradesForAllStudents                = ""
	RouteTeacherKeysAssignGradeToStudent                      = ""
	RouteTeacherKeysAssignPodGradeForAStudentAsTeacher        = ""
	RouteTeacherKeysGetStudentProfile                         = ""
)

const (
	RouteTemplatesGetKeyTemplates   = ""
	RouteTemplatesGetBlockTemplates = ""
	RouteTemplatesGetPodTemplates   = ""
)

const (
	RouteUserGetUsers              = ""
	RouteUserGetUserByUuid         = ""
	RouteUserDeactivateUserAccount = ""
)

const (
	RouteVersionGetLatestVersion = ""
	RouteVersionGetAppStatus     = ""
)
