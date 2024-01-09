# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'UuidResult',
    'AwaitableUuidResult',
    'uuid',
    'uuid_output',
]

@pulumi.output_type
class UuidResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableUuidResult(UuidResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return UuidResult(
            result=self.result)


def uuid(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableUuidResult:
    """
    Returns a unique identifier string, generated and formatted as required by RFC 4122.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:uuid', __args__, opts=opts, typ=UuidResult).value

    return AwaitableUuidResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(uuid)
def uuid_output(opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[UuidResult]:
    """
    Returns a unique identifier string, generated and formatted as required by RFC 4122.
    """
    ...
