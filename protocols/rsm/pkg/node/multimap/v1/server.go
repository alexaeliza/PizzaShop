// SPDX-FileCopyrightText: 2022-present Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package v1

import (
	"context"
	multimapprotocolv1 "github.com/atomix/atomix/protocols/rsm/api/multimap/v1"
	"github.com/atomix/atomix/protocols/rsm/pkg/node"
	"github.com/atomix/atomix/runtime/pkg/logging"
	streams "github.com/atomix/atomix/runtime/pkg/stream"
	"github.com/gogo/protobuf/proto"
	"google.golang.org/grpc"
)

var log = logging.GetLogger()

const truncLen = 200

func RegisterServer(node *node.Node) {
	node.RegisterService(func(server *grpc.Server) {
		multimapprotocolv1.RegisterMultiMapServer(server, NewMultiMapServer(node))
	})
}

var serverCodec = node.NewCodec[*multimapprotocolv1.MultiMapInput, *multimapprotocolv1.MultiMapOutput](
	func(input *multimapprotocolv1.MultiMapInput) ([]byte, error) {
		return proto.Marshal(input)
	},
	func(bytes []byte) (*multimapprotocolv1.MultiMapOutput, error) {
		output := &multimapprotocolv1.MultiMapOutput{}
		if err := proto.Unmarshal(bytes, output); err != nil {
			return nil, err
		}
		return output, nil
	})

func NewMultiMapServer(protocol node.Protocol) multimapprotocolv1.MultiMapServer {
	return &multiMapServer{
		handler: node.NewHandler[*multimapprotocolv1.MultiMapInput, *multimapprotocolv1.MultiMapOutput](protocol, serverCodec),
	}
}

type multiMapServer struct {
	handler node.Handler[*multimapprotocolv1.MultiMapInput, *multimapprotocolv1.MultiMapOutput]
}

func (s *multiMapServer) Size(ctx context.Context, request *multimapprotocolv1.SizeRequest) (*multimapprotocolv1.SizeResponse, error) {
	log.Debugw("Size",
		logging.Stringer("SizeRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Size_{
			Size_: request.SizeInput,
		},
	}
	output, headers, err := s.handler.Query(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Size",
			logging.Stringer("SizeRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.SizeResponse{
		Headers:    headers,
		SizeOutput: output.GetSize_(),
	}
	log.Debugw("Size",
		logging.Stringer("SizeRequest", request),
		logging.Stringer("SizeResponse", response))
	return response, nil
}

func (s *multiMapServer) Put(ctx context.Context, request *multimapprotocolv1.PutRequest) (*multimapprotocolv1.PutResponse, error) {
	log.Debugw("Put",
		logging.Stringer("PutRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Put{
			Put: request.PutInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Put",
			logging.Stringer("PutRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.PutResponse{
		Headers:   headers,
		PutOutput: output.GetPut(),
	}
	log.Debugw("Put",
		logging.Stringer("PutRequest", request),
		logging.Stringer("PutResponse", response))
	return response, nil
}

func (s *multiMapServer) PutAll(ctx context.Context, request *multimapprotocolv1.PutAllRequest) (*multimapprotocolv1.PutAllResponse, error) {
	log.Debugw("PutAll",
		logging.Stringer("PutAllRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_PutAll{
			PutAll: request.PutAllInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("PutAll",
			logging.Stringer("PutAllRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.PutAllResponse{
		Headers:      headers,
		PutAllOutput: output.GetPutAll(),
	}
	log.Debugw("PutAll",
		logging.Stringer("PutAllRequest", request),
		logging.Stringer("PutAllResponse", response))
	return response, nil
}

func (s *multiMapServer) PutEntries(ctx context.Context, request *multimapprotocolv1.PutEntriesRequest) (*multimapprotocolv1.PutEntriesResponse, error) {
	log.Debugw("PutEntries",
		logging.Stringer("PutEntriesRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_PutEntries{
			PutEntries: request.PutEntriesInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("PutEntries",
			logging.Stringer("PutEntriesRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.PutEntriesResponse{
		Headers:          headers,
		PutEntriesOutput: output.GetPutEntries(),
	}
	log.Debugw("PutEntries",
		logging.Stringer("PutEntriesRequest", request),
		logging.Stringer("PutEntriesResponse", response))
	return response, nil
}

func (s *multiMapServer) Replace(ctx context.Context, request *multimapprotocolv1.ReplaceRequest) (*multimapprotocolv1.ReplaceResponse, error) {
	log.Debugw("Replace",
		logging.Stringer("ReplaceRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Replace{
			Replace: request.ReplaceInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Replace",
			logging.Stringer("ReplaceRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.ReplaceResponse{
		Headers:       headers,
		ReplaceOutput: output.GetReplace(),
	}
	log.Debugw("Replace",
		logging.Stringer("ReplaceRequest", request),
		logging.Stringer("ReplaceResponse", response))
	return response, nil
}

func (s *multiMapServer) Contains(ctx context.Context, request *multimapprotocolv1.ContainsRequest) (*multimapprotocolv1.ContainsResponse, error) {
	log.Debugw("Contains",
		logging.Stringer("ContainsRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Contains{
			Contains: request.ContainsInput,
		},
	}
	output, headers, err := s.handler.Query(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Contains",
			logging.Stringer("ContainsRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.ContainsResponse{
		Headers:        headers,
		ContainsOutput: output.GetContains(),
	}
	log.Debugw("Contains",
		logging.Stringer("ContainsRequest", request),
		logging.Stringer("ContainsResponse", response))
	return response, nil
}

func (s *multiMapServer) Get(ctx context.Context, request *multimapprotocolv1.GetRequest) (*multimapprotocolv1.GetResponse, error) {
	log.Debugw("Get",
		logging.Stringer("GetRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Get{
			Get: request.GetInput,
		},
	}
	output, headers, err := s.handler.Query(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Get",
			logging.Stringer("GetRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.GetResponse{
		Headers:   headers,
		GetOutput: output.GetGet(),
	}
	log.Debugw("Get",
		logging.Stringer("GetRequest", request),
		logging.Stringer("GetResponse", response))
	return response, nil
}

func (s *multiMapServer) Remove(ctx context.Context, request *multimapprotocolv1.RemoveRequest) (*multimapprotocolv1.RemoveResponse, error) {
	log.Debugw("Remove",
		logging.Stringer("RemoveRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Remove{
			Remove: request.RemoveInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Remove",
			logging.Stringer("RemoveRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.RemoveResponse{
		Headers:      headers,
		RemoveOutput: output.GetRemove(),
	}
	log.Debugw("Remove",
		logging.Stringer("RemoveRequest", request),
		logging.Stringer("RemoveResponse", response))
	return response, nil
}

func (s *multiMapServer) RemoveAll(ctx context.Context, request *multimapprotocolv1.RemoveAllRequest) (*multimapprotocolv1.RemoveAllResponse, error) {
	log.Debugw("RemoveAll",
		logging.Stringer("RemoveAllRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_RemoveAll{
			RemoveAll: request.RemoveAllInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("RemoveAll",
			logging.Stringer("RemoveAllRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.RemoveAllResponse{
		Headers:         headers,
		RemoveAllOutput: output.GetRemoveAll(),
	}
	log.Debugw("RemoveAll",
		logging.Stringer("RemoveAllRequest", request),
		logging.Stringer("RemoveAllResponse", response))
	return response, nil
}

func (s *multiMapServer) RemoveEntries(ctx context.Context, request *multimapprotocolv1.RemoveEntriesRequest) (*multimapprotocolv1.RemoveEntriesResponse, error) {
	log.Debugw("RemoveEntries",
		logging.Stringer("RemoveEntriesRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_RemoveEntries{
			RemoveEntries: request.RemoveEntriesInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("RemoveEntries",
			logging.Stringer("RemoveEntriesRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.RemoveEntriesResponse{
		Headers:             headers,
		RemoveEntriesOutput: output.GetRemoveEntries(),
	}
	log.Debugw("RemoveEntries",
		logging.Stringer("RemoveEntriesRequest", request),
		logging.Stringer("RemoveEntriesResponse", response))
	return response, nil
}

func (s *multiMapServer) Clear(ctx context.Context, request *multimapprotocolv1.ClearRequest) (*multimapprotocolv1.ClearResponse, error) {
	log.Debugw("Clear",
		logging.Stringer("ClearRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Clear{
			Clear: request.ClearInput,
		},
	}
	output, headers, err := s.handler.Propose(ctx, input, request.Headers)
	if err != nil {
		log.Warnw("Clear",
			logging.Stringer("ClearRequest", request),
			logging.Error("Error", err))
		return nil, err
	}
	response := &multimapprotocolv1.ClearResponse{
		Headers:     headers,
		ClearOutput: output.GetClear(),
	}
	log.Debugw("Clear",
		logging.Stringer("ClearRequest", request),
		logging.Stringer("ClearResponse", response))
	return response, nil
}

func (s *multiMapServer) Events(request *multimapprotocolv1.EventsRequest, server multimapprotocolv1.MultiMap_EventsServer) error {
	log.Debugw("Events",
		logging.Stringer("EventsRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Events{
			Events: request.EventsInput,
		},
	}

	stream := streams.NewBufferedStream[*node.StreamProposalResponse[*multimapprotocolv1.MultiMapOutput]]()
	go func() {
		err := s.handler.StreamPropose(server.Context(), input, request.Headers, stream)
		if err != nil {
			log.Warnw("Events",
				logging.Stringer("EventsRequest", request),
				logging.Error("Error", err))
			stream.Error(err)
			stream.Close()
		}
	}()

	for {
		result, ok := stream.Receive()
		if !ok {
			return nil
		}

		if result.Failed() {
			log.Warnw("Events",
				logging.Stringer("EventsRequest", request),
				logging.Error("Error", result.Error))
			return result.Error
		}

		response := &multimapprotocolv1.EventsResponse{
			Headers:      result.Value.Headers,
			EventsOutput: result.Value.Output.GetEvents(),
		}
		log.Debugw("Events",
			logging.Stringer("EventsRequest", request),
			logging.Stringer("EventsResponse", response))
		if err := server.Send(response); err != nil {
			log.Warnw("Events",
				logging.Stringer("EventsRequest", request),
				logging.Error("Error", err))
			return err
		}
	}
}

func (s *multiMapServer) Entries(request *multimapprotocolv1.EntriesRequest, server multimapprotocolv1.MultiMap_EntriesServer) error {
	log.Debugw("Entries",
		logging.Stringer("EntriesRequest", request))
	input := &multimapprotocolv1.MultiMapInput{
		Input: &multimapprotocolv1.MultiMapInput_Entries{
			Entries: request.EntriesInput,
		},
	}

	stream := streams.NewBufferedStream[*node.StreamQueryResponse[*multimapprotocolv1.MultiMapOutput]]()
	go func() {
		err := s.handler.StreamQuery(server.Context(), input, request.Headers, stream)
		if err != nil {
			log.Warnw("Entries",
				logging.Stringer("EntriesRequest", request),
				logging.Error("Error", err))
			stream.Error(err)
			stream.Close()
		}
	}()

	for {
		result, ok := stream.Receive()
		if !ok {
			return nil
		}

		if result.Failed() {
			log.Warnw("Entries",
				logging.Stringer("EntriesRequest", request),
				logging.Error("Error", result.Error))
			return result.Error
		}

		response := &multimapprotocolv1.EntriesResponse{
			Headers:       result.Value.Headers,
			EntriesOutput: result.Value.Output.GetEntries(),
		}
		log.Debugw("Entries",
			logging.Stringer("EntriesRequest", request),
			logging.Stringer("EntriesResponse", response))
		if err := server.Send(response); err != nil {
			log.Warnw("Entries",
				logging.Stringer("EntriesRequest", request),
				logging.Error("Error", err))
			return err
		}
	}
}
