// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"

	datacatalog "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/datacatalog"
	grpc "google.golang.org/grpc"

	mock "github.com/stretchr/testify/mock"
)

// DataCatalogClient is an autogenerated mock type for the DataCatalogClient type
type DataCatalogClient struct {
	mock.Mock
}

// AddTag provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) AddTag(ctx context.Context, in *datacatalog.AddTagRequest, opts ...grpc.CallOption) (*datacatalog.AddTagResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.AddTagResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.AddTagRequest, ...grpc.CallOption) *datacatalog.AddTagResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.AddTagResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.AddTagRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateArtifact provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) CreateArtifact(ctx context.Context, in *datacatalog.CreateArtifactRequest, opts ...grpc.CallOption) (*datacatalog.CreateArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.CreateArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.CreateArtifactRequest, ...grpc.CallOption) *datacatalog.CreateArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.CreateArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.CreateArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// CreateDataset provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) CreateDataset(ctx context.Context, in *datacatalog.CreateDatasetRequest, opts ...grpc.CallOption) (*datacatalog.CreateDatasetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.CreateDatasetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.CreateDatasetRequest, ...grpc.CallOption) *datacatalog.CreateDatasetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.CreateDatasetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.CreateDatasetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetArtifact provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) GetArtifact(ctx context.Context, in *datacatalog.GetArtifactRequest, opts ...grpc.CallOption) (*datacatalog.GetArtifactResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GetArtifactResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetArtifactRequest, ...grpc.CallOption) *datacatalog.GetArtifactResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetArtifactResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetArtifactRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataset provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) GetDataset(ctx context.Context, in *datacatalog.GetDatasetRequest, opts ...grpc.CallOption) (*datacatalog.GetDatasetResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.GetDatasetResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.GetDatasetRequest, ...grpc.CallOption) *datacatalog.GetDatasetResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.GetDatasetResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.GetDatasetRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListArtifacts provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) ListArtifacts(ctx context.Context, in *datacatalog.ListArtifactsRequest, opts ...grpc.CallOption) (*datacatalog.ListArtifactsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.ListArtifactsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.ListArtifactsRequest, ...grpc.CallOption) *datacatalog.ListArtifactsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.ListArtifactsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.ListArtifactsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListDatasets provides a mock function with given fields: ctx, in, opts
func (_m *DataCatalogClient) ListDatasets(ctx context.Context, in *datacatalog.ListDatasetsRequest, opts ...grpc.CallOption) (*datacatalog.ListDatasetsResponse, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, in)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *datacatalog.ListDatasetsResponse
	if rf, ok := ret.Get(0).(func(context.Context, *datacatalog.ListDatasetsRequest, ...grpc.CallOption) *datacatalog.ListDatasetsResponse); ok {
		r0 = rf(ctx, in, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*datacatalog.ListDatasetsResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *datacatalog.ListDatasetsRequest, ...grpc.CallOption) error); ok {
		r1 = rf(ctx, in, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
