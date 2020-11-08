from typing import get_type_hints
from functools import wraps
from inspect import getfullargspec

# 函数调用和返回值的类型声明
# def add(a: int, b: int) -> int:
#     validate_input(add, a=a, b=b)
#     return a + b


def validate_input(obj, **kwargs):
    hints = get_type_hints(obj)
    for para_name, para_type in hints.items():
        if para_name == 'return':
            continue

        if not isinstance(kwargs[para_name], para_type):
            raise TypeError('参数: {} 类型错误，　应该是: {}'.format(
                para_name, para_type))


def type_check(decorator):
    @wraps(decorator)
    def wrapped_decorator(*args, **kwargs):
        func_args = getfullargspec(decorator)[0]
        kwargs.update(dict(zip(func_args, args)))

        validate_input(decorator, **kwargs)
        return decorator(**kwargs)

    return wrapped_decorator


# 函数调用和返回值的类型声明
@type_check
def add(a: int, b: int) -> int:
    return a + b


if __name__ == "__main__":
    # add('a', '1') 会报错
    print(add(2, 5))