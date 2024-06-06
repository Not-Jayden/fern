# This file was auto-generated by Fern from our API Definition.

import datetime as dt
import typing
import uuid
from json.decoder import JSONDecodeError

from ..core.api_error import ApiError
from ..core.client_wrapper import AsyncClientWrapper, SyncClientWrapper
from ..core.datetime_utils import serialize_datetime
from ..core.jsonable_encoder import jsonable_encoder
from ..core.pydantic_utilities import pydantic_v1
from ..core.request_options import RequestOptions
from .types.nested_user import NestedUser
from .types.user import User


class UserClient:
    def __init__(self, *, client_wrapper: SyncClientWrapper):
        self._client_wrapper = client_wrapper

    def get_username(
        self,
        *,
        limit: typing.Optional[int] = None,
        id: typing.Optional[uuid.UUID] = None,
        date: typing.Optional[dt.date] = None,
        deadline: typing.Optional[dt.datetime] = None,
        bytes: typing.Optional[str] = None,
        user: typing.Optional[User] = None,
        key_value: typing.Optional[typing.Dict[str, str]] = None,
        optional_string: typing.Optional[str] = None,
        nested_user: typing.Optional[NestedUser] = None,
        optional_user: typing.Optional[User] = None,
        exclude_user: typing.Optional[typing.Union[User, typing.Sequence[User]]] = None,
        filter: typing.Optional[typing.Union[str, typing.Sequence[str]]] = None,
        request_options: typing.Optional[RequestOptions] = None
    ) -> User:
        """
        Parameters
        ----------
        limit : typing.Optional[int]

        id : typing.Optional[uuid.UUID]

        date : typing.Optional[dt.date]

        deadline : typing.Optional[dt.datetime]

        bytes : typing.Optional[str]

        user : typing.Optional[User]

        key_value : typing.Optional[typing.Dict[str, str]]

        optional_string : typing.Optional[str]

        nested_user : typing.Optional[NestedUser]

        optional_user : typing.Optional[User]

        exclude_user : typing.Optional[typing.Union[User, typing.Sequence[User]]]

        filter : typing.Optional[typing.Union[str, typing.Sequence[str]]]

        request_options : typing.Optional[RequestOptions]
            Request-specific configuration.

        Returns
        -------
        User

        Examples
        --------
        import uuid

        from seed.client import SeedQueryParameters

        client = SeedQueryParameters(
            base_url="https://yourhost.com/path/to/api",
        )
        client.user.get_username(
            limit=5,
            id=uuid.UUID(
                "4ff45b32-ca63-462d-b988-cf4eec41397a",
            ),
        )
        """
        _response = self._client_wrapper.httpx_client.request(
            "user",
            method="GET",
            params={
                "limit": limit,
                "id": jsonable_encoder(id),
                "date": str(date) if date is not None else None,
                "deadline": serialize_datetime(deadline) if deadline is not None else None,
                "bytes": jsonable_encoder(bytes),
                "user": jsonable_encoder(user),
                "keyValue": jsonable_encoder(key_value),
                "optionalString": optional_string,
                "nestedUser": jsonable_encoder(nested_user),
                "optionalUser": jsonable_encoder(optional_user),
                "excludeUser": jsonable_encoder(exclude_user),
                "filter": filter,
            },
            request_options=request_options,
        )
        if 200 <= _response.status_code < 300:
            return pydantic_v1.parse_obj_as(User, _response.json())  # type: ignore
        try:
            _response_json = _response.json()
        except JSONDecodeError:
            raise ApiError(status_code=_response.status_code, body=_response.text)
        raise ApiError(status_code=_response.status_code, body=_response_json)


class AsyncUserClient:
    def __init__(self, *, client_wrapper: AsyncClientWrapper):
        self._client_wrapper = client_wrapper

    async def get_username(
        self,
        *,
        limit: typing.Optional[int] = None,
        id: typing.Optional[uuid.UUID] = None,
        date: typing.Optional[dt.date] = None,
        deadline: typing.Optional[dt.datetime] = None,
        bytes: typing.Optional[str] = None,
        user: typing.Optional[User] = None,
        key_value: typing.Optional[typing.Dict[str, str]] = None,
        optional_string: typing.Optional[str] = None,
        nested_user: typing.Optional[NestedUser] = None,
        optional_user: typing.Optional[User] = None,
        exclude_user: typing.Optional[typing.Union[User, typing.Sequence[User]]] = None,
        filter: typing.Optional[typing.Union[str, typing.Sequence[str]]] = None,
        request_options: typing.Optional[RequestOptions] = None
    ) -> User:
        """
        Parameters
        ----------
        limit : typing.Optional[int]

        id : typing.Optional[uuid.UUID]

        date : typing.Optional[dt.date]

        deadline : typing.Optional[dt.datetime]

        bytes : typing.Optional[str]

        user : typing.Optional[User]

        key_value : typing.Optional[typing.Dict[str, str]]

        optional_string : typing.Optional[str]

        nested_user : typing.Optional[NestedUser]

        optional_user : typing.Optional[User]

        exclude_user : typing.Optional[typing.Union[User, typing.Sequence[User]]]

        filter : typing.Optional[typing.Union[str, typing.Sequence[str]]]

        request_options : typing.Optional[RequestOptions]
            Request-specific configuration.

        Returns
        -------
        User

        Examples
        --------
        import uuid

        from seed.client import AsyncSeedQueryParameters

        client = AsyncSeedQueryParameters(
            base_url="https://yourhost.com/path/to/api",
        )
        await client.user.get_username(
            limit=5,
            id=uuid.UUID(
                "4ff45b32-ca63-462d-b988-cf4eec41397a",
            ),
        )
        """
        _response = await self._client_wrapper.httpx_client.request(
            "user",
            method="GET",
            params={
                "limit": limit,
                "id": jsonable_encoder(id),
                "date": str(date) if date is not None else None,
                "deadline": serialize_datetime(deadline) if deadline is not None else None,
                "bytes": jsonable_encoder(bytes),
                "user": jsonable_encoder(user),
                "keyValue": jsonable_encoder(key_value),
                "optionalString": optional_string,
                "nestedUser": jsonable_encoder(nested_user),
                "optionalUser": jsonable_encoder(optional_user),
                "excludeUser": jsonable_encoder(exclude_user),
                "filter": filter,
            },
            request_options=request_options,
        )
        if 200 <= _response.status_code < 300:
            return pydantic_v1.parse_obj_as(User, _response.json())  # type: ignore
        try:
            _response_json = _response.json()
        except JSONDecodeError:
            raise ApiError(status_code=_response.status_code, body=_response.text)
        raise ApiError(status_code=_response.status_code, body=_response_json)
