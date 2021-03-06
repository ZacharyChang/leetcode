# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: hack/readme/leetcode/leetcode.proto

import sys

_b = sys.version_info[0] < 3 and (lambda x: x) or (lambda x: x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database

# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()

DESCRIPTOR = _descriptor.FileDescriptor(
    name='hack/readme/leetcode/leetcode.proto',
    package='login',
    syntax='proto3',
    serialized_options=None,
    serialized_pb=_b(
        '\n#hack/readme/leetcode/leetcode.proto\x12\x05login\".\n\x0cLoginRequest\x12\x0c\n\x04name\x18\x01 \x01(\t\x12\x10\n\x08password\x18\x02 \x01(\t\"\x18\n\x16ListAllProblemsRequest\"#\n\x13QueryProblemRequest\x12\x0c\n\x04slug\x18\x01 \x01(\t\"\x1a\n\x08Response\x12\x0e\n\x06result\x18\x01 \x01(\t2\xc6\x01\n\x0fLeetcodeService\x12/\n\x05Login\x12\x13.login.LoginRequest\x1a\x0f.login.Response\"\x00\x12\x43\n\x0fListAllProblems\x12\x1d.login.ListAllProblemsRequest\x1a\x0f.login.Response\"\x00\x12=\n\x0cQueryProblem\x12\x1a.login.QueryProblemRequest\x1a\x0f.login.Response\"\x00\x62\x06proto3')
)

_LOGINREQUEST = _descriptor.Descriptor(
    name='LoginRequest',
    full_name='login.LoginRequest',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='name', full_name='login.LoginRequest.name', index=0,
            number=1, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=_b("").decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
        _descriptor.FieldDescriptor(
            name='password', full_name='login.LoginRequest.password', index=1,
            number=2, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=_b("").decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=46,
    serialized_end=92,
)

_LISTALLPROBLEMSREQUEST = _descriptor.Descriptor(
    name='ListAllProblemsRequest',
    full_name='login.ListAllProblemsRequest',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=94,
    serialized_end=118,
)

_QUERYPROBLEMREQUEST = _descriptor.Descriptor(
    name='QueryProblemRequest',
    full_name='login.QueryProblemRequest',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='slug', full_name='login.QueryProblemRequest.slug', index=0,
            number=1, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=_b("").decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=120,
    serialized_end=155,
)

_RESPONSE = _descriptor.Descriptor(
    name='Response',
    full_name='login.Response',
    filename=None,
    file=DESCRIPTOR,
    containing_type=None,
    fields=[
        _descriptor.FieldDescriptor(
            name='result', full_name='login.Response.result', index=0,
            number=1, type=9, cpp_type=9, label=1,
            has_default_value=False, default_value=_b("").decode('utf-8'),
            message_type=None, enum_type=None, containing_type=None,
            is_extension=False, extension_scope=None,
            serialized_options=None, file=DESCRIPTOR),
    ],
    extensions=[
    ],
    nested_types=[],
    enum_types=[
    ],
    serialized_options=None,
    is_extendable=False,
    syntax='proto3',
    extension_ranges=[],
    oneofs=[
    ],
    serialized_start=157,
    serialized_end=183,
)

DESCRIPTOR.message_types_by_name['LoginRequest'] = _LOGINREQUEST
DESCRIPTOR.message_types_by_name['ListAllProblemsRequest'] = _LISTALLPROBLEMSREQUEST
DESCRIPTOR.message_types_by_name['QueryProblemRequest'] = _QUERYPROBLEMREQUEST
DESCRIPTOR.message_types_by_name['Response'] = _RESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

LoginRequest = _reflection.GeneratedProtocolMessageType('LoginRequest', (_message.Message,), {
    'DESCRIPTOR': _LOGINREQUEST,
    '__module__': 'hack.readme.leetcode.leetcode_pb2'
    # @@protoc_insertion_point(class_scope:login.LoginRequest)
})
_sym_db.RegisterMessage(LoginRequest)

ListAllProblemsRequest = _reflection.GeneratedProtocolMessageType('ListAllProblemsRequest', (_message.Message,), {
    'DESCRIPTOR': _LISTALLPROBLEMSREQUEST,
    '__module__': 'hack.readme.leetcode.leetcode_pb2'
    # @@protoc_insertion_point(class_scope:login.ListAllProblemsRequest)
})
_sym_db.RegisterMessage(ListAllProblemsRequest)

QueryProblemRequest = _reflection.GeneratedProtocolMessageType('QueryProblemRequest', (_message.Message,), {
    'DESCRIPTOR': _QUERYPROBLEMREQUEST,
    '__module__': 'hack.readme.leetcode.leetcode_pb2'
    # @@protoc_insertion_point(class_scope:login.QueryProblemRequest)
})
_sym_db.RegisterMessage(QueryProblemRequest)

Response = _reflection.GeneratedProtocolMessageType('Response', (_message.Message,), {
    'DESCRIPTOR': _RESPONSE,
    '__module__': 'hack.readme.leetcode.leetcode_pb2'
    # @@protoc_insertion_point(class_scope:login.Response)
})
_sym_db.RegisterMessage(Response)

_LEETCODESERVICE = _descriptor.ServiceDescriptor(
    name='LeetcodeService',
    full_name='login.LeetcodeService',
    file=DESCRIPTOR,
    index=0,
    serialized_options=None,
    serialized_start=186,
    serialized_end=384,
    methods=[
        _descriptor.MethodDescriptor(
            name='Login',
            full_name='login.LeetcodeService.Login',
            index=0,
            containing_service=None,
            input_type=_LOGINREQUEST,
            output_type=_RESPONSE,
            serialized_options=None,
        ),
        _descriptor.MethodDescriptor(
            name='ListAllProblems',
            full_name='login.LeetcodeService.ListAllProblems',
            index=1,
            containing_service=None,
            input_type=_LISTALLPROBLEMSREQUEST,
            output_type=_RESPONSE,
            serialized_options=None,
        ),
        _descriptor.MethodDescriptor(
            name='QueryProblem',
            full_name='login.LeetcodeService.QueryProblem',
            index=2,
            containing_service=None,
            input_type=_QUERYPROBLEMREQUEST,
            output_type=_RESPONSE,
            serialized_options=None,
        ),
    ])
_sym_db.RegisterServiceDescriptor(_LEETCODESERVICE)

DESCRIPTOR.services_by_name['LeetcodeService'] = _LEETCODESERVICE

# @@protoc_insertion_point(module_scope)
