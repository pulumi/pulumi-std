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
    'MaxResult',
    'AwaitableMaxResult',
    'max',
    'max_output',
]

@pulumi.output_type
class MaxResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, float):
            raise TypeError("Expected argument 'result' to be a float")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> float:
        return pulumi.get(self, "result")


class AwaitableMaxResult(MaxResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return MaxResult(
            result=self.result)


def max(input: Optional[Sequence[float]] = None,
        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableMaxResult:
    """
    Returns the largest of the floats.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:max', __args__, opts=opts, typ=MaxResult).value

    return AwaitableMaxResult(
        result=pulumi.get(__ret__, 'result'))
def max_output(input: Optional[pulumi.Input[Sequence[float]]] = None,
               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[MaxResult]:
    """
    Returns the largest of the floats.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:max', __args__, opts=opts, typ=MaxResult)
    return __ret__.apply(lambda __response__: MaxResult(
        result=pulumi.get(__response__, 'result')))
