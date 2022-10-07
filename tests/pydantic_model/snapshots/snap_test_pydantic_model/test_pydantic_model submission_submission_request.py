from __future__ import annotations

import typing

import pydantic
import typing_extensions

from .initialize_problem_request import InitializeProblemRequest
from .stop_request import StopRequest
from .submit_request_v_2 import SubmitRequestV2
from .workspace_submit_request import WorkspaceSubmitRequest

T_Result = typing.TypeVar("T_Result")


class _Factory:
    def initialize_problem_request(self, value: InitializeProblemRequest) -> SubmissionRequest:
        return SubmissionRequest(
            __root__=_SubmissionRequest.InitializeProblemRequest(**dict(value), type="initializeProblemRequest")
        )

    def initialize_workspace_request(self) -> SubmissionRequest:
        return SubmissionRequest(
            __root__=_SubmissionRequest.InitializeWorkspaceRequest(type="initializeWorkspaceRequest")
        )

    def submit_v_2(self, value: SubmitRequestV2) -> SubmissionRequest:
        return SubmissionRequest(__root__=_SubmissionRequest.SubmitV2(**dict(value), type="submitV2"))

    def workspace_submit(self, value: WorkspaceSubmitRequest) -> SubmissionRequest:
        return SubmissionRequest(__root__=_SubmissionRequest.WorkspaceSubmit(**dict(value), type="workspaceSubmit"))

    def stop(self, value: StopRequest) -> SubmissionRequest:
        return SubmissionRequest(__root__=_SubmissionRequest.Stop(**dict(value), type="stop"))


class SubmissionRequest(pydantic.BaseModel):
    factory: typing.ClassVar[_Factory] = _Factory()

    def get_as_union(
        self,
    ) -> typing.Union[
        _SubmissionRequest.InitializeProblemRequest,
        _SubmissionRequest.InitializeWorkspaceRequest,
        _SubmissionRequest.SubmitV2,
        _SubmissionRequest.WorkspaceSubmit,
        _SubmissionRequest.Stop,
    ]:
        return self.__root__

    def visit(
        self,
        initialize_problem_request: typing.Callable[[InitializeProblemRequest], T_Result],
        initialize_workspace_request: typing.Callable[[], T_Result],
        submit_v_2: typing.Callable[[SubmitRequestV2], T_Result],
        workspace_submit: typing.Callable[[WorkspaceSubmitRequest], T_Result],
        stop: typing.Callable[[StopRequest], T_Result],
    ) -> T_Result:
        if self.__root__.type == "initializeProblemRequest":
            return initialize_problem_request(self.__root__)
        if self.__root__.type == "initializeWorkspaceRequest":
            return initialize_workspace_request()
        if self.__root__.type == "submitV2":
            return submit_v_2(self.__root__)
        if self.__root__.type == "workspaceSubmit":
            return workspace_submit(self.__root__)
        if self.__root__.type == "stop":
            return stop(self.__root__)

    __root__: typing_extensions.Annotated[
        typing.Union[
            _SubmissionRequest.InitializeProblemRequest,
            _SubmissionRequest.InitializeWorkspaceRequest,
            _SubmissionRequest.SubmitV2,
            _SubmissionRequest.WorkspaceSubmit,
            _SubmissionRequest.Stop,
        ],
        pydantic.Field(discriminator="type"),
    ]

    @pydantic.root_validator
    def _validate(cls, values: typing.Dict[str, typing.Any]) -> typing.Dict[str, typing.Any]:
        value = typing.cast(
            typing.Union[
                _SubmissionRequest.InitializeProblemRequest,
                _SubmissionRequest.InitializeWorkspaceRequest,
                _SubmissionRequest.SubmitV2,
                _SubmissionRequest.WorkspaceSubmit,
                _SubmissionRequest.Stop,
            ],
            values.get("__root__"),
        )
        for validator in SubmissionRequest.Validators._validators:
            value = validator(value)
        return {**values, "__root__": value}

    class Validators:
        _validators: typing.ClassVar[
            typing.List[
                typing.Callable[
                    [
                        typing.Union[
                            _SubmissionRequest.InitializeProblemRequest,
                            _SubmissionRequest.InitializeWorkspaceRequest,
                            _SubmissionRequest.SubmitV2,
                            _SubmissionRequest.WorkspaceSubmit,
                            _SubmissionRequest.Stop,
                        ]
                    ],
                    typing.Union[
                        _SubmissionRequest.InitializeProblemRequest,
                        _SubmissionRequest.InitializeWorkspaceRequest,
                        _SubmissionRequest.SubmitV2,
                        _SubmissionRequest.WorkspaceSubmit,
                        _SubmissionRequest.Stop,
                    ],
                ]
            ]
        ] = []

        @classmethod
        def validate(
            cls,
            validator: typing.Callable[
                [
                    typing.Union[
                        _SubmissionRequest.InitializeProblemRequest,
                        _SubmissionRequest.InitializeWorkspaceRequest,
                        _SubmissionRequest.SubmitV2,
                        _SubmissionRequest.WorkspaceSubmit,
                        _SubmissionRequest.Stop,
                    ]
                ],
                typing.Union[
                    _SubmissionRequest.InitializeProblemRequest,
                    _SubmissionRequest.InitializeWorkspaceRequest,
                    _SubmissionRequest.SubmitV2,
                    _SubmissionRequest.WorkspaceSubmit,
                    _SubmissionRequest.Stop,
                ],
            ],
        ) -> None:
            cls._validators.append(validator)

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults: typing.Any = {"by_alias": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    class Config:
        frozen = True


class _SubmissionRequest:
    class InitializeProblemRequest(InitializeProblemRequest):
        type: typing_extensions.Literal["initializeProblemRequest"]

        class Config:
            frozen = True

    class InitializeWorkspaceRequest(pydantic.BaseModel):
        type: typing_extensions.Literal["initializeWorkspaceRequest"]

        class Config:
            frozen = True

    class SubmitV2(SubmitRequestV2):
        type: typing_extensions.Literal["submitV2"]

        class Config:
            frozen = True

    class WorkspaceSubmit(WorkspaceSubmitRequest):
        type: typing_extensions.Literal["workspaceSubmit"]

        class Config:
            frozen = True

    class Stop(StopRequest):
        type: typing_extensions.Literal["stop"]

        class Config:
            frozen = True


SubmissionRequest.update_forward_refs()
