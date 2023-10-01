package analytics

import "idp_system/app/framework/ctx"

type Analytics interface {
	// Identify ties a user to actions and traits.
	Identify(userID string, traits map[string]string)
	// Track records the a action the user performs.
	Track(eventName string, properties map[string]string, userID string, ctx ctx.ExecutionContext)
	// Group associates an identified user with a group.
	Group(userID string, groupID string)
	// Alias associates an identity with another.
	Alias(prevUserID string, newUserID string)
}
