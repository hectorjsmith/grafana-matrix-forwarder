package formatter

import "grafana-matrix-forwarder/model"

func GenerateReaction(alert model.AlertData) string {
	if alert.State == model.AlertStateResolved {
		return resolvedReactionStr
	}
	return ""
}
