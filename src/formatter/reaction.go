package formatter

import "grafana-matrix-forwarder/model"

func GenerateReaction(alert model.Data) string {
	if alert.State == model.AlertStateResolved {
		return resolvedReactionStr
	}
	return ""
}
