package social_grpc

import social "msa_big_tech/social/pkg/proto/v1"

func ConvertStatusToProto(status string) social.Status {
	switch status {
	case "PENDING", "pending":
		return social.Status_STATUS_PENDING
	case "ACCEPTED", "accepted", "APPROVED", "approved":
		return social.Status_STATUS_APPROVED
	case "DECLINED", "declined":
		return social.Status_STATUS_DECLINED
	default:
		return social.Status_STATUS_PENDING
	}
}
