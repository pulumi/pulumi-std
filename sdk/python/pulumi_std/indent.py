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
    'IndentResult',
    'AwaitableIndentResult',
    'indent',
    'indent_output',
]

@pulumi.output_type
class IndentResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableIndentResult(IndentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return IndentResult(
            result=self.result)


def indent(input: Optional[str] = None,
           spaces: Optional[int] = None,
           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableIndentResult:
    """
    Adds a given number of spaces after each newline character in the given string.
    """
    __args__ = dict()
    __args__['input'] = input
    __args__['spaces'] = spaces
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:indent', __args__, opts=opts, typ=IndentResult).value

    return AwaitableIndentResult(
        result=pulumi.get(__ret__, 'result'))


@_utilities.lift_output_func(indent)
def indent_output(input: Optional[pulumi.Input[str]] = None,
                  spaces: Optional[pulumi.Input[int]] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[IndentResult]:
    """
    Adds a given number of spaces after each newline character in the given string.
    """
    ...
