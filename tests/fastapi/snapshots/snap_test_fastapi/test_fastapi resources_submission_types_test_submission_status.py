# This file was auto-generated by Fern from our API Definition.

from __future__ import annotations

import datetime as dt
import typing

import pydantic
import typing_extensions

from ....core.datetime_utils import serialize_datetime
from .error_info import ErrorInfo
from .running_submission_state import RunningSubmissionState
from .submission_status_for_test_case import SubmissionStatusForTestCase

T_Result = typing.TypeVar("T_Result")


class _Factory:
    def stopped(self) -> TestSubmissionStatus:
        return TestSubmissionStatus(__root__=_TestSubmissionStatus.Stopped())

    def errored(self, value: ErrorInfo) -> TestSubmissionStatus:
        return TestSubmissionStatus(__root__=_TestSubmissionStatus.Errored(value=value))

    def running(self, value: RunningSubmissionState) -> TestSubmissionStatus:
        return TestSubmissionStatus(__root__=_TestSubmissionStatus.Running(value=value))

    def test_case_id_to_state(self, value: typing.Dict[str, SubmissionStatusForTestCase]) -> TestSubmissionStatus:
        return TestSubmissionStatus(__root__=_TestSubmissionStatus.TestCaseIdToState(value=value))


class TestSubmissionStatus(pydantic.BaseModel):
    factory: typing.ClassVar[_Factory] = _Factory()

    def get_as_union(
        self,
    ) -> typing.Union[
        _TestSubmissionStatus.Stopped,
        _TestSubmissionStatus.Errored,
        _TestSubmissionStatus.Running,
        _TestSubmissionStatus.TestCaseIdToState,
    ]:
        return self.__root__

    def visit(
        self,
        stopped: typing.Callable[[], T_Result],
        errored: typing.Callable[[ErrorInfo], T_Result],
        running: typing.Callable[[RunningSubmissionState], T_Result],
        test_case_id_to_state: typing.Callable[[typing.Dict[str, SubmissionStatusForTestCase]], T_Result],
    ) -> T_Result:
        if self.__root__.type == "stopped":
            return stopped()
        if self.__root__.type == "errored":
            return errored(self.__root__.value)
        if self.__root__.type == "running":
            return running(self.__root__.value)
        if self.__root__.type == "testCaseIdToState":
            return test_case_id_to_state(self.__root__.value)

    __root__: typing_extensions.Annotated[
        typing.Union[
            _TestSubmissionStatus.Stopped,
            _TestSubmissionStatus.Errored,
            _TestSubmissionStatus.Running,
            _TestSubmissionStatus.TestCaseIdToState,
        ],
        pydantic.Field(discriminator="type"),
    ]

    class Validators:
        """
        Use this class to add validators to the Pydantic model.

            @TestSubmissionStatus.Validators.validate
            def validate(value: typing.Union[_TestSubmissionStatus.Stopped, _TestSubmissionStatus.Errored, _TestSubmissionStatus.Running, _TestSubmissionStatus.TestCaseIdToState]) -> typing.Union[_TestSubmissionStatus.Stopped, _TestSubmissionStatus.Errored, _TestSubmissionStatus.Running, _TestSubmissionStatus.TestCaseIdToState]:
                ...
        """

        _validators: typing.ClassVar[
            typing.List[
                typing.Callable[
                    [
                        typing.Union[
                            _TestSubmissionStatus.Stopped,
                            _TestSubmissionStatus.Errored,
                            _TestSubmissionStatus.Running,
                            _TestSubmissionStatus.TestCaseIdToState,
                        ]
                    ],
                    typing.Union[
                        _TestSubmissionStatus.Stopped,
                        _TestSubmissionStatus.Errored,
                        _TestSubmissionStatus.Running,
                        _TestSubmissionStatus.TestCaseIdToState,
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
                        _TestSubmissionStatus.Stopped,
                        _TestSubmissionStatus.Errored,
                        _TestSubmissionStatus.Running,
                        _TestSubmissionStatus.TestCaseIdToState,
                    ]
                ],
                typing.Union[
                    _TestSubmissionStatus.Stopped,
                    _TestSubmissionStatus.Errored,
                    _TestSubmissionStatus.Running,
                    _TestSubmissionStatus.TestCaseIdToState,
                ],
            ],
        ) -> None:
            cls._validators.append(validator)

    @pydantic.root_validator(pre=False)
    def _validate(cls, values: typing.Dict[str, typing.Any]) -> typing.Dict[str, typing.Any]:
        value = typing.cast(
            typing.Union[
                _TestSubmissionStatus.Stopped,
                _TestSubmissionStatus.Errored,
                _TestSubmissionStatus.Running,
                _TestSubmissionStatus.TestCaseIdToState,
            ],
            values.get("__root__"),
        )
        for validator in TestSubmissionStatus.Validators._validators:
            value = validator(value)
        return {**values, "__root__": value}

    def json(self, **kwargs: typing.Any) -> str:
        kwargs_with_defaults: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        return super().json(**kwargs_with_defaults)

    def dict(self, **kwargs: typing.Any) -> typing.Dict[str, typing.Any]:
        kwargs_with_defaults: typing.Any = {"by_alias": True, "exclude_unset": True, **kwargs}
        return super().dict(**kwargs_with_defaults)

    class Config:
        frozen = True
        extra = pydantic.Extra.forbid
        json_encoders = {dt.datetime: serialize_datetime}


class _TestSubmissionStatus:
    class Stopped(pydantic.BaseModel):
        type: typing_extensions.Literal["stopped"] = "stopped"

        class Config:
            frozen = True

    class Errored(pydantic.BaseModel):
        type: typing_extensions.Literal["errored"] = "errored"
        value: ErrorInfo

        class Config:
            frozen = True

    class Running(pydantic.BaseModel):
        type: typing_extensions.Literal["running"] = "running"
        value: RunningSubmissionState

        class Config:
            frozen = True

    class TestCaseIdToState(pydantic.BaseModel):
        type: typing_extensions.Literal["testCaseIdToState"] = "testCaseIdToState"
        value: typing.Dict[str, SubmissionStatusForTestCase]

        class Config:
            frozen = True


TestSubmissionStatus.update_forward_refs()
