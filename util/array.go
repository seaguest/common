package util

import "bytes"

func ContainString(list []string, e string) bool {
	for _, a := range list {
		if a == e {
			return true
		}
	}
	return false
}

/*
	This method allows to compute the Intersect of two list.

	input:
		oldIds			:	a list of old ids
		newIds			:	a list of new ids

	return:
		excludedIds		:	the ids which are in old ids but not in new ids
		keptIds			:	the ids which are both in old ids and new ids
		addedIds		:	the ids which are in new ids but not in old olds
*/
func IntersectUuid(oldIds, newIds [][]byte) (excludedIds [][]byte, keptIds [][]byte, addedIds [][]byte) {
	for _, oldId := range oldIds {
		found := false
		for _, newId := range newIds {
			if bytes.Equal(oldId, newId) {
				found = true
				keptIds = append(keptIds, oldId)
				break
			}
		}
		if !found {
			excludedIds = append(excludedIds, oldId)
		}
	}

	for _, newId := range newIds {
		found := false
		for _, keptId := range keptIds {
			if bytes.Equal(newId, keptId) {
				found = true
				break
			}
		}
		if !found {
			addedIds = append(addedIds, newId)
		}
	}
	return excludedIds, keptIds, addedIds
}

/*
	This method allows to compute the Intersect of two list.

	input:
		oldIds			:	a list of old ids
		newIds			:	a list of new ids

	return:
		excludedIds		:	the ids which are in old ids but not in new ids
		keptIds			:	the ids which are both in old ids and new ids
		addedIds		:	the ids which are in new ids but not in old olds
*/
func IntersectInt(oldIds, newIds []uint32) (excludedIds, keptIds, addedIds []uint32) {
	for _, oldId := range oldIds {
		found := false
		for _, newId := range newIds {
			if oldId == newId {
				found = true
				keptIds = append(keptIds, oldId)
				break
			}
		}
		if !found {
			excludedIds = append(excludedIds, oldId)
		}
	}

	for _, newId := range newIds {
		found := false
		for _, keptId := range keptIds {
			if newId == keptId {
				found = true
				break
			}
		}
		if !found {
			addedIds = append(addedIds, newId)
		}
	}
	return excludedIds, keptIds, addedIds
}

func IntersectString(oldIds, newIds []string) (excludedIds, keptIds, addedIds []string) {
	for _, oldId := range oldIds {
		found := false
		for _, newId := range newIds {
			if oldId == newId {
				found = true
				keptIds = append(keptIds, oldId)
				break
			}
		}
		if !found {
			excludedIds = append(excludedIds, oldId)
		}
	}

	for _, newId := range newIds {
		found := false
		for _, keptId := range keptIds {
			if newId == keptId {
				found = true
				break
			}
		}
		if !found {
			addedIds = append(addedIds, newId)
		}
	}
	return excludedIds, keptIds, addedIds
}
