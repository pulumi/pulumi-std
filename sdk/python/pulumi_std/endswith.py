# coding=utf-8
# *** WARNING: this file was generated by pulumi. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'EndswithResult',
    'AwaitableEndswithResult',
    'endswith',
    'endswith_output',
]

@pulumi.output_type
class EndswithResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, bool):
            raise TypeError("Expected argument 'result' to be a bool")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> bool:
        return pulumi.get(self, "result")


class AwaitableEndswithResult(EndswithResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return EndswithResult(
            result=self.result)


def endswith(input: Optional[str] = None,
             suffix: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableEndswithResult:
    """
    Determines if the input string ends with the suffix.
    """
    __args__ = dict()
    __args__['input'] = input
    __args__['suffix'] = suffix
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:endswith', __args__, opts=opts, typ=EndswithResult).value

    return AwaitableEndswithResult(
        result=__ret__.result)


@_utilities.lift_output_func(endswith)
def endswith_output(input: Optional[pulumi.Input[str]] = None,
                    suffix: Optional[pulumi.Input[str]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[EndswithResult]:
    """
    Determines if the input string ends with the suffix.
    """
    ...
