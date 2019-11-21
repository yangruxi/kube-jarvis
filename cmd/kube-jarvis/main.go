package main

import (
	"context"

	"github.com/RayHuangCN/kube-jarvis/pkg/logger"
	_ "github.com/RayHuangCN/kube-jarvis/pkg/plugins/coordinate/all"
	_ "github.com/RayHuangCN/kube-jarvis/pkg/plugins/diagnose/all"
	_ "github.com/RayHuangCN/kube-jarvis/pkg/plugins/evaluate/all"
	_ "github.com/RayHuangCN/kube-jarvis/pkg/plugins/export/all"
)

func main() {
	config, err := GetConfig("conf/default.yaml", logger.NewLogger())
	if err != nil {
		panic(err)
	}

	cli, err := config.GetClusterClient()
	if err != nil {
		panic(err)
	}

	coordinator, err := config.GetCoordinator()
	if err != nil {
		panic(err)
	}

	trans, err := config.GetTranslator()
	if err != nil {
		panic(err)
	}

	diagnostics, err := config.GetDiagnostics(cli, trans)
	if err != nil {
		panic(err)
	}

	for _, d := range diagnostics {
		coordinator.AddDiagnostic(d)
	}

	evaluators, err := config.GetEvaluators(trans)
	if err != nil {
		panic(err)
	}

	for _, e := range evaluators {
		coordinator.AddEvaluate(e)
	}

	exporters, err := config.GetExporters()
	if err != nil {
		panic(err)
	}

	for _, e := range exporters {
		coordinator.AddExporter(e)
	}

	coordinator.Run(context.Background())
}