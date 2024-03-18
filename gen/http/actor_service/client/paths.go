// Code generated by goa v3.15.2, DO NOT EDIT.
//
// HTTP request path constructors for the ActorService service.
//
// Command:
// $ goa gen film-lib/design

package client

import (
	"fmt"
)

// GetAllActorsActorServicePath returns the URL path to the ActorService service getAllActors HTTP endpoint.
func GetAllActorsActorServicePath() string {
	return "/actors"
}

// AddActorActorServicePath returns the URL path to the ActorService service addActor HTTP endpoint.
func AddActorActorServicePath() string {
	return "/actors"
}

// UpdateActorInfoActorServicePath returns the URL path to the ActorService service updateActorInfo HTTP endpoint.
func UpdateActorInfoActorServicePath(actorID uint64) string {
	return fmt.Sprintf("/actors/%v", actorID)
}

// DeleteActorActorServicePath returns the URL path to the ActorService service deleteActor HTTP endpoint.
func DeleteActorActorServicePath(actorID uint64) string {
	return fmt.Sprintf("/actors/%v", actorID)
}