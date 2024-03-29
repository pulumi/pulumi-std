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
    'CoalescelistResult',
    'AwaitableCoalescelistResult',
    'coalescelist',
    'coalescelist_output',
]

@pulumi.output_type
class CoalescelistResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[Any]:
        return pulumi.get(self, "result")


class AwaitableCoalescelistResult(CoalescelistResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return CoalescelistResult(
            result=self.result)


def coalescelist(input: Optional[Sequence[Sequence[Any]]] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableCoalescelistResult:
    """
    Returns the first non-empty list from the given list of lists.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:coalescelist', __args__, opts=opts, typ=CoalescelistResult).value

    return AwaitableCoalescelistResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(coalescelist)
def coalescelist_output(input: Optional[pulumi.Input[Sequence[Sequence[Any]]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[CoalescelistResult]:
    """
    Returns the first non-empty list from the given list of lists.
    """
    ...
