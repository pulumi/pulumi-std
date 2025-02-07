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
    'TimecmpResult',
    'AwaitableTimecmpResult',
    'timecmp',
    'timecmp_output',
]

@pulumi.output_type
class TimecmpResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, int):
            raise TypeError("Expected argument 'result' to be a int")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> int:
        return pulumi.get(self, "result")


class AwaitableTimecmpResult(TimecmpResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return TimecmpResult(
            result=self.result)


def timecmp(timestampa: Optional[str] = None,
            timestampb: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableTimecmpResult:
    """
    Compares two timestamps and returns a number that represents the ordering
    	of the instants those timestamps represent.
    	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
    	Both timestamps must be strings adhering this syntax, i.e. "2017-11-22T00:00:00Z".
    	If 'timestamp_a' is before 'timestamp_b', -1 is returned.
    	If 'timestamp_a' is equal to 'timestamp_b', 0 is returned.
    	If 'timestamp_a' is after 'timestamp_b', 1 is returned.
    """
    __args__ = dict()
    __args__['timestampa'] = timestampa
    __args__['timestampb'] = timestampb
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:timecmp', __args__, opts=opts, typ=TimecmpResult).value

    return AwaitableTimecmpResult(
        result=pulumi.get(__ret__, 'result'))
def timecmp_output(timestampa: Optional[pulumi.Input[str]] = None,
                   timestampb: Optional[pulumi.Input[str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[TimecmpResult]:
    """
    Compares two timestamps and returns a number that represents the ordering
    	of the instants those timestamps represent.
    	Timestamps are represented as strings using RFC 3339 "Date and time format" syntax.
    	Both timestamps must be strings adhering this syntax, i.e. "2017-11-22T00:00:00Z".
    	If 'timestamp_a' is before 'timestamp_b', -1 is returned.
    	If 'timestamp_a' is equal to 'timestamp_b', 0 is returned.
    	If 'timestamp_a' is after 'timestamp_b', 1 is returned.
    """
    __args__ = dict()
    __args__['timestampa'] = timestampa
    __args__['timestampb'] = timestampb
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:timecmp', __args__, opts=opts, typ=TimecmpResult)
    return __ret__.apply(lambda __response__: TimecmpResult(
        result=pulumi.get(__response__, 'result')))
