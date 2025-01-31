# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = [
    'TimestampResult',
    'AwaitableTimestampResult',
    'timestamp',
    'timestamp_output',
]

@pulumi.output_type
class TimestampResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableTimestampResult(TimestampResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TimestampResult(
            result=self.result)


def timestamp(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTimestampResult:
    """
    Returns a UTC timestamp string of the current time in RFC 3339 format
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:timestamp', __args__, opts=opts, typ=TimestampResult).value

    return AwaitableTimestampResult(
        result=pulumi.get(__ret__, 'result'))
def timestamp_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[TimestampResult]:
    """
    Returns a UTC timestamp string of the current time in RFC 3339 format
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:timestamp', __args__, opts=opts, typ=TimestampResult)
    return __ret__.apply(lambda __response__: TimestampResult(
        result=pulumi.get(__response__, 'result')))
