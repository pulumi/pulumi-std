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
    'SortResult',
    'AwaitableSortResult',
    'sort',
    'sort_output',
]

@pulumi.output_type
class SortResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, list):
            raise TypeError("Expected argument 'result' to be a list")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> Sequence[str]:
        return pulumi.get(self, "result")


class AwaitableSortResult(SortResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return SortResult(
            result=self.result)


def sort(input: Optional[Sequence[str]] = None,
         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableSortResult:
    """
    Returns a list of strings sorted lexicographically.
    """
    __args__ = dict()
    __args__['input'] = input
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:sort', __args__, opts=opts, typ=SortResult).value

    return AwaitableSortResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(sort)
def sort_output(input: Optional[pulumi.Input[Sequence[str]]] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[SortResult]:
    """
    Returns a list of strings sorted lexicographically.
    """
    ...
