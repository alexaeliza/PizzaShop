// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package primitive

import (
	"context"
	"github.com/atomix/runtime/pkg/driver"
	"google.golang.org/grpc"
)

type Registrar[T Primitive] func(*grpc.Server, *Service[T], *Registry[T])

func New[T Primitive](resolver ClientResolver[T], registrar Registrar[T]) Type {
	return &primitiveType[T]{
		resolver:  resolver,
		registrar: registrar,
		registry:  NewRegistry[T](),
	}
}

type Type interface {
	Register(server *grpc.Server, connector Connector)
}

type primitiveType[T Primitive] struct {
	resolver  ClientResolver[T]
	registrar Registrar[T]
	registry  *Registry[T]
}

func (a *primitiveType[T]) Register(server *grpc.Server, connector Connector) {
	a.registrar(server, NewService[T](connector, a.resolver, a.registry), a.registry)
}

type Primitive interface {
	Close(ctx context.Context) error
}

type ClientResolver[T Primitive] func(client driver.Client) (*Client[T], bool)

// NewClient creates a new client for the given primitive type
func NewClient[T Primitive](getter func(ctx context.Context, name string) (T, error)) *Client[T] {
	return &Client[T]{
		getter: getter,
	}
}

type Client[T Primitive] struct {
	getter func(ctx context.Context, name string) (T, error)
}

func (c *Client[T]) GetPrimitive(ctx context.Context, name string) (T, error) {
	return c.getter(ctx, name)
}