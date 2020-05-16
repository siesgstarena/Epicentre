package utils

import (
	"fmt"
)

// IdentifyEventTypeGithub This will identify the exact key with which we have stored in DB
func IdentifyEventTypeGithub(event string) string {
    fmt.Print("In the IdentifyEventTypeGithub")
    switch event {
        case "checkRun":
            return "check_run"
        case "issueComment":
            return "issue_comment"
        case "pullRequest":
            return "pull_request"
        case "pullRequestReview":
            return "pull_request_review"
        case "pullRequestReviewComment":
            return "pull_request_review_comment"
        case "projectCard":
            return "project_card"
        case "projectColumn":
            return "project_column"
        default:
            return event
    }
}

// IdentifyEventTypeHeroku This will identify the exact key with which we have stored in DB
func IdentifyEventTypeHeroku(event string) string {
    fmt.Print("In the IdentifyEventTypeHeroku")
    switch event {
        case "addonAttachment":
            return "addon-attachment"
        default:
            return event
    }
}