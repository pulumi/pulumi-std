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
    'ReplaceResult',
    'AwaitableReplaceResult',
    'replace',
    'replace_output',
]

@pulumi.output_type
class ReplaceResult:
    def __init__(__self__, result=None):
        if result and not isinstance(result, str):
            raise TypeError("Expected argument 'result' to be a str")
        pulumi.set(__self__, "result", result)

    @property
    @pulumi.getter
    def result(self) -> str:
        return pulumi.get(self, "result")


class AwaitableReplaceResult(ReplaceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return ReplaceResult(
            result=self.result)


def replace(replace: Optional[str] = None,
            search: Optional[str] = None,
            text: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableReplaceResult:
    """
    Does a search and replace on the given string.
    All instances of search are replaced with the value of replace.
    If search is wrapped in forward slashes, it is treated as a regular expression.
    If using a regular expression, replace can reference subcaptures in the regular expression by
    using $n where n is the index or name of the subcapture. If using a regular expression,
    the syntax conforms to the re2 regular expression syntax.
    """
    __args__ = dict()
    __args__['replace'] = replace
    __args__['search'] = search
    __args__['text'] = text
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('std:index:replace', __args__, opts=opts, typ=ReplaceResult).value

    return AwaitableReplaceResult(
        result=pulumi.get(__ret__, 'result'))
def replace_output(replace: Optional[pulumi.Input[str]] = None,
                   search: Optional[pulumi.Input[str]] = None,
                   text: Optional[pulumi.Input[str]] = None,
                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[ReplaceResult]:
    """
    Does a search and replace on the given string.
    All instances of search are replaced with the value of replace.
    If search is wrapped in forward slashes, it is treated as a regular expression.
    If using a regular expression, replace can reference subcaptures in the regular expression by
    using $n where n is the index or name of the subcapture. If using a regular expression,
    the syntax conforms to the re2 regular expression syntax.
    """
    __args__ = dict()
    __args__['replace'] = replace
    __args__['search'] = search
    __args__['text'] = text
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('std:index:replace', __args__, opts=opts, typ=ReplaceResult)
    return __ret__.apply(lambda __response__: ReplaceResult(
        result=pulumi.get(__response__, 'result')))
