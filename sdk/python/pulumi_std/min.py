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
    'MinResult',
    'AwaitableMinResult',
    'min',
    'min_output',
]

@pulumi.output_type
class MinResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, float):
            raise TypeError("Expected argument 'result' to be a float")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> float:
        return pulumi.get(self, "result")


class AwaitableMinResult(MinResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return MinResult(
            result=self.result)


def min(input: Optional[Sequence[float]] = None,
        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableMinResult:
    """
    Returns the smallest of the floats.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:min', __args__, opts=opts, typ=MinResult).value

    return AwaitableMinResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(min)
def min_output(input: Optional[pulumi.Input[Sequence[float]]] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[MinResult]:
    """
    Returns the smallest of the floats.
    """
    ...
