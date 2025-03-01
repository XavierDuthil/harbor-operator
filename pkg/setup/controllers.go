package setup

import (
	"context"
	"fmt"

	"github.com/goharbor/harbor-operator/controllers"
	"github.com/goharbor/harbor-operator/controllers/goharbor/chartmuseum"
	"github.com/goharbor/harbor-operator/controllers/goharbor/configuration"
	"github.com/goharbor/harbor-operator/controllers/goharbor/core"
	"github.com/goharbor/harbor-operator/controllers/goharbor/harbor"
	"github.com/goharbor/harbor-operator/controllers/goharbor/harborcluster"
	"github.com/goharbor/harbor-operator/controllers/goharbor/jobservice"
	notaryserver "github.com/goharbor/harbor-operator/controllers/goharbor/notaryserver"
	notarysigner "github.com/goharbor/harbor-operator/controllers/goharbor/notarysigner"
	"github.com/goharbor/harbor-operator/controllers/goharbor/portal"
	"github.com/goharbor/harbor-operator/controllers/goharbor/registry"
	"github.com/goharbor/harbor-operator/controllers/goharbor/registryctl"
	"github.com/goharbor/harbor-operator/controllers/goharbor/trivy"
	"github.com/goharbor/harbor-operator/pkg/config"
	commonCtrl "github.com/goharbor/harbor-operator/pkg/controller"
	"github.com/ovh/configstore"
	"github.com/pkg/errors"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const (
	ControllerDisabledSuffixConfigKey = "controller-disabled"
)

var controllersBuilder = map[controllers.Controller]func(context.Context, *configstore.Store) (commonCtrl.Reconciler, error){
	controllers.Core:                core.New,
	controllers.Harbor:              harbor.New,
	controllers.JobService:          jobservice.New,
	controllers.Registry:            registry.New,
	controllers.NotaryServer:        notaryserver.New,
	controllers.NotarySigner:        notarysigner.New,
	controllers.RegistryController:  registryctl.New,
	controllers.Portal:              portal.New,
	controllers.ChartMuseum:         chartmuseum.New,
	controllers.Trivy:               trivy.New,
	controllers.HarborCluster:       harborcluster.New,
	controllers.HarborConfiguration: configuration.New,
}

type ControllerFactory func(context.Context, string, string, *configstore.Store) (commonCtrl.Reconciler, error)

type Controller interface {
	WithManager(context.Context, manager.Manager) error
	IsEnabled(context.Context) (bool, error)
}

type controller struct {
	Name controllers.Controller
	New  func(context.Context, *configstore.Store) (commonCtrl.Reconciler, error)
}

func (c *controller) GetConfig(ctx context.Context) (*configstore.Store, error) {
	configStore := config.NewConfigWithDefaults()
	configStore.Env(c.Name.String())

	return configStore, nil
}

func (c *controller) WithManager(ctx context.Context, mgr manager.Manager) error {
	configStore, err := c.GetConfig(ctx)
	if err != nil {
		return errors.Wrap(err, "get configuration")
	}

	controller, err := c.New(ctx, configStore)
	if err != nil {
		return errors.Wrap(err, "create")
	}

	err = controller.SetupWithManager(ctx, mgr)

	return errors.Wrap(err, "setup")
}

func (c *controller) IsEnabled(ctx context.Context) (bool, error) {
	configKey := fmt.Sprintf("%s-%s", c.Name, ControllerDisabledSuffixConfigKey)

	disabled, err := configstore.GetItemValueBool(configKey)
	if err == nil {
		return !disabled, nil
	}

	if config.IsNotFound(err, configKey) {
		return true, nil
	}

	return false, err
}
