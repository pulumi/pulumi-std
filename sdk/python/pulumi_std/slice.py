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
    'SliceResult',
    'AwaitableSliceResult',
    'slice',
    'slice_output',
]

@pulumi.output_type
class SliceResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[Any]:
        return pulumi.get(self, "result")


class AwaitableSliceResult(SliceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SliceResult(
            result=self.result)


def slice(from_: Optional[int] = None,
          list: Optional[Sequence[Any]] = None,
          to: Optional[int] = None,
          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSliceResult:
    """
    Returns the portion of list between from (inclusive) and to (exclusive).
    """
    __args__ = dict()
    __args__['from'] = from_
    __args__['list'] = list
    __args__['to'] = to
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:slice', __args__, opts=opts, typ=SliceResult).value

    return AwaitableSliceResult(
        result=pulumi.get(__ret__, 'result'))
def slice_output(from_: Optional[pulumi.Input[Optional[int]]] = None,
                 list: Optional[pulumi.Input[Sequence[Any]]] = None,
                 to: Optional[pulumi.Input[Optional[int]]] = None,
                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[SliceResult]:
    """
    Returns the portion of list between from (inclusive) and to (exclusive).
    """
    __args__ = dict()
    __args__['from'] = from_
    __args__['list'] = list
    __args__['to'] = to
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:slice', __args__, opts=opts, typ=SliceResult)
    return __ret__.apply(lambda __response__: SliceResult(
        result=pulumi.get(__response__, 'result')))
