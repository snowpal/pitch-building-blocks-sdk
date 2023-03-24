package main

import log "github.com/sirupsen/logrus"

func main() {
	log.Info("Recipe1: Register few users")
	RegisterFewUsers()

	log.Info("Recipe2: Get resource attributes")
	GetResourceAttributes()

	log.Info("Recipe3: Create private conversation")
	CreatePrivateConversation()

	log.Info("Recipe4: Add and link resources")
	AddAndLinkResources()

	log.Info("Recipe5: Share a block")
	ShareBlock()

	log.Info("Recipe6: Get all keys")
	GetAllKeys()

	log.Info("Recipe7: Add & remove favorite")
	AddFavorite()

	log.Info("Recipe8: Fetch scheduler events")
	FetchScheduler()

	log.Info("Recipe9: Add & remove relations")
	AddRelation()

	log.Info("Recipe10: Publish pod grade for a student")
	PublishStudentGrade()

	log.Info("Recipe11: Add project list and project pod")
	AddProjectList()

	log.Info("Recipe12: Grant access level for block")
	GrantAclOnCustomBlock()

	log.Info("Recipe13: Update attributes")
	UpdateAttributes()
}
