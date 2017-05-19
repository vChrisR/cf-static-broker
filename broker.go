package main

import (
	"context"

	"code.cloudfoundry.org/lager"
	"github.com/pivotal-cf/brokerapi"
)

//New Returns a ServiceBroker
func New(services []brokerapi.Service, logger lager.Logger, env Config) brokerapi.ServiceBroker {
	return &broker{services: services, logger: logger, env: env}
}

type broker struct {
	services []brokerapi.Service
	creds    staticStoreType
	logger   lager.Logger
	env      Config
}

func (b *broker) Services(context context.Context) []brokerapi.Service {
	return b.services
}

func (b *broker) Provision(context context.Context, instanceID string, details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (b *broker) Deprovision(context context.Context, instanceID string, details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (b *broker) Bind(context context.Context, instanceID, bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	var bindingSpec brokerapi.Binding
	bindingSpec.Credentials = StaticStoreGetCredsForPlanID(b.creds, details.PlanID)
	return bindingSpec, nil
}

func (b *broker) Unbind(context context.Context, instanceID, bindingID string, details brokerapi.UnbindDetails) error {
	return nil
}

func (b *broker) Update(context context.Context, instanceID string, details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (b *broker) LastOperation(context context.Context, instanceID, operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}
